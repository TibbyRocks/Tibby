package utils

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func ReturnInteractionMessageUrl(i *discordgo.InteractionCreate) string {
	channelId := i.ChannelID
	messageID := i.ApplicationCommandData().TargetID
	guildID := i.GuildID

	var url string

	if guildID != "" {
		url = fmt.Sprintf("https://discord.com/channels/%v/%v/%v", guildID, channelId, messageID)
	} else {
		url = fmt.Sprintf("https://discord.com/channels/@me/%v/%v", channelId, messageID)
	}

	return url

}
