package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/FedorLap2006/disgolf"
	"github.com/Keanu73/ScreenshipBot/handlers"
	"github.com/Keanu73/botutils"
	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
)

var envVars = []string{
	"ADMIN_ROLE_ID",
	"BOT_TOKEN",
	"BOT_GUILD_ID",
	"DATABASE_URI",
}

func main() {
	// Checks if env variables were set

	botutils.Env.Check(envVars)

	// If all good... let's continue

	bot, err := disgolf.New(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(fmt.Errorf("failed to initialise session: %w", err))
	}

	// Adds disgolf command interaction handler & ready event
	bot.AddHandler(bot.Router.HandleInteraction)
	bot.AddHandler(
		func(session *discordgo.Session, ready *discordgo.Ready) {
			log.Printf("[READY] %s#%s is online!", ready.User.Username, ready.User.Discriminator)
		},
	)

	bot.AddHandler(handlers.MessageCreate)
	bot.AddHandler(handlers.VoiceStateUpdate)

	// Opens bot's session using token
	err = bot.Open()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to open session: %w", err))
	}

	// Syncs disgolf router with discordgo bot
	err = bot.Router.Sync(bot.Session, "", os.Getenv("BOT_GUILD_ID"))
	if err != nil {
		log.Fatal(fmt.Errorf("failed to sync commands: %w", err))
	}

	// Updates bot status
	_ = bot.Session.UpdateStatusComplex(
		discordgo.UpdateStatusData{
			Activities: []*discordgo.Activity{
				{
					Name: "people use screens less",
					Type: discordgo.ActivityTypeWatching,
				},
			},
		},
	)

	// Schedules timekeeping & VC access crons
	// modules.Timekeeping.ScheduleCrons(bot.Session)
	// modules.VCAccess.ScheduleCrons(bot.Session)

	// Allows for graceful Ctrl + C
	ech := make(chan os.Signal)
	signal.Notify(ech, os.Kill, syscall.SIGTERM) //nolint:govet,staticcheck
	<-ech
	_ = bot.Close()
}
