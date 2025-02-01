package scheduler

import (
	"fmt"
	"time"

	"github.com/dark-person/gf2-dc-bot/internal/discord"
	"github.com/dark-person/lazydb"
	"github.com/robfig/cron/v3"
)

// Scheduler struct to manage cron jobs.
type Scheduler struct {
	c  *cron.Cron     // Root cron job
	db *lazydb.LazyDB // Database connection

	bot discord.Bot // Discord bot manager
}

// Create a new scheduler instance, which without discord bot settings.
// If the discord bot is not available, then this scheduler will not send any message to discord.
func NewScheduler(c *cron.Cron, db *lazydb.LazyDB) *Scheduler {
	return &Scheduler{c: c, db: db, bot: nil}
}

// Set the discord bot.
func (s *Scheduler) SetBot(bm discord.Bot) {
	if bm == nil {
		panic("You must provide a bot reference for scheduler creation")
	}

	// Set bot
	s.bot = bm
}

// Send a message of reminder, through discord channel.
// If the bot is not available, then this function has no effect.
func (s *Scheduler) sendReminder(message string) error {
	if s.bot == nil {
		fmt.Println("Warning: discord bot is not available")
		return nil
	}

	return s.bot.SendReminder(message)
}

// Get current time in opinionated formatted string.
func currentTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
