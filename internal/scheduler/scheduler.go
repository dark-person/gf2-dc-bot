package scheduler

import (
	"time"

	"github.com/dark-person/gf2-dc-bot/internal/calendar"
	"github.com/dark-person/lazydb"
	"github.com/robfig/cron/v3"
)

// Scheduler struct to manage cron jobs.
type Scheduler struct {
	c   *cron.Cron         // Root cron job
	db  *lazydb.LazyDB     // Database connection
	cal *calendar.Calendar // Calendar for events
}

// Create a new scheduler instance, which without discord bot settings.
// If the discord bot is not available, then this scheduler will not send any message to discord.
func NewScheduler(c *cron.Cron, db *lazydb.LazyDB) *Scheduler {
	return &Scheduler{c: c, db: db, cal: calendar.NewCalendar(db)}
}

// Get current time in opinionated formatted string.
func currentTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
