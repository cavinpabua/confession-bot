package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var GuildID = ""

func checkNilErr(e error) {
	if e != nil {
		log.Fatalf("Error message '%v'", e)
	}
}

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "confess",
		Description: "Add a confession",
	},
}

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"confess": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseModal,
			Data: &discordgo.InteractionResponseData{
				Title:    "Confession Form",
				CustomID: "confession_form_" + i.Interaction.Member.User.ID,
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID:    "name",
								Label:       "Hidden Name",
								Style:       discordgo.TextInputShort,
								Placeholder: "Boy Tapang...",
								Required:    true,
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID:    "confession",
								Label:       "Your Conession",
								Style:       discordgo.TextInputParagraph,
								Placeholder: "Enter your conession here...",
								Required:    true,
							},
						},
					},
				},
			},
		})
	},
}

func registerCommands(discord *discordgo.Session) {
	for _, cmd := range Commands {
		_, err := discord.ApplicationCommandCreate(discord.State.User.ID, GuildID, cmd)
		checkNilErr(err)
	}

	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
				h(s, i)
			}
		case discordgo.InteractionModalSubmit:
			onInteractionCreate(s, i)
		}
	})
}
