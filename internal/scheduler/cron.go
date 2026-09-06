package scheduler

import (
	"fmt"
	"time"

	"github.com/dark-person/gf2-dc-bot/pkg/api"
	"github.com/rs/zerolog/log"
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
		log.Error().Err(err).Msg("Error calculating next date in startup.")
		return
	}
	log.Debug().
		Str("Phase", "Startup").
		Msg("Next date updated.")

	// Recalculate ranged schedule date when every day start
	s.c.AddFunc("0 0 * * *", func() {
		err := s.calcNext()
		if err != nil {
			log.Error().Err(err).Msg("Error calculating next ranged date.")
			return
		}
		log.Debug().
			Str("Phase", "Daily").
			Msg("Next date updated.")
	})

	// Send message at 22:00 of computer
	s.c.AddFunc("0 22 * * *", func() {
		log.Debug().
			Str("Phase", "Daily").
			Msg("Start send daily reminder message.")

		// Send discord message by combined all
		err := bot.SendReminder()
		if err != nil {
			log.Error().Err(err).Msg("Error sending discord message:")
			return
		}
		log.Debug().
			Str("Phase", "Daily").
			Msg("Message sent.")
	})

	// Send message at 23:00 of computer
	s.c.AddFunc("0 23 * * *", func() {
		isGunSmoke, err := s.cal.IsGunSmokeFrontline(time.Now())
		if err != nil {
			log.Error().Err(err).Msg("Error when get database value.")
			return
		}

		log.Debug().
			Str("Phase", "Daily").
			Bool("isGunSmoke", isGunSmoke).
			Msg("Check if event message needed.")

		if !isGunSmoke {
			return
		}

		err = bot.SendGunSmokeReminder()
		if err != nil {
			log.Error().Err(err).Msg("Error sending discord message:")
			return
		}

		log.Debug().
			Str("Phase", "Daily").
			Msg("[Daily] Gun smoke reminder sent.")
	})

	// Send message at 23:01 of computer
	s.c.AddFunc("1 23 * * *", func() {
		isFrontierConquest, err := s.cal.IsFrontierConquest(time.Now())
		if err != nil {
			log.Error().Err(err).Msg("Error when get database value.")
			return
		}

		log.Debug().
			Str("Phase", "Daily").
			Bool("isFrontierConquest", isFrontierConquest).
			Msg("Check if event message needed.")

		if !isFrontierConquest {
			return
		}

		err = bot.SendFrontierConquestReminder()
		if err != nil {
			log.Error().Err(err).Msg("Error sending discord message:")
			return
		}

		log.Debug().
			Str("Phase", "Daily").
			Msg("[Daily] Frontier Conquest reminder sent.")
	})
}
