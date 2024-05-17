package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	_ "github.com/tibbyrocks/tibby/internal/autoload"
	"github.com/tibbyrocks/tibby/internal/commands"
	"github.com/tibbyrocks/tibby/internal/commands/magic8ball"
	"github.com/tibbyrocks/tibby/internal/commands/radlibs"
	"github.com/tibbyrocks/tibby/internal/commands/sorrygenerator"
	"github.com/tibbyrocks/tibby/internal/commands/tibbycmds"
	"github.com/tibbyrocks/tibby/internal/commands/translations"
	"github.com/tibbyrocks/tibby/internal/commands/wisdom"
	"github.com/tibbyrocks/tibby/internal/utils"
)

var (
	GuildID            = ""
	log                = utils.Log
	unregisterCommands = flag.Bool("unregister", false, "Use this flag to unregister all registered bot commands")
	debugMode          = flag.Bool("debug", false, "Use this flag to enable debug mode")
	customs            = utils.GetCustoms()
	dc                 *discordgo.Session
	err                error
)

var (
	botCommandSlices []*[]commands.Command = []*[]commands.Command{
		&tibbycmds.Commands,
		&magic8ball.Commands,
		&radlibs.Commands,
		&translations.Commands,
		&wisdom.Commands,
		&sorrygenerator.Commands,
	}
	botCommands, interactionMap = splitBotCommands(botCommandSlices)
)

func init() {
	flag.Parse()
	if *debugMode {
		utils.LogLevel.Set(slog.LevelDebug)
	}
}

func splitBotCommands(cmdSlices []*[]commands.Command) (map[string]commands.Command, map[string]commands.Command) {
	var commandSlice []commands.Command
	var commandMap map[string]commands.Command = make(map[string]commands.Command)
	var interactionMap map[string]commands.Command = make(map[string]commands.Command)
	for _, v := range cmdSlices {
		commandSlice = append(commandSlice, *v...)
	}
	for _, v := range commandSlice {
		commandMap[v.AppComm.Name] = v
		if len(v.InteractionIDPrefixes) > 0 {
			for _, idPrefix := range v.InteractionIDPrefixes {
				interactionMap[idPrefix] = v
			}
		}
	}
	return commandMap, interactionMap
}

func main() {
	log.Info("Starting " + customs.BotName)

	tibbycmds.RegisterAppStart()
	setupDiscordSession()
	addDiscordHandlers()
	openDiscordConnection()

	registerDiscordCommands()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	if *unregisterCommands {
		unregisterDiscordCommands()
	}

	closeDiscordConnection()
}

func setupDiscordSession() {
	dc, err = discordgo.New("Bot " + os.Getenv("WB_DC_TOKEN"))
	if err != nil {
		log.Error("Couldn't set up the Discord session", err)
		return
	}
	dc.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentGuildMembers
}

func addDiscordHandlers() {
	dc.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if h, ok := botCommands[i.ApplicationCommandData().Name]; ok {
				h.Handler(s, i)
			}
		case discordgo.InteractionModalSubmit:
			if h, ok := interactionMap[utils.GetCustomIDPrefix(i.ModalSubmitData().CustomID)]; ok {
				h.Handler(s, i)
			}
		case discordgo.InteractionMessageComponent:
			if h, ok := interactionMap[utils.GetCustomIDPrefix(i.MessageComponentData().CustomID)]; ok {
				h.Handler(s, i)
			}
		}

	})
}

func openDiscordConnection() {
	err = dc.Open()
	if err != nil {
		log.Error("Error opening connection", err)
		return
	}
	log.Info(fmt.Sprintf("%s is running with the username '%s' and ID '%s'", customs.BotName, dc.State.User.Username, dc.State.User.ID))
}

func registerDiscordCommands() {
	log.Info("Registering commands with the Discord API")
	for _, v := range botCommands {
		_, err := dc.ApplicationCommandCreate(dc.State.User.ID, "", v.AppComm)
		if err != nil {
			log.Error(fmt.Sprintf("Cannot create '%s' command: %s", v.AppComm.Name, err))
		}
	}

}

func unregisterDiscordCommands() {
	log.Info("Unregistering commands...")
	commandsToRemove, err := dc.ApplicationCommands(dc.State.User.ID, "")
	if err != nil {
		log.Error("Could not get registered commands")
	}
	for _, c := range commandsToRemove {
		err := dc.ApplicationCommandDelete(dc.State.User.ID, "", c.ID)
		if err != nil {
			log.Error(fmt.Sprintf("Cannot delete '%s' command: %s", c.Name, err))
		}
	}
}

func closeDiscordConnection() {
	log.Info("Gracefully shutting down")
	err = dc.Close()
	if err != nil {
		log.Error(err.Error())
	}
}
