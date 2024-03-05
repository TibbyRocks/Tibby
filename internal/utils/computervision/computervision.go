package computervision

type TextRecognizer struct {
	RecognizeText func(imageURL string) (string, error)
}

func GetTextFromImage(imageURL string) (string, error) {
	return "", nil
}
