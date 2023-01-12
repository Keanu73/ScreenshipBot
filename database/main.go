package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var db *pgx.Conn
var err error

func main() {
	db, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	if err := db.Ping(context.Background()); err != nil {
		defer db.Close(context.Background()) // close database connection
		log.Fatal(fmt.Errorf("error, not sent ping to database, %w", err))
	}
}
