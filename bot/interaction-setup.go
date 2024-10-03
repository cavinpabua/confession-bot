package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func onInteractionClick(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.MessageComponentData().CustomID == "submit_confession" {
		confessor(s, i)
	}
}

func onInteractionModal(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ModalSubmitData().CustomID == "confession_form_"+i.Interaction.Member.User.ID {
		name := i.ModalSubmitData().Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
		confession := i.ModalSubmitData().Components[1].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value

		embed := &discordgo.MessageEmbed{
			Title:       "Anonymous Confession",
			Description: "@here is a confession by **" + name + "** 💌💌💌",
			Color:       0x00ff00,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Confession",
					Value:  confession,
					Inline: false,
				},
			},
		}
		var receipientId = ""
		if i.ChannelID == "" {
			receipientId = i.User.ID
		} else {
			receipientId = i.ChannelID
		}
		_, err := s.ChannelMessageSendComplex(receipientId, &discordgo.MessageSend{
			Embed: embed,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Submit Confession",
							Style:    discordgo.PrimaryButton,
							CustomID: "submit_confession",
						},
					},
				},
			},
		})
		if err != nil {
			log.Printf("Error sending embed: %v", err)
		}

		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseDeferredMessageUpdate,
		})
		if err != nil {
			log.Printf("Error acknowledging interaction: %v", err)
		}
	}
}
