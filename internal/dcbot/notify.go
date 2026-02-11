package dcbot

import (
	"fmt"
	"time"

	"github.com/dark-person/gf2-dc-bot/internal/calendar"
	"github.com/rs/zerolog/log"
)

// Get daily notification as string. This function is public method due to more flexibility.
func (bm *BotManager) getDailyReminderMsg() string {
	// Get instance of current time
	t := time.Now()

	log.Debug().
		Time("demoAt", t).
		Msg("Reminder message creating.")
	prefix, suffix := bm.persona.GetReminderCustomizedStr(t)

	// Set fixed daily info
	msg := prefix
	msg += "### 每日\n"
	msg += "- 品質甄選 > 常駐商店 > 每日禮包\n"
	msg += "- 實兵演習 3 場\n"
	msg += "- 檢查 邊界推進\n"

	// Get info details
	calItems, err := bm.cal.GetCurrentActivity(t)
	if err != nil {
		log.Error().Err(err).Msg("Error getting reminder details.")
		return ""
	}

	// Check if activity is going on
	if calendar.HasScheduleName(calItems, "活動物資") {
		log.Debug().Msg("Activity Battle Detected.")
		item := calendar.GetFirstByScheduleName(calItems, "活動物資")

		start := calendar.ConvertIntToTime(item.StartAt)
		end := calendar.ConvertIntToTime(item.EndAt)

		msg += "- 活動自律 3 場 (" + start.Format("2006-01-02") + " ~ " + end.Format("2006-01-02") + ")\n"
	}

	// Check if weekday is sunday
	if t.Weekday() == time.Sunday || t.Weekday() == time.Saturday {
		log.Debug().Msg("Saturday/Sunday detected.")
		msg += "\n### 每周特別提醒:\n"
		msg += "- 首領挑戰自律 3 場\n"
		msg += "- 公會商店兌換 **鍋鍋沙**\n"
		msg += "- 首領商店兌換 **紫核、好感度道具**\n"
		msg += "- 調度商店兌換 **抽抽、好感度道具，能全掃就掃**\n"
	}

	// Cycle Event reminder
	cycled, err := bm.cal.GetUpcomingDeadline(t)
	if err != nil {
		log.Error().Err(err).Msg("Error getting upcoming deadline.")
		return ""
	}

	// Get event that require reminder
	cycled = calendar.FilterNoRemindCycledEvents(cycled, t)

	if len(cycled) > 0 {
		log.Debug().Msg("Cycle event detected.")
		msg += "\n### 循環活動提醒:\n"
		for _, evt := range cycled {
			msg += fmt.Sprintf("- %s (剩餘 %d 天)\n", evt.Name, evt.RemainDays(t))
		}
	}

	// Check if today is last two day of current month
	currentYear, currentMonth, currentDay := t.Date()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, t.Location())

	last2DayOfMonth := firstOfMonth.AddDate(0, 1, -2).Day()
	last6DayOfMonth := firstOfMonth.AddDate(0, 1, -6).Day()

	if currentDay >= last6DayOfMonth {
		log.Debug().Msg("Monthly reminder is needed.")

		log.Debug().Msg("Last two day of month detected.")
		msg += "\n### 月底特別提醒:\n"

		if currentDay >= last2DayOfMonth {
			msg += "- 首領商店兌換 **肥霰**、**抽抽**\n"
			msg += "- 易物所兌換 **抽抽**\n"
		}
	}

	// Check if gun smoke frontline is running
	if calendar.HasScheduleName(calItems, "塵煙") {
		log.Debug().Msg("GunSmoke Frontline Detected.")

		item := calendar.GetFirstByScheduleName(calItems, "塵煙")

		start := calendar.ConvertIntToTime(item.StartAt)
		end := calendar.ConvertIntToTime(item.EndAt)

		msg += "\n### 特別注意!\n**塵煙活動開放中, 記得要出2刀** (" +
			start.Format("2006-01-02") + " ~ " + end.Format("2006-01-02") + ")\n"
	}

	msg += suffix

	// Add horizontal line break
	msg += "__                                        __"
	return msg
}

// Send message to discord channel. Please note that discord token must be set before call this function.
func (bm *BotManager) SendReminder() error {
	if !bm.initialized {
		return fmt.Errorf("bot manager not initialized")
	}

	// Prepare message
	message := bm.getDailyReminderMsg()

	_, err := bm.session.ChannelMessageSend(bm.ReminderChannel, message)
	if err != nil {
		return fmt.Errorf("failed to send message to discord: %v", err)
	}
	return nil
}

// Send message to discord channel, to notify specific role to remember gun-smoke frontline event.
// Please note that discord token must be set before call this function.
func (bm *BotManager) SendGunSmokeReminder() error {
	if !bm.initialized {
		return fmt.Errorf("bot manager not initialized")
	}

	// Prepare message
	message := bm.persona.GunSmokeRemindDialog(bm.cfg.GunSmokeRemindRole)

	_, err := bm.session.ChannelMessageSend(bm.ReminderChannel, message)
	if err != nil {
		return fmt.Errorf("failed to send message to discord: %v", err)
	}
	return nil
}
