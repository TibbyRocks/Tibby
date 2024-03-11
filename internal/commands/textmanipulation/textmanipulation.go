package textmanipulation

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tibbyrocks/tibby/internal/commands"
	"github.com/tibbyrocks/tibby/internal/utils"
)

var (
	Command  commands.Command
	Commands []commands.Command
	log      = utils.Log
	//customs = utils.GetCustoms()
)

var TextManipulationCommand = discordgo.ApplicationCommand{
	Name: "Text Manipulation",
	Type: discordgo.MessageApplicationCommand,
}

func init() {
	Commands = append(Commands, commands.Command{
		AppComm: &TextManipulationCommand,
		Handler: TextManipulationCommandHandler,
	})
}

func TextManipulationCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Title: "Text manipulation",
			Flags: discordgo.MessageFlagsEphemeral,
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       i.ApplicationCommandData().Resolved.Messages[i.ApplicationCommandData().TargetID].Author.Username,
					Description: i.ApplicationCommandData().Resolved.Messages[i.ApplicationCommandData().TargetID].Content,
				},
			},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.SelectMenu{
							CustomID: "text-manipulation",
							Options: []discordgo.SelectMenuOption{
								{
									Label: "Uwuify",
									Value: "uwuify",
									Emoji: discordgo.ComponentEmoji{
										Name: "🐍",
									},
								},
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Error(err.Error())
	}
}
