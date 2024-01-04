package utils

import (
	"encoding/json"
	"os"
)

type CustomizationOptions struct {
	BotName string `json:"BotName"`
	DocsURL string `json:"DocsURL"`
	CDN     struct {
		BaseURL string
		Files   map[string]string
	}
}

var BotCustoms CustomizationOptions

func init() {
	customsFile, err := os.Open("customization.json")
	if err != nil {
		Log.Error("Couldn't load customizations file: " + err.Error())
	}

	jsonParser := json.NewDecoder(customsFile)
	if err = jsonParser.Decode(&BotCustoms); err != nil {
		Log.Error("Couldn't parse customizations file: " + err.Error())
	}
}

func GetCdnUri(fileName string) string {
	if val, ok := BotCustoms.CDN.Files[fileName]; ok {
		return BotCustoms.CDN.BaseURL + val
	} else {
		Log.Error("Couldn't file an customization entry for the file " + fileName)
	}
	return ""
}
