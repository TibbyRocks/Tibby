package tibbycmds

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/tibbyrocks/tibby/internal/utils"
)

func getDocs(i *discordgo.InteractionCreate) string {
	utils.LogCmd(i)
	return fmt.Sprintf("[Read the %s docs here](%s)", customs.BotName, customs.DocsURL)
}
