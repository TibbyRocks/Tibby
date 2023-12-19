package translations

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/mozoarella/wombot/internal/utils"
)

/*
	As it stands I am using Google translations and I use the Microsoft Translations APIs for getting a language's name
	Google only offers up translated names for languages in one language at a time (specified with a parameter)
	rather than Microsoft offering both the name in English and the native one.

	The Microsoft Translations API also doesn't require credentials for getting the language names
	but I keep them in place in case I have to switch translation backends.
*/

type SingleTranslation struct {
	fromLang       string
	fromLangNative string
	toLang         string
	toLangNative   string
	translatedText string
	originalText   string
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

	//tl := msAnyToLanguage(msg.Content, "en")
	tl := gtAnyToLanguage(msg.Content, "en")

	return buildMessage(tl, msgUrl)
}

func buildMessage(translation SingleTranslation, msgURL string) string {
	var translationFormatString string = `**%v (%v)**
%v

**%v (%v)**
%v

[Go to the original message](%v)
`
	builtMessage := fmt.Sprintf(translationFormatString, translation.fromLang, translation.fromLangNative, translation.originalText, translation.toLang, translation.toLangNative, translation.translatedText, msgURL)

	return builtMessage
}
