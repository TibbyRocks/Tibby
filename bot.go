package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/tibbyrocks/tibby/internal/commands/magic8ball"
	"github.com/tibbyrocks/tibby/internal/commands/radlibs"
	"github.com/tibbyrocks/tibby/internal/commands/textmanipulation"
	"github.com/tibbyrocks/tibby/internal/commands/tibbycmds"
	"github.com/tibbyrocks/tibby/internal/commands/translations"
	"github.com/tibbyrocks/tibby/internal/commands/wisdom"
	"github.com/tibbyrocks/tibby/internal/types"
	"github.com/tibbyrocks/tibby/internal/utils"
)

type Randomizer = types.Randomizer

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
	botCommands = []*discordgo.ApplicationCommand{
		{
			Name:        "radlibs",
			Description: "Replaces certain tokens with words",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "msg",
					Description: "Message with the tokens you want to radlib",
					Required:    true,
				},
			},
		},
		{
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
		},
		{
			Name: "Translate to English",
			Type: discordgo.MessageApplicationCommand,
		},
		{
			Name: "Mockify",
			Type: discordgo.MessageApplicationCommand,
		},
		{
			Name: "UwUify",
			Type: discordgo.MessageApplicationCommand,
		},
		{
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
		},
		{
			Name:        customs.RootCommand,
			Description: fmt.Sprintf("General commands for %s", customs.BotName),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "docs",
					Description: fmt.Sprintf("Get the %s docs", customs.BotName),
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
				{
					Name:        "info",
					Description: fmt.Sprintf("Get the %s runtime information", customs.BotName),
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
			},
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"radlibs": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

			libbedMsg := radlibs.DoRadlibs(i)

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Description: libbedMsg,
							Author: &discordgo.MessageEmbedAuthor{
								Name:    fmt.Sprintf("%s RadLibs 😎", customs.BotName),
								IconURL: s.State.User.AvatarURL("1024"),
							},
						},
					},
				},
			})
		},
		"8ball": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

			log.Info(fmt.Sprintf("User '%s' called the 8ball command", utils.GetUsernameFromInteraction(i)))
			ballResponse := magic8ball.ShakeTheBall(i)

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
		},
		"Translate to English": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Description: translations.MsgTranslationToEnglish(i),
							Author: &discordgo.MessageEmbedAuthor{
								Name:    fmt.Sprintf("%s Translator", customs.BotName),
								IconURL: s.State.User.AvatarURL("1024"),
							},
						},
					},
				},
			})
		},
		"Mockify": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Description: textmanipulation.MockText(i),
							Author: &discordgo.MessageEmbedAuthor{
								Name:    "Mock",
								IconURL: s.State.User.AvatarURL("1024"),
							},
						},
					},
				},
			})
		},
		"UwUify": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Description: textmanipulation.Uwuify(i),
							Author: &discordgo.MessageEmbedAuthor{
								Name:    "UwUify",
								IconURL: s.State.User.AvatarURL("1024"),
							},
						},
					},
				},
			})
		},
		"wisdom": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

			quoteMsg := wisdom.GetQuote(i)

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
		},
		customs.RootCommand: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options

			switch options[0].Name {
			case "docs":
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Flags: discordgo.MessageFlagsEphemeral,
						Embeds: []*discordgo.MessageEmbed{
							{
								Description: fmt.Sprintf("[Read the %s docs here](%s)", customs.BotName, customs.DocsURL),
								Author: &discordgo.MessageEmbedAuthor{
									Name:    fmt.Sprintf("%s docs", customs.BotName),
									IconURL: s.State.User.AvatarURL("1024"),
								},
							},
						},
					},
				})
			case "info":
				infoMsg := tibbycmds.GetInfo(i, s)

				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Flags: discordgo.MessageFlagsEphemeral,
						Embeds: []*discordgo.MessageEmbed{
							{
								Color:       int(utils.HexToDec("BFEA7C")),
								Description: infoMsg,
								Author: &discordgo.MessageEmbedAuthor{
									Name:    fmt.Sprintf("%s Info", customs.BotName),
									IconURL: s.State.User.AvatarURL("1024"),
								},
							},
						},
					},
				})
			}
		},
	}
)

func init() {
	flag.Parse()
	if *debugMode {
		utils.LogLevel.Set(slog.LevelDebug)
	}
	err := godotenv.Load()
	if err != nil {
		log.Debug("Couldn't load .env, this is probably fine")
	} else {
		log.Debug("Loaded .env file(s)")
	}
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
	dc.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
}

func addDiscordHandlers() {
	dc.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
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
	registeredCommands := make([]*discordgo.ApplicationCommand, len(botCommands))
	for i, v := range botCommands {
		cmd, err := dc.ApplicationCommandCreate(dc.State.User.ID, "", v)
		if err != nil {
			log.Error(fmt.Sprintf("Cannot create '%s' command: %s", v.Name, err))
		}
		registeredCommands[i] = cmd
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
