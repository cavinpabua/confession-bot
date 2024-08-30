package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func onInteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
		_, err := s.ChannelMessageSendEmbed(i.ChannelID, embed)
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
