package sorrygenerator

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tibbyrocks/tibby/internal/commands"
	"github.com/tibbyrocks/tibby/internal/utils"
)

var (
	Commands []commands.Command
)

func init() {
	Commands = append(Commands, commands.Command{
		AppComm:               &SorryGeneratorCommand,
		Handler:               SorryGeneratorCommandHandler,
		InteractionIDPrefixes: []string{"sorry_theme"},
	})
}

var SorryGeneratorCommand = discordgo.ApplicationCommand{
	Name:        "sorry",
	Description: "Generate an apology",
}

func SorryGeneratorCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "Select a theme for your apology image",
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.SelectMenu{

								CustomID: "sorry_theme-",
								MenuType: discordgo.StringSelectMenu,
								Options: []discordgo.SelectMenuOption{
									{
										Label: "Cyberpunk",
										Value: "cyberpunk",
									},
								},
							},
						},
					},
				},
			},
		})
	case discordgo.InteractionMessageComponent:
		switch utils.GetCustomIDPrefix(i.MessageComponentData().CustomID) {
		case "sorry_theme":
			respondWithLatexImage(s, i)
		}
	}

}

func respondWithLatexImage(s *discordgo.Session, i *discordgo.InteractionCreate) {
	//textToLatexify := i.ModalSubmitData().Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
	//textToLatexify := i.MessageComponentData().Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
	textToLatexify := "Helloooo"
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: textToLatexify,
		},
	})
}
