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

func confessor(s *discordgo.Session, i *discordgo.InteractionCreate) {
	channel, err := s.Channel(i.ChannelID)
	if err != nil {
		log.Printf("Error fetching channel: %v", err)
		return
	}

	isDM := channel.Type == discordgo.ChannelTypeDM

	// Construct the custom ID based on the context
	var customID string
	if isDM {
		customID = "confession_form_dm_" + i.Interaction.User.ID // Use User ID for DMs
	} else {
		customID = "confession_form_" + i.Interaction.Member.User.ID // Use Member ID for guild
	}
	
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			Title:    "Confession Form",
			CustomID: customID,
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
							Label:       "Your Confession",
							Style:       discordgo.TextInputParagraph,
							Placeholder: "Enter your confession here...",
							Required:    true,
						},
					},
				},
			},
		},
	})
}

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"confess": confessor,
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
			onInteractionModal(s, i)

		case discordgo.InteractionMessageComponent:
			onInteractionClick(s, i)
		}
	})
}
