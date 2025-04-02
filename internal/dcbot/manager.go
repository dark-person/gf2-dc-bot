package dcbot

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/dark-person/gf2-dc-bot/internal/calendar"
	"github.com/dark-person/gf2-dc-bot/internal/config"
	"github.com/dark-person/gf2-dc-bot/pkg/api"
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
	cfg     *config.DiscordConfig // Original configuration
	cal     *calendar.Calendar    // Calendar for calculation
	persona api.BotPersona        // Discord bot personality control

	initialized bool               // Only true when this manager is initialized
	session     *discordgo.Session // Discord session that designed for notification

	ReminderChannel string // Channel ID for daily reminder notification
}

// Interface check
var _ api.DiscordBot = (*BotManager)(nil)

// Create a new empty discord bot manager, with WA2000 persona chosen.
func NewManager(persona api.BotPersona) *BotManager {
	return &BotManager{
		cfg:             nil,
		initialized:     false,
		session:         nil,
		persona:         persona,
		ReminderChannel: "",
	}
}

// Init this bot manager with given configuration,
// which also validate the configuration is able to run or not.
func (bm *BotManager) Init(cfg *config.DiscordConfig, cal *calendar.Calendar) error {
	// Perform validation of the configuration
	if cfg.Token == "" || cfg.ReminderChannel == "" {
		return fmt.Errorf("discord token or channel ID not set")
	}

	bm.cfg = cfg // Store original configuration for future reference
	bm.ReminderChannel = cfg.ReminderChannel

	var err error

	// Create a discord connection session
	bm.session, err = discordgo.New("Bot " + cfg.Token)
	if err != nil {
		return fmt.Errorf("failed to create discord bot: %v", err)
	}

	// Add interaction listener for message
	bm.session.Identify.Intents |= discordgo.IntentMessageContent
	bm.session.AddHandler(bm.messageCreate)

	err = bm.session.Open()
	if err != nil {
		return fmt.Errorf("failed to open discord connection: %v", err)
	}

	// Set calendar object
	bm.cal = cal

	bm.initialized = true
	return nil
}

// Get current time in opinionated formatted string.
func currentTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
