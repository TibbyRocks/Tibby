package wisdom

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/bwmarrin/discordgo"
	"github.com/tibbyrocks/tibby/internal/commands"
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
	Commands  []commands.Command
	customs   = utils.GetCustoms()
)

func init() {
	Commands = append(Commands, commands.Command{
		AppComm: &WisdomCommand,
		Handler: WisdomCommandHandler,
	})
}

var WisdomCommand = discordgo.ApplicationCommand{
	Name:        "wisdom",
	Description: "Get a (random) quote",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "quoteid",
			Description: "quotable.io quote ID",
			Required:    false,
		},
	},
}

func WisdomCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	quoteMsg := GetQuote(i)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Description: quoteMsg,
					Author: &discordgo.MessageEmbedAuthor{
						Name:    fmt.Sprintf("%s Wisdom", customs.BotName),
						IconURL: s.State.User.AvatarURL("1024"),
					},
				},
			},
		},
	})
}

func getRandomQuotes(amount int) []quote {
	var quoteArray []quote

	reqUrl, _ := url.Parse(quotehost + "/quotes/random")
	q := reqUrl.Query()
	q.Add("limit", fmt.Sprint(amount))
	reqUrl.RawQuery = q.Encode()

	var failedQuote quote = quote{
		Id:      "404",
		Content: "Could not get a quote...",
		Author:  customs.BotName,
	}

	req, err := http.NewRequest("GET", reqUrl.String(), nil)
	if err != nil {
		log.Error(err.Error())
		quoteArray = append(quoteArray, failedQuote)
		return quoteArray
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error(err.Error())
		quoteArray = append(quoteArray, failedQuote)
		return quoteArray
	}

	if err := json.NewDecoder(res.Body).Decode(&quoteArray); err != nil {
		log.Error(err.Error())
		quoteArray = append(quoteArray, failedQuote)
		return quoteArray
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

	var failedQuote quote = quote{
		Id:      "404",
		Content: "Could not get a quote...",
		Author:  customs.BotName,
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error(err.Error())
		resolvedQuote = failedQuote
		return resolvedQuote
	}

	if res.StatusCode == 404 {
		resolvedQuote = quote{
			Content: "Four-Oh-Four",
			Author:  "The Server",
			Id:      "N0T-F0UND",
		}
	} else {
		if err := json.NewDecoder(res.Body).Decode(&resolvedQuote); err != nil {
			log.Error(err.Error())
			resolvedQuote = failedQuote
			return resolvedQuote
		}
	}

	return resolvedQuote
}

func GetQuote(i *discordgo.InteractionCreate) string {
	utils.LogCmd(i)
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
