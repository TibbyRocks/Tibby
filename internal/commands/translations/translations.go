package translations

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/mozoarella/wombot/internal/utils"
)

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

	tl := msAnyToLanguage(msg.Content, "en")

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
