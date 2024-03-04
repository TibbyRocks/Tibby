package tibbycmds

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/tibbyrocks/tibby/internal/commands"
	"github.com/tibbyrocks/tibby/internal/utils"
)

var (
	Commands []commands.Command
	customs  = utils.GetCustoms()
)

func init() {
	Commands = append(Commands, commands.Command{
		AppComm: &RootCommand,
		Handler: RootCommandHandler,
	})
}

var RootCommand = discordgo.ApplicationCommand{
	Name:        customs.RootCommand,
	Description: fmt.Sprintf("General commands for %s", customs.BotName),
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "docs",
			Description: fmt.Sprintf("Get the %s docs", customs.BotName),
			Type:        discordgo.ApplicationCommandOptionSubCommand,
		},
		{
			Name:        "info",
			Description: fmt.Sprintf("Get the %s runtime information", customs.BotName),
			Type:        discordgo.ApplicationCommandOptionSubCommand,
		},
	},
}

func RootCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	switch options[0].Name {
	case "docs":
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags: discordgo.MessageFlagsEphemeral,
				Embeds: []*discordgo.MessageEmbed{
					{
						Color:       int(utils.HexToDec("BFEA7C")),
						Description: getDocs(i),
						Author: &discordgo.MessageEmbedAuthor{
							Name:    fmt.Sprintf("%s docs", customs.BotName),
							IconURL: s.State.User.AvatarURL("1024"),
						},
					},
				},
			},
		})
	case "info":
		infoMsg := GetInfo(i, s)
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags: discordgo.MessageFlagsEphemeral,
				Embeds: []*discordgo.MessageEmbed{
					{
						Color:       int(utils.HexToDec("BFEA7C")),
						Description: infoMsg,
						Author: &discordgo.MessageEmbedAuthor{
							Name:    fmt.Sprintf("%s Info", customs.BotName),
							IconURL: s.State.User.AvatarURL("1024"),
						},
					},
				},
			},
		})
	}
}
