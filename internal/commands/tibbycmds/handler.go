package tibbycmds

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/tibbyrocks/tibby/internal/utils"
)

var (
	customs = utils.GetCustoms()
)

func HandleRootCommand(i *discordgo.InteractionCreate, s *discordgo.Session) {
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
