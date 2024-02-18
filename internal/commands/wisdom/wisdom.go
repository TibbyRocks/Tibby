package wisdom

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/bwmarrin/discordgo"
	"github.com/tibbyrocks/tibby/internal/utils"
)

type quote struct {
	Id      string `json:"_id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

var (
	quotehost string = "https://api.quotable.io"
	log              = utils.Log
)

func getRandomQuotes(amount int) []quote {
	var quoteArray []quote

	reqUrl, _ := url.Parse(quotehost + "/quotes/random")
	q := reqUrl.Query()
	q.Add("limit", fmt.Sprint(amount))
	reqUrl.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", reqUrl.String(), nil)
	if err != nil {
		log.Error(err.Error())
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error(err.Error())
	}

	if err := json.NewDecoder(res.Body).Decode(&quoteArray); err != nil {
		log.Error(err.Error())
	}

	return quoteArray
}

func getQuoteByID(id string) quote {
	var resolvedQuote quote

	reqUrl, _ := url.Parse(quotehost + "/quotes/" + id)
	req, err := http.NewRequest("GET", reqUrl.String(), nil)
	if err != nil {
		log.Error(err.Error())
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error(err.Error())
	}

	if res.StatusCode == 404 {
		resolvedQuote.Content = "Four-Oh-Four"
		resolvedQuote.Author = "The Server"
		resolvedQuote.Id = "N0T-F0UND"
	} else {
		if err := json.NewDecoder(res.Body).Decode(&resolvedQuote); err != nil {
			log.Error(err.Error())
		}
	}

	return resolvedQuote
}

func GetQuote(i *discordgo.InteractionCreate) string {
	optionMap := utils.GetOptionsFromInteraction(i)
	var quoteID string
	var result quote

	if val, ok := optionMap["quoteid"]; ok {
		quoteID = val.StringValue()
	}

	if len(quoteID) == 0 {
		result = getRandomQuotes(1)[0]
	} else {
		result = getQuoteByID(quoteID)
	}

	return buildMessageSingle(result)
}

func buildMessageSingle(quote quote) string {
	var formatString string = `
			*%v*
		-%v

	ID: %v
`
	return fmt.Sprintf(formatString, quote.Content, quote.Author, quote.Id)
}
