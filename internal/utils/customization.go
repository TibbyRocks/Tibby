package utils

import (
	"encoding/json"
	"os"
	"slices"

	"github.com/bwmarrin/discordgo"
)

type CustomizationOptions struct {
	BotName     string `json:"BotName"`
	RootCommand string `json:"RootCommand"`
	DocsURL     string `json:"DocsURL"`
	CDN         struct {
		BaseURL string
		Files   map[string]string
	}
	BotAdmins []string `json:"BotAdmins"`
}

func GetCustoms() CustomizationOptions {
	customsFile, err := os.Open("customizations/botproperties.json")
	if err != nil {
		Log.Error("Couldn't load customizations file: " + err.Error())
		os.Exit(1)
	}

	var BotCustoms CustomizationOptions

	jsonParser := json.NewDecoder(customsFile)
	if err = jsonParser.Decode(&BotCustoms); err != nil {
		Log.Error("Couldn't parse customizations file: " + err.Error())
		os.Exit(1)
	}

	return BotCustoms
}

func GetCdnUri(fileName string) string {
	BotCustoms := GetCustoms()
	if val, ok := BotCustoms.CDN.Files[fileName]; ok {
		return BotCustoms.CDN.BaseURL + val
	} else {
		Log.Error("Couldn't find an customization entry for the file " + fileName)
	}
	return ""
}

func UserIsBotAdmin(userID string) bool {
	BotCustoms := GetCustoms()
	if slices.Contains(BotCustoms.BotAdmins, userID) {
		return true
	} else {
		return false
	}
}

func InteractionAdminCheck(s *discordgo.Session, i *discordgo.InteractionCreate) bool {
	if !UserIsBotAdmin(GetUserobjectFromInteraction(i).ID) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "You need to be a bot admin to perform that command.",
			},
		})
		return false
	} else {
		return true
	}
}
