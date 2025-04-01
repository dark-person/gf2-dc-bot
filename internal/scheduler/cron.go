package scheduler

import (
	"fmt"

	"github.com/dark-person/gf2-dc-bot/pkg/api"
)

// Calculate next date, include ranged date & cycled event.
func (s *Scheduler) calcNext() error {
	// Recalculate ranged schedule date
	err := s.cal.CalcNextRangedDate()
	if err != nil {
		return fmt.Errorf("error when calculating next ranged date: %v", err)
	}

	// Recalculate next deadline of cycled event
	err = s.cal.CalcNextDeadline()
	if err != nil {
		return fmt.Errorf("error when calculating next deadline: %v", err)
	}

	return nil
}

// Init cron task that work daily.
func (s *Scheduler) AddDailyCron(bot api.DiscordBot) {
	// Recalculate dates when startup
	err := s.calcNext()
	if err != nil {
		fmt.Println(currentTimeStr(), "Error calculating next date in startup: ", err)
		return
	}
	fmt.Println(currentTimeStr(), "[Startup] Next date updated.")

	// Recalculate ranged schedule date when every day start
	s.c.AddFunc("0 0 * * *", func() {
		err := s.calcNext()
		if err != nil {
			fmt.Println(currentTimeStr(), "Error calculating next ranged date:", err)
			return
		}
		fmt.Println(currentTimeStr(), "[Daily] Next date updated.")
	})

	// Send message at 22:00 of computer
	s.c.AddFunc("0 22 * * *", func() {
		fmt.Println(currentTimeStr(), "[Daily] Start send daily reminder message.")

		// Send discord message by combined all
		err := bot.SendReminder()
		if err != nil {
			fmt.Println(currentTimeStr(), "Error sending discord message:", err)
			return
		}
		fmt.Println(currentTimeStr(), "[Daily] Message sent.")
	})
}
