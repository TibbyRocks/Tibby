package utils

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

// Gets the actual Discord username for the user that invoked an interaction, don't use this when you want want a nickname instead.
func GetUsernameFromInteraction(i *discordgo.InteractionCreate) string {
	var username string
	if i.User != nil {
		username = i.User.Username
	} else {
		username = i.Member.User.Username
	}
	return username
}

func GetUserobjectFromInteraction(i *discordgo.InteractionCreate) *discordgo.User {
	if i.User != nil {
		return i.User
	} else {
		return i.Member.User
	}
}

// Gets the nickname for the user that invoked the interaction in a guild if set, if not set or invoked in a direct message you get the username.
func GetNickFromInteraction(i *discordgo.InteractionCreate) string {
	var username string
	if i.User != nil {
		username = i.User.Username
	} else if i.Member.Nick != "" {
		username = i.Member.Nick
	} else {
		username = i.Member.User.Username
	}
	return username
}

func GetOptionsFromInteraction(i *discordgo.InteractionCreate) map[string]*discordgo.ApplicationCommandInteractionDataOption {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}
	return optionMap
}

func GetOptionsStringsFromInteraction(i *discordgo.InteractionCreate) map[string]string {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]string, len(options))
	for _, opt := range options {
		switch opt.Type {
		case discordgo.ApplicationCommandOptionString:
			optionMap[opt.Name] = opt.StringValue()
		case discordgo.ApplicationCommandOptionSubCommand:
			optionMap["subcommand"] = opt.Name
		case discordgo.ApplicationCommandOptionSubCommandGroup:
			optionMap["subcommandgroup"] = opt.Name
		default:
			optionMap[opt.Name] = opt.Type.String()
		}

	}
	return optionMap
}

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

func HexToDec(hex string) int {
	dec, _ := strconv.ParseInt(hex, 16, 32)

	return int(dec)
}
