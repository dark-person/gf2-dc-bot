package main

import (
	"fmt"
	"time"

	"github.com/dark-person/gf2-dc-bot/internal/calendar"
	"github.com/dark-person/gf2-dc-bot/internal/config"
	"github.com/dark-person/gf2-dc-bot/internal/dcbot"
	"github.com/dark-person/gf2-dc-bot/internal/scheduler"
	"github.com/dark-person/gf2-dc-bot/pkg/persona/zh/wa2000"
	"github.com/robfig/cron/v3"
)

// Get current time in opinionated formatted string.
func currentTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

var cfg *config.DiscordConfig

func main() {
	var err error

	// Start up log
	fmt.Println(currentTimeStr(), "Starting up...")

	// Load config
	cfg, err = config.LoadYaml("config.yaml")
	if err != nil {
		panic(err) // Program will never run properly when config not loaded
	}

	fmt.Println(currentTimeStr(),
		"Config loaded. Token: ", cfg.Token, "Channels: ", cfg.ReminderChannel)

	// Init discord
	bot := dcbot.NewManager(wa2000.New())

	// Setup database
	err = setup()
	if err != nil {
		panic(err) // Program will never run properly when database init fails
	}
	fmt.Println(currentTimeStr(), "Database ready.")

	// Setup calendar
	cal := calendar.NewCalendar(db)

	// Init cron jobs
	c := cron.New()
	s := scheduler.NewScheduler(c, cal, bot)
	s.AddDailyCron(bot)

	// Start discord bot
	bot.ReminderMsgGenerator = s.GetDailyReminderMsg
	err = bot.Init(cfg)
	if err != nil {
		panic(err) // Program will never run properly when discord init fails
	}

	// Start cron job
	c.Start()

	// Keep the main program running
	select {}
}
