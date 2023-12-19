package utils

import "github.com/bwmarrin/discordgo"

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

// Gets the nickname for the user that invoked the interaction in a guild if set, if not set or invoked in a direct message you get the username.
func GetNickFromInteraction(i *discordgo.InteractionCreate) string {
	var username string
	if i.User != nil {
		username = i.User.Username
	} else {
		username = i.Member.Nick
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
