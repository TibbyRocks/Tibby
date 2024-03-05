package googletranslatev3

import (
	"context"
	"fmt"
	"os"

	translate "cloud.google.com/go/translate/apiv3"
	"cloud.google.com/go/translate/apiv3/translatepb"
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

func Translate(fromLang string, toLang string, translatable string) (types.SingleTranslation, error) {
	ctx := context.Background()
	client, err := translate.NewTranslationClient(ctx)
	if err != nil {
		log.Error(err.Error())
		var emptyResponse types.SingleTranslation
		return emptyResponse, err
	}
	defer client.Close()

	req := &translatepb.TranslateTextRequest{
		Parent:             fmt.Sprintf("projects/%s/locations/global", os.Getenv("GOOGLE_PROJECT")),
		SourceLanguageCode: fromLang,
		TargetLanguageCode: toLang,
		MimeType:           "text/plain",
		Contents:           []string{translatable},
	}

	resp, err := client.TranslateText(ctx, req)
	if err != nil {
		log.Error(err.Error())
		var emptyResponse types.SingleTranslation
		return emptyResponse, err
	}

	translations := resp.GetTranslations()

	finalTranslation := types.SingleTranslation{
		TranslatedText: translations[0].TranslatedText,
		FromLangCode:   translations[0].DetectedLanguageCode,
		ToLangCode:     toLang,
		OriginalText:   translatable,
	}
	return finalTranslation, nil
}
