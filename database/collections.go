package database

import (
	"log"
	"os"
	"time"

	"github.com/Keanu73/botutils"
	"go.mongodb.org/mongo-driver/mongo"
)

type Settings struct {
	ActionLoggingChannelID string `bson:"action_logging_channel_id"`
	AutoKickDuration       string `bson:"channel_id"`
	CharacterLimit         int    `bson:"character_limit"`
}

type VoiceChannelReservation struct {
	UserID    string    `bson:"user_id"`
	ChannelID string    `bson:"channel_id"`
	StartTime time.Time `bson:"start_time"`
	EndTime   time.Time `bson:"end_time"`
}

var VCRCollection *mongo.Collection

func init() {
	client, err := botutils.Database.NewClient(os.Getenv("MONGODB_URI"))
	if err != nil {
		log.Fatal("[MONGO] Failed to connect to database:", err)
	}

	VCRCollection = client.Database("screenship").Collection("voice_channel_records")
}
