package translations

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	googletranslate "github.com/tibbyrocks/tibby/internal/commands/translations/GoogleTranslate"
	googletranslatev3 "github.com/tibbyrocks/tibby/internal/commands/translations/GoogleTranslateV3"
	microsofttranslate "github.com/tibbyrocks/tibby/internal/commands/translations/MicrosoftTranslate"
	"github.com/tibbyrocks/tibby/internal/types"
	"github.com/tibbyrocks/tibby/internal/utils"
)

/*
	As it stands I am using Google translations and I use the Microsoft Translations APIs for getting a language's name
	Google only offers up translated names for languages in one language at a time (specified with a parameter)
	rather than Microsoft offering both the name in English and the native one.

	The Microsoft Translations API also doesn't require credentials for getting the language names
	but I keep them in place in case I have to switch translation backends.
*/

var (
	Translators map[string]types.Translator = make(map[string]types.Translator)
)

func init() {
	Translators["microsoft"] = microsofttranslate.Translator
	Translators["google"] = googletranslate.Translator
	Translators["googlev3"] = googletranslatev3.Translator
}

func MsgTranslationToEnglish(i *discordgo.InteractionCreate) string {
	msg := i.ApplicationCommandData().Resolved.Messages[i.ApplicationCommandData().TargetID]
	msgUrl := utils.ReturnInteractionMessageUrl(i)

	if len(msg.Embeds) > 0 {
		return fmt.Sprintf("I can't translate embeds, sorry.\n\n[Go to the original message](%v)", msgUrl)
	}

	if len(msg.Content) < 1 {
		return fmt.Sprintf("I can't translate messages without text, sorry.\n\n[Go to the original message](%v)", msgUrl)
	}

	var erroredTranslations string = `Failed to translate
	
	[Go to the original message](%v)
	`

	translation, err := Translators[os.Getenv("WB_TRANSLATOR")].Translate("", "en", msg.Content)
	if err != nil {
		return fmt.Sprintf(erroredTranslations, msgUrl)
	}

	translationWithLanguages := Translators[os.Getenv("WB_LANGUAGELOOKUP")].FillLanguagesFromCodes(translation)

	return buildMessage(translationWithLanguages, msgUrl)
}

func buildMessage(translation types.SingleTranslation, msgURL string) string {
	var translationFormatString string = `**%v (%v)**
%v

**%v (%v)**
%v

[Go to the original message](%v)
`
	builtMessage := fmt.Sprintf(translationFormatString, translation.FromLang, translation.FromLangNative, translation.OriginalText, translation.ToLang, translation.ToLangNative, translation.TranslatedText, msgURL)

	return builtMessage
}
