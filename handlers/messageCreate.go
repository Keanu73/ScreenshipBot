package handlers

import (
	"github.com/Keanu73/ScreenshipBot/modules"
	"github.com/bwmarrin/discordgo"
)

// All handlers will call modules in some way, shape or form

func MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Checks if message is longer than 170 characters
	modules.CheckCharacterLimit(session, message)
}
