package radlibs

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/tibbyrocks/tibby/internal/commands"
	"github.com/tibbyrocks/tibby/internal/types"
	"github.com/tibbyrocks/tibby/internal/utils"
)

type Randomizer = types.Randomizer

var (
	singularNouns Randomizer
	pluralNouns   Randomizer
	verbs         Randomizer
	adjectives    Randomizer
	animals       Randomizer
	fruit         Randomizer
	gerunds       Randomizer
	adverbs       Randomizer
	Command       commands.Command
	Commands      []commands.Command
	customs       = utils.GetCustoms()
)

var RadLibsCommand = discordgo.ApplicationCommand{
	Name:        "radlibs",
	Description: "Replaces certain tokens with words",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "msg",
			Description: "Message with the tokens you want to radlib",
			Required:    true,
		},
	},
}

func init() {
	Commands = append(Commands, commands.Command{
		AppComm: &RadLibsCommand,
		Handler: RadlibsCommandHandler,
	})
}

func init() {
	pluralNouns.Fill("customizations/pluralnouns.txt", true)
	verbs.Fill("customizations/verbs.txt", true)
	adjectives.Fill("customizations/adjectives.txt", true)
	singularNouns.Fill("customizations/singularnouns.txt", true)
	animals.Fill("customizations/animals.txt", true)
	fruit.Fill("customizations/fruit.txt", true)
	singularNouns.Combine(&animals, &fruit)
	gerunds.Fill("customizations/gerunds.txt", true)
	adverbs.Fill("customizations/adverbs.txt", true)
}

func RadlibsCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {

	libbedMsg := DoRadlibs(s, i)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Description: libbedMsg,
					Author: &discordgo.MessageEmbedAuthor{
						Name:    fmt.Sprintf("%s RadLibs 😎", customs.BotName),
						IconURL: s.State.User.AvatarURL("1024"),
					},
				},
			},
		},
	})
}

// This function splits up the message string and loops over every string
func DoRadlibs(s *discordgo.Session, i *discordgo.InteractionCreate) string {
	utils.LogCmd(i)
	optionMap := utils.GetOptionsFromInteraction(i)
	message := optionMap["msg"].StringValue()
	members, _ := s.GuildMembers(i.GuildID, "", 1000)

	splitMessage := strings.Split(message, " ")
	var workSlice []string
	for _, s := range splitMessage {
		workSlice = append(workSlice, replaceRadlibToken(s, members))
	}
	return strings.Join(workSlice, " ")
}

func replaceRadlibToken(token string, members []*discordgo.Member) string {
	token = strings.ReplaceAll(token, "$ANIMAL", animals.Random())
	token = strings.ReplaceAll(token, "$FRUIT", fruit.Random())
	token = strings.ReplaceAll(token, "$NOUNS", pluralNouns.Random())
	token = strings.ReplaceAll(token, "$NOUN", singularNouns.Random())
	token = strings.ReplaceAll(token, "$VERBING", gerunds.Random())
	token = strings.ReplaceAll(token, "$VERB", verbs.Random())
	token = strings.ReplaceAll(token, "$ADJ", adjectives.Random())
	token = strings.ReplaceAll(token, "$ADVERB", adverbs.Random())
	token = strings.ReplaceAll(token, "$CHATTER", utils.RandomMemberName(members))
	return token
}
