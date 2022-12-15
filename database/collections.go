package database

import (
	"log"
	"os"
	"time"

	"github.com/Keanu73/botutils"
	"go.mongodb.org/mongo-driver/mongo"
)

// VoiceChannelRecord represents a record of a user's time spent in a voice channel.
type VoiceChannelRecord struct {
	UserID     string    `bson:"user_id"`
	ChannelID  string    `bson:"channel_id"`
	StartTime  time.Time `bson:"start_time"`
	EndTime    time.Time `bson:"end_time"`
	Duration   time.Duration
	RecordedAt time.Time `bson:"recorded_at"`
}

var VCRCollection *mongo.Collection

func init() {
	client, err := botutils.Database.NewClient(os.Getenv("MONGODB_URI"))
	if err != nil {
		log.Fatal("[MONGO] Failed to connect to database:", err)
	}

	VCRCollection = client.Database("screenship").Collection("voice_channel_records")
}
