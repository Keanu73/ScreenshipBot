package database

import (
	"context"
)

func ModifySetting(setting string, value interface{}) error {
	_, err := db.Exec(context.Background(), "UPDATE 'settings' SET value = $1 WHERE setting = $2", value, setting)
	return err
}
