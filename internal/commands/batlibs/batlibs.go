package batlibs

import (
	"strings"

	"github.com/mozoarella/wombot/internal/commands"
	"github.com/mozoarella/wombot/internal/types"
)

type Randomizer = types.Randomizer

var (
	singularNouns    Randomizer
	pluralNouns      Randomizer
	verbs            Randomizer
	adjectives       Randomizer
	thirdPersonVerbs Randomizer
	animals          Randomizer
	fruit            Randomizer
)

func init() {
	commands.BotCommands["batlibs"] = &commands.Command{
		Name: "Batlibs",
		Help: "Does Batlibs",
	}
}

func init() {
	pluralNouns.Append("foxes", "dogs", "cats", "houses", "men")
	verbs.Fill("./data/verbs.txt", true)
	adjectives.Fill("./data/adjectives.txt", true)
	singularNouns.Fill("./data/singularnouns.txt", true)
	animals.Fill("./data/animals.txt", true)
	fruit.Fill("./data/fruit.txt", true)
	singularNouns.Combine(&animals, &fruit)

	//adjectives.Append("blue", "cute", "wet", "gassy")
	thirdPersonVerbs.Append("walks", "jumps", "cooks", "drives", "swims")
}

// This function splits up the message string and loops over every string
func DoBatlibs(message string) string {
	splitMessage := strings.Split(message, " ")
	var workSlice []string
	for _, s := range splitMessage {
		workSlice = append(workSlice, replaceBatlibToken(s))
	}
	return strings.Join(workSlice, " ")
}

func replaceBatlibToken(token string) string {
	token = strings.ReplaceAll(token, "$ANIMAL", animals.Random())
	token = strings.ReplaceAll(token, "$FRUIT", fruit.Random())
	token = strings.ReplaceAll(token, "$NOUNS", pluralNouns.Random())
	token = strings.ReplaceAll(token, "$NOUN", singularNouns.Random())
	token = strings.ReplaceAll(token, "$VERBS", thirdPersonVerbs.Random())
	token = strings.ReplaceAll(token, "$VERB", verbs.Random())
	token = strings.ReplaceAll(token, "$ADJ", adjectives.Random())
	return token
}
