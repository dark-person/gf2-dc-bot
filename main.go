package main

import (
	"github.com/dark-person/gf2-dc-bot/internal/calendar"
	"github.com/dark-person/gf2-dc-bot/internal/config"
	"github.com/dark-person/gf2-dc-bot/internal/dcbot"
	"github.com/dark-person/gf2-dc-bot/internal/scheduler"
	"github.com/dark-person/gf2-dc-bot/pkg/persona/zh/wa2000"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var cfg *config.DiscordConfig

func main() {
	var err error

	// Start up log
	setupLogger(zerolog.TraceLevel)
	log.Info().Msg("Starting up...")

	// Load config
	cfg, err = config.LoadYaml("config.yaml")
	if err != nil {
		panic(err) // Program will never run properly when config not loaded
	}

	log.Debug().
		Str("Token", cfg.Token).
		Str("Channels", cfg.ReminderChannel).
		Str("Role", cfg.GunSmokeRemindRole).
		Msg("Config loaded.")

	// Init discord
	bot := dcbot.NewManager(wa2000.New())

	// Setup database
	err = setup()
	if err != nil {
		panic(err) // Program will never run properly when database init fails
	}
	log.Debug().Msg("Database ready.")

	// Setup calendar
	cal := calendar.NewCalendar(db)

	// Init cron jobs
	c := cron.New()
	s := scheduler.NewScheduler(c, cal, bot)
	s.AddDailyCron(bot)

	// Start discord bot
	err = bot.Init(cfg, cal)
	if err != nil {
		panic(err) // Program will never run properly when discord init fails
	}

	// Start cron job
	c.Start()
	log.Info().Msg("All service ready.")

	// Keep the main program running
	select {}
}
