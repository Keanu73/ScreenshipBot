package modules

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func CheckCharacterLimit(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Delete messages longer than 250 characters
	if len(m.Content) > 250 {
		// if message is longer than 250 characters, time them out
		initialTimer := m.Timestamp.Add(5 * time.Minute)

		err := s.GuildMemberTimeout(m.GuildID, m.Author.ID, &initialTimer)
		if err != nil {
			return
		}

		// send ephemeral reply
		_, err = s.ChannelMessageSend(
			m.Author.ID,
			":negative_squared_cross_mark: Your message was deleted because it was over 250 characters.\n"+
				"Please try and keep messages short, as we are a **CALL-focused** server.\n"+
				"You have 5 minutes to edit the length of your message. If you do not, you will be timed out for an hour.",
		)
		if err != nil {
			return
		}
	}
}
