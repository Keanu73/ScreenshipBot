package database

import (
	"math/big"
	"time"
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

type vcr struct{}

var VoiceChannelRecords = &vcr{}

// AddRecord query for creating a new user by given email and password hash.
func (v *vcr) AddRecord(record *VoiceChannelRecord) error {
	// Define query string.
	query := `INSERT INTO voice_channel VALUES ($1, $2, $3, $4, $5)`

	// Send query to database.
	_, err := q.Exec(
		query,
		u.ID, u.CreatedAt, u.UpdatedAt, u.Email, u.PasswordHash,
	)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// GetRecord gets a voice channel record by the user ID or channel ID.
func (v *vcr) GetRecord(user_id big.Int) error {
	// Define User variable.
	user := models.User{}

	// Define query string.
	query := `SELECT * FROM users WHERE email = $1`

	// Send query to database.
	err := q.Get(&user, query, email)
	if err != nil {
		// Return empty object and error.
		return user, err
	}

	// Return query result.
	return user, nil
}
