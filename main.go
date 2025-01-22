package main

import (
	"fmt"
	"time"

	"github.com/dark-person/gf2-dc-bot/internal/config"
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
		"Config loaded. Token: ", cfg.Token, "Channels: ", cfg.ChannelID)

	// Init task
	initDailyTask()

	// Keep the main program running
	select {}
}
