package handlers

import (
	"github.com/Keanu73/ScreenshipBot/modules"
	"github.com/bwmarrin/discordgo"
)

func VoiceStateUpdate(session *discordgo.Session, event *discordgo.VoiceStateUpdate) {
	modules.ActionLogger(session, event)
}
