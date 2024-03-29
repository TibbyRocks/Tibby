package textmanipulation

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/tibbyrocks/tibby/internal/commands"
	"github.com/tibbyrocks/tibby/internal/types"
	"github.com/tibbyrocks/tibby/internal/utils"
)

var UwuifyCommand = discordgo.ApplicationCommand{
	Name: "UwUify",
	Type: discordgo.MessageApplicationCommand,
}

type Randomizer = types.Randomizer

var (
	furryFlavourText Randomizer
	Commands         []commands.Command
)

func init() {
	Commands = append(Commands, commands.Command{
		AppComm: &UwuifyCommand,
		Handler: UwuifyCommandHandler,
	})
}

func UwuifyCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Description: Uwuify(i),
					Author: &discordgo.MessageEmbedAuthor{
						Name:    "UwUify",
						IconURL: s.State.User.AvatarURL("1024"),
					},
				},
			},
		},
	})
}

func Uwuify(i *discordgo.InteractionCreate) string {
	utils.LogCmd(i)
	msg := i.ApplicationCommandData().Resolved.Messages[i.ApplicationCommandData().TargetID]
	input := msg.Content
	msgUrl := utils.ReturnInteractionMessageUrl(i)

	if len(msg.Embeds) > 0 {
		return fmt.Sprintf("I can't mock embeds, sorry.\n\n[Go to the original message](%v)", msgUrl)
	}

	if len(msg.Content) < 1 {
		return fmt.Sprintf("I can't mock messages without text, sorry.\n\n[Go to the original message](%v)", msgUrl)
	}

	/*wordReplacer := strings.NewReplacer(
	"friend", "fwiendo",
	"ove", "uv",
	"hugs", "huggies",
	"hug", "huggy",
	"kisses", "smoochies",
	"bird", "birb",
	"chicken", "chinkem")*/
	letterReplacer := strings.NewReplacer(
		"r", "w",
		"R", "W",
		"l", "w",
		"L", "W")

	replacedText := letterReplacer.Replace(input)
	flavouredText := addFlavour(replacedText)

	var outputFormat string = `%v
	
	[Go to the original message](%v)
	`

	builtMessage := fmt.Sprintf(outputFormat, flavouredText, msgUrl)
	return builtMessage
}

func addFlavour(input string) string {
	furryFlavourText.Append(" UwU ", " OwO ", " ^_^ ", " :3 ", ` \*nuzzles u\* `, "~ ", " nya~ ", " (・`ω´・) ")

	textSlice := strings.Split(input, "")

	for k, v := range textSlice {
		if v == "!" {
			textSlice[k] = furryFlavourText.Random()
		}
	}
	return strings.Join(textSlice, "")
}
