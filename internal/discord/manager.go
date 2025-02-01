package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/dark-person/gf2-dc-bot/internal/config"
)

// Manager for control static functions reference of this discord package.
//
// Before using this manager, you should call Init function like this:
//
//	m := discord.NewManager()
//	m.Init(someConfig)
//
// Otherwise, this bot manager will never work properly.
type BotManager struct {
	cfg        *config.DiscordConfig // Original configuration
	initalized bool                  // Only true when this manager is initialized
	session    *discordgo.Session    // Discord session that designed for notification

	ReminderChannel string // Channel ID for daily reminder notification
}

// Create a new empty discord bot manager.
func NewManager() *BotManager {
	return &BotManager{
		cfg:             nil,
		initalized:      false,
		session:         nil,
		ReminderChannel: "",
	}
}

// Init this bot manager with given configuration,
// which also validate the configuration is able to run or not.
func (bm *BotManager) Init(cfg *config.DiscordConfig) error {
	// Perform validation of the configuration
	if cfg.Token == "" || cfg.ReminderChannel == "" {
		return fmt.Errorf("discord token or channel ID not set")
	}

	bm.cfg = cfg // Store original configuration for future reference
	bm.ReminderChannel = cfg.ReminderChannel

	var err error

	bm.session, err = discordgo.New("Bot " + cfg.Token)
	if err != nil {
		return fmt.Errorf("failed to create discord bot: %v", err)
	}

	err = bm.session.Open()
	if err != nil {
		return fmt.Errorf("failed to open discord connection: %v", err)
	}

	bm.initalized = true
	return nil
}
