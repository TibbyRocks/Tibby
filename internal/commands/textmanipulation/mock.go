package textmanipulation

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/tibbyrocks/tibby/internal/commands"
	"github.com/tibbyrocks/tibby/internal/utils"
)

var (
	Command commands.Command
	//customs = utils.GetCustoms()
)

var MockifyCommand = discordgo.ApplicationCommand{
	Name: "Mockify",
	Type: discordgo.MessageApplicationCommand,
}

func init() {
	Commands = append(Commands, commands.Command{
		AppComm: &MockifyCommand,
		Handler: MockifyCommandHandler,
	})
}

func MockifyCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Description: MockText(i),
					Author: &discordgo.MessageEmbedAuthor{
						Name:    "Mock",
						IconURL: s.State.User.AvatarURL("1024"),
					},
				},
			},
		},
	})
}

func MockText(i *discordgo.InteractionCreate) string {
	utils.LogCmd(i)
	msg := i.ApplicationCommandData().Resolved.Messages[i.ApplicationCommandData().TargetID]
	input := msg.Content
	msgUrl := utils.ReturnInteractionMessageUrl(i)

	var mockedText string

	if len(msg.Embeds) > 0 {
		return fmt.Sprintf("I can't mock embeds, sorry.\n\n[Go to the original message](%v)", msgUrl)
	}

	if len(msg.Content) < 1 {
		return fmt.Sprintf("I can't mock messages without text, sorry.\n\n[Go to the original message](%v)", msgUrl)
	}

	textSlice := strings.Split(input, "")

	for k, v := range textSlice {
		if k%2 == 0 {
			textSlice[k] = strings.ToUpper(v)
		} else {
			textSlice[k] = strings.ToLower(v)
		}
	}

	var outputFormat string = `%v
	
	[Go to the original message](%v)
	`

	mockedText = strings.Join(textSlice, "")
	builtNessage := fmt.Sprintf(outputFormat, mockedText, msgUrl)
	return builtNessage
}
