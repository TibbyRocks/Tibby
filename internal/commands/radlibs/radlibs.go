package radlibs

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/tibbyrocks/tibby/internal/commands"
	"github.com/tibbyrocks/tibby/internal/types"
	"github.com/tibbyrocks/tibby/internal/utils"
)

type Randomizer = types.Randomizer

var (
	singularNouns Randomizer
	pluralNouns   Randomizer
	verbs         Randomizer
	adjectives    Randomizer
	animals       Randomizer
	fruit         Randomizer
	gerunds       Randomizer
	adverbs       Randomizer
	log           = utils.Log
)

func init() {
	commands.BotCommands["radlibs"] = &commands.Command{
		Name: "Radlibs",
		Help: "Does Radlibs",
	}
}

func init() {
	pluralNouns.Fill("customizations/pluralnouns.txt", true)
	verbs.Fill("customizations/verbs.txt", true)
	adjectives.Fill("customizations/adjectives.txt", true)
	singularNouns.Fill("customizations/singularnouns.txt", true)
	animals.Fill("customizations/animals.txt", true)
	fruit.Fill("customizations/fruit.txt", true)
	singularNouns.Combine(&animals, &fruit)
	gerunds.Fill("customizations/gerunds.txt", true)
	adverbs.Fill("customizations/adverbs.txt", true)
}

// This function splits up the message string and loops over every string
func DoRadlibs(i *discordgo.InteractionCreate) string {
	optionMap := utils.GetOptionsFromInteraction(i)
	log.Info(fmt.Sprintf("User '%s' called the nounverb command with text '%s'", utils.GetUsernameFromInteraction(i), optionMap["msg"].StringValue()))
	message := optionMap["msg"].StringValue()
	splitMessage := strings.Split(message, " ")
	var workSlice []string
	for _, s := range splitMessage {
		workSlice = append(workSlice, replaceRadlibToken(s))
	}
	return strings.Join(workSlice, " ")
}

func replaceRadlibToken(token string) string {
	token = strings.ReplaceAll(token, "$ANIMAL", animals.Random())
	token = strings.ReplaceAll(token, "$FRUIT", fruit.Random())
	token = strings.ReplaceAll(token, "$NOUNS", pluralNouns.Random())
	token = strings.ReplaceAll(token, "$NOUN", singularNouns.Random())
	token = strings.ReplaceAll(token, "$VERBING", gerunds.Random())
	token = strings.ReplaceAll(token, "$VERB", verbs.Random())
	token = strings.ReplaceAll(token, "$ADJ", adjectives.Random())
	token = strings.ReplaceAll(token, "$ADVERB", adverbs.Random())
	return token
}
