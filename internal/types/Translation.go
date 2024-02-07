package types

type Translator struct {
	Translate              func(fromLang string, toLang string, translatable string) SingleTranslation
	FillLanguagesFromCodes func(SingleTranslation) SingleTranslation
}

type SingleTranslation struct {
	FromLangCode   string
	FromLang       string
	FromLangNative string
	ToLangCode     string
	ToLang         string
	ToLangNative   string
	TranslatedText string
	OriginalText   string
}
