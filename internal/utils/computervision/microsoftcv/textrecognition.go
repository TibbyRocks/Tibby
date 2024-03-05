package microsoftcv

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"os"

	"github.com/tibbyrocks/tibby/internal/utils"
	"github.com/tibbyrocks/tibby/internal/utils/computervision"
)

var (
	TextRecognizer computervision.TextRecognizer
	log            = utils.Log
)

func init() {
	TextRecognizer = computervision.TextRecognizer{
		RecognizeText: RecognizeText,
	}
}

func RecognizeText(imageURL string) (string, error) {
	endpoint := os.Getenv("WB_MS_CV_ENDPOINT")
	apikey := os.Getenv("WB_MS_CV_KEY")
	region := os.Getenv("WB_MS_TRANSLATE_REGION")

	reqUrl, _ := url.Parse(endpoint + "/computervision/imageanalysis:analyze")
	q := reqUrl.Query()
	q.Add("features", "read")
	q.Add("model-version", "latest")
	q.Add("api-version", "2024-02-01")

	body := []struct {
		url string
	}{
		{url: imageURL},
	}
	reqBody, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", reqUrl.String(), bytes.NewBuffer(reqBody))
	if err != nil {
		log.Error(err.Error())
		return "", err
	}

	req.Header.Add("Ocp-Apim-Subscription-Key", apikey)
	req.Header.Add("Ocp-Apim-Subscription-Region", region)

	return "", nil
}
