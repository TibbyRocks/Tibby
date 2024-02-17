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
var log = Log

func init() {
	customsFile, err := os.Open("customizations/botproperties.json")
	if err != nil {
		log.Error("Couldn't load customizations file: " + err.Error())
	}

	jsonParser := json.NewDecoder(customsFile)
	if err = jsonParser.Decode(&BotCustoms); err != nil {
		log.Error("Couldn't parse customizations file: " + err.Error())
	}
}

func GetCdnUri(fileName string) string {
	if val, ok := BotCustoms.CDN.Files[fileName]; ok {
		return BotCustoms.CDN.BaseURL + val
	} else {
		log.Error("Couldn't file an customization entry for the file " + fileName)
	}
	return ""
}
