package translations

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"os"

	"github.com/mozoarella/wombot/internal/utils"
)

var (
	log         = utils.Log
	endpoint    = "https://api.cognitive.microsofttranslator.com"
	languageMap map[string]Language
)

func init() {
	languageMap = msGetLanguages()
}

type TranslationLanguages struct {
	Languages map[string]map[string]string `json:"translation"`
}

type Language struct {
	direction  string
	name       string
	nativeName string
}

type TranslationResult struct {
	DetectedLanguage struct {
		Name  string  `json:"language"`
		Score float32 `json:"score"`
	} `json:"detectedLanguage"`
	Translations []struct {
		Text       string `json:"text"`
		ToLanguage string `json:"to"`
	}
}

/*
Returns the full names of the language from the code given.

If the language is supported you get a slice with 2 items. The English name for a language, and the native name.

If the language is not supported you get an empty slice.
*/
func msGetLanguageByCode(code string) []string {
	var response []string

	if val, ok := languageMap[code]; ok {
		response = append(response, val.name, val.nativeName)
	}

	return response
}

func msGetLanguages() map[string]Language {
	reqUrl, _ := url.Parse(endpoint + "/languages")
	q := reqUrl.Query()
	q.Add("api-version", "3.0")
	q.Add("scope", "translation")
	reqUrl.RawQuery = q.Encode()

	body := []struct {
	}{}
	reqBody, _ := json.Marshal(body)

	req, err := http.NewRequest("GET", reqUrl.String(), bytes.NewBuffer(reqBody))
	if err != nil {
		log.Error("1" + err.Error())
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("2" + err.Error())
	}
	defer res.Body.Close()
	var result TranslationLanguages
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Error("3" + err.Error())
	}

	languageMap := make(map[string]Language)

	for k, v := range result.Languages {
		var l Language
		l.direction = v["dir"]
		l.name = v["name"]
		l.nativeName = v["nativeName"]
		languageMap[k] = l
	}

	return languageMap

}

func msAnyToLanguage(text string, language string) SingleTranslation {
	key := os.Getenv("WB_MS_TRANSLATE_KEY")
	region := os.Getenv("WB_MS_TRANSLATE_REGION")
	reqUrl, _ := url.Parse(endpoint + "/translate")
	q := reqUrl.Query()
	q.Add("api-version", "3.0")
	q.Add("to", language)
	reqUrl.RawQuery = q.Encode()

	body := []struct {
		Text string
	}{
		{Text: text},
	}
	reqBody, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", reqUrl.String(), bytes.NewBuffer(reqBody))
	if err != nil {
		log.Error("1" + err.Error())
	}

	req.Header.Add("Ocp-Apim-Subscription-Key", key)
	req.Header.Add("Ocp-Apim-Subscription-Region", region)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("2" + err.Error())
	}

	var result []TranslationResult
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Error("3" + err.Error())
	}

	finalTranslation := SingleTranslation{
		translatedText: result[0].Translations[0].Text,
		fromLang:       msGetLanguageByCode(result[0].DetectedLanguage.Name)[0],
		fromLangNative: msGetLanguageByCode(result[0].DetectedLanguage.Name)[1],
		toLang:         msGetLanguageByCode(result[0].Translations[0].ToLanguage)[0],
		toLangNative:   msGetLanguageByCode(result[0].Translations[0].ToLanguage)[1],
		originalText:   text,
	}

	return finalTranslation

}
