package tibbycmds

import (
	"bytes"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/tibbyrocks/tibby/internal/utils"
)

var (
	log = utils.Log
)

func ShowBotGuilds(s *discordgo.Session, i *discordgo.InteractionCreate) {
	guilds, _ := s.UserGuilds(200, "", "", false)
	var guildsString bytes.Buffer
	for _, v := range guilds {
		fmt.Fprintf(&guildsString, "%s : %s\n", v.ID, v.Name)
	}
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
			Embeds: []*discordgo.MessageEmbed{
				{
					Color:       int(utils.HexToDec("BFEA7C")),
					Description: guildsString.String(),
					Author: &discordgo.MessageEmbedAuthor{
						Name:    fmt.Sprintf("%s Info", customs.BotName),
						IconURL: s.State.User.AvatarURL("1024"),
					},
				},
			},
		},
	})
}

func LeaveGuild(s *discordgo.Session, i *discordgo.InteractionCreate) {
	guildID := i.ApplicationCommandData().Options[0].Options[0].Options[0].Value.(string)

	if guildID == i.GuildID {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: fmt.Sprintf("Trying to leave %s, which is the current server", guildID),
			},
		})
		s.GuildLeave(guildID)
	} else {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: fmt.Sprintf("Trying to leave %s", guildID),
			},
		})
		err := s.GuildLeave(guildID)
		if err != nil {
			log.Error(fmt.Sprintf("Couldn't leave guild %s because %s", guildID, err.Error()))
		}
	}
}
