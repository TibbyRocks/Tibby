package magic8ball

import (
	"fmt"
	"math/rand"

	"github.com/bwmarrin/discordgo"
	"github.com/tibbyrocks/tibby/internal/commands"
	"github.com/tibbyrocks/tibby/internal/types"
	"github.com/tibbyrocks/tibby/internal/utils"
)

type Randomizer = types.Randomizer

var (
	yesResponses Randomizer
	noResponses  Randomizer
	Commands     []commands.Command
	customs      = utils.GetCustoms()
)

var EightBallCommand = discordgo.ApplicationCommand{
	Name:        "8ball",
	Description: "Shake a magic 8-ball",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "question",
			Description: "Optional question for the 8-ball",
			Required:    false,
		},
	},
}

func init() {
	yesResponses.Fill("customizations/8ball/yes.txt", true)
	noResponses.Fill("customizations/8ball/no.txt", true)
	Commands = append(Commands, commands.Command{
		AppComm: &EightBallCommand,
		Handler: EightBallCommandHandler,
	})
}

func EightBallCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	ballResponse := ShakeTheBall(i)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Description: ballResponse,
					Author: &discordgo.MessageEmbedAuthor{
						Name:    fmt.Sprintf("%s Magic 8-Ball", customs.BotName),
						IconURL: utils.GetCdnUri("8ball-icon"),
					},
				},
			},
		},
	})
}

func ShakeTheBall(i *discordgo.InteractionCreate) string {
	utils.LogCmd(i)
	optionMap := utils.GetOptionsFromInteraction(i)

	var shaker string
	var question string
	var ballResponse string

	randNum := rand.Intn(100) + 1

	if randNum <= 50 {
		ballResponse = yesResponses.Random()
	} else {
		ballResponse = noResponses.Random()
	}

	shaker = utils.GetNickFromInteraction(i)

	if val, ok := optionMap["question"]; ok {
		question = val.StringValue()
	}

	var fullResponse string

	var responseWithoutQ string = `*%v shakes the Magic 8-ball*
	
The Magic 8-ball says: 
**%v**`

	var responseWithQ string = `%v asks the Magic 8-ball: "%v"

*They give the ball a good shake*

The Magic 8-ball says: 
**%v**`

	if len(question) == 0 {
		fullResponse = fmt.Sprintf(responseWithoutQ, shaker, ballResponse)
	} else {
		fullResponse = fmt.Sprintf(responseWithQ, shaker, question, ballResponse)
	}

	return fullResponse
}
