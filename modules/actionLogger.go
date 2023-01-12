package modules

import (
	"context"
	"time"

	"github.com/Keanu73/ScreenshipBot/database"
	"github.com/bwmarrin/discordgo"
)

// Store a map of voice channel join times for each user.
var joinTimes = make(map[string]time.Time)

// Sends a message: <USERNAME> meditated/exercised/journaled] for <DURATION> with <MENTIONS> in <CHANNEL>!
// Formats it in embed?

func ActionLogger(_ *discordgo.Session, event *discordgo.VoiceStateUpdate) {
	oldChannelID := event.BeforeUpdate.ChannelID
	newChannelID := event.ChannelID

	// Get a handle to the voice channel records collection.
	collection := &database.VCRCollection

	// If the user just joined the voice channel, store the current time.
	if oldChannelID == "" && newChannelID != "" {
		joinTimes[event.UserID] = time.Now()
		return
	}

	// If the user just left a voice channel, calculate the duration they were in the channel and store it in a VoiceChannelRecord.
	if oldChannelID != "" && newChannelID == "" {
		joinTime, ok := joinTimes[event.UserID]
		if !ok {
			return
		}
		delete(joinTimes, event.UserID)

		duration := time.Since(joinTime)

		record := &database.VoiceChannelRecord{
			UserID:     event.UserID,
			ChannelID:  event.ChannelID,
			StartTime:  joinTime,
			EndTime:    joinTime.Add(duration),
			Duration:   duration,
			RecordedAt: time.Now(),
		}

		// Insert the record into the collection.
		_, err := (*collection).InsertOne(context.TODO(), record)
		if err != nil {
			// Handle error.
			return
		}
	}
}
