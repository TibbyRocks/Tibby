package googletranslate

import (
	"encoding/json"
	"html"
	"net/http"
	"net/url"
	"os"

	"github.com/tibbyrocks/tibby/internal/types"
	"github.com/tibbyrocks/tibby/internal/utils"
)

var (
	log        = utils.Log
	Translator types.Translator
)

func init() {
	Translator = types.Translator{
		Translate: Translate,
	}
}

var (
	google_endpoint string = "https://translation.googleapis.com"
)

type translationResult struct {
	Data struct {
		Translations []struct {
			DetectedSourceLanguage string `json:"detectedSourceLanguage"`
			TranslatedText         string `json:"translatedText"`
		} `json:"translations"`
	} `json:"data"`
}

func Translate(fromLang string, toLang string, translatable string) types.SingleTranslation {
	key := os.Getenv("WB_GOOGLE_API_KEY")
	reqUrl, _ := url.Parse(google_endpoint + "/language/translate/v2")
	q := reqUrl.Query()
	if fromLang != "" {
		q.Add("source", fromLang)
	}
	q.Add("q", translatable)
	q.Add("target", toLang)
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

	var result translationResult
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Error("3" + err.Error())
	}

	finalTranslation := types.SingleTranslation{
		// Google Translate takes the translatable string as a query parameter and has to be escaped.
		// So when the result comes back we have to unescape it back to the original symbols.
		TranslatedText: html.UnescapeString(result.Data.Translations[0].TranslatedText),
		FromLangCode:   result.Data.Translations[0].DetectedSourceLanguage,
		ToLangCode:     toLang,
		OriginalText:   translatable,
	}

	return finalTranslation

}
