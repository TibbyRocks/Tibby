package translations

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
)

var (
	google_endpoint string = "https://translation.googleapis.com"
)

type gtTranslationResult struct {
	Data struct {
		Translations []struct {
			DetectedSourceLanguage string `json:"detectedSourceLanguage"`
			TranslatedText         string `json:"translatedText"`
		} `json:"translations"`
	} `json:"data"`
}

func gtAnyToLanguage(text string, language string) SingleTranslation {
	key := os.Getenv("WB_GOOGLE_API_KEY")
	reqUrl, _ := url.Parse(google_endpoint + "/language/translate/v2")
	q := reqUrl.Query()
	q.Add("q", text)
	q.Add("target", language)
	reqUrl.RawQuery = q.Encode()

	req, err := http.NewRequest("POST", reqUrl.String(), nil)
	if err != nil {
		log.Error("1" + err.Error())
	}

	req.Header.Add("x-goog-api-key", key)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("2" + err.Error())
	}

	var result gtTranslationResult
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Error("3" + err.Error())
	}

	finalTranslation := SingleTranslation{
		translatedText: result.Data.Translations[0].TranslatedText,
		fromLang:       msGetLanguageByCode(result.Data.Translations[0].DetectedSourceLanguage)[0],
		fromLangNative: msGetLanguageByCode(result.Data.Translations[0].DetectedSourceLanguage)[1],
		toLang:         msGetLanguageByCode(language)[0],
		toLangNative:   msGetLanguageByCode(language)[1],
		originalText:   text,
	}

	return finalTranslation

}
