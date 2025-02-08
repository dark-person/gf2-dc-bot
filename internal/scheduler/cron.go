package scheduler

import (
	"fmt"
	"time"

	"github.com/dark-person/gf2-dc-bot/internal/calendar"
	"github.com/dark-person/gf2-dc-bot/pkg/api"
)

// Get daily notification as string. This function is public method due to more flexibility.
func (s *Scheduler) GetDailyReminderMsg() string {
	// Get instance of current time
	t := time.Now()

	// Set fixed daily info
	msg := ""
	msg += "哼，是時候確認一下自己的:\n"
	msg += "### 每日\n"
	msg += "- 品質甄選 > 常駐商店 > 每日禮包\n"
	msg += "- 實兵演習 3 場\n"

	// Get info details
	calItems, err := s.cal.GetCurrentActivity(t)
	if err != nil {
		fmt.Println(currentTimeStr(), "Error getting reminder details:", err)
		return ""
	}

	// Check if activity is going on
	if calendar.HasScheduleName(calItems, "活動物資") {
		fmt.Println(currentTimeStr(), "Activity Battle Detected.")
		item := calendar.GetFirstByScheduleName(calItems, "活動物資")

		start := calendar.ConvertIntToTime(item.StartAt)
		end := calendar.ConvertIntToTime(item.EndAt)

		msg += "- 活動自律 3 場 (" + start.Format("2006-01-02") + " ~ " + end.Format("2006-01-02") + ")\n"
	}

	// Check if weekday is sunday
	if t.Weekday() == time.Sunday || t.Weekday() == time.Saturday {
		fmt.Println(currentTimeStr(), "Saturday/Sunday detected.")
		msg += "\n### 每周特別提醒:\n"
		msg += "- 首領挑戰自律 3 場\n"
		msg += "- 公會商店兌換 **鍋鍋沙**\n"
		msg += "- 首領商店兌換 **紫核、好感度道具**\n"
		msg += "- 調度商店兌換 **抽抽、好感度道具，能全掃就掃**\n"
	}

	// Check if today is last two day of current month
	currentYear, currentMonth, currentDay := t.Date()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, t.Location())
	last2DayOfMonth := firstOfMonth.AddDate(0, 1, -2).Day()

	if currentDay >= last2DayOfMonth {
		fmt.Println(currentTimeStr(), "Last two day of month detected.")
		msg += "\n### 月底特別提醒:\n"
		msg += "- 首領商店兌換 **肥霰**、**抽抽**\n"
		msg += "- 易物所兌換 **抽抽**\n"
	}

	// Check if gunsmoke frontline is running
	if calendar.HasScheduleName(calItems, "塵煙") {
		fmt.Println(currentTimeStr(), "GunSmoke Frontline Detected.")
		msg += "\n### 特別注意!\n**塵煙活動開放中, 記得要出2刀**\n"
	}

	msg += "\n_WA醬 現在在測試中, 現在模擬的日期是 " + t.Format("2006-01-02") + "_\n" // TODO: REMOVE

	// Add horizontal line break
	msg += "__                                        __"
	return msg
}

// Init cron task that work daily.
func (s *Scheduler) AddDailyCron(bot api.DiscordBot) {
	// Recalculate ranged schedule date when startup
	s.calcNextRangedDate()
	fmt.Println(currentTimeStr(), "[Startup] Next gun-smoke frontline date updated.")

	// Recalculate ranged schedule date when every day start
	s.c.AddFunc("0 0 * * *", func() {
		s.calcNextRangedDate()
		fmt.Println(currentTimeStr(), "[Daily] Next gun-smoke frontline date updated.")
	})

	// Send message at 22:00 of computer
	s.c.AddFunc("0 22 * * *", func() {
		fmt.Println(currentTimeStr(), "[Daily] Start send daily reminder message.")
		msg := s.GetDailyReminderMsg()
		fmt.Println("Message: \n\n", msg)

		// Send discord message by combined all
		err := bot.SendReminder(msg)
		if err != nil {
			fmt.Println(currentTimeStr(), "Error sending discord message:", err)
			return
		}
		fmt.Println(currentTimeStr(), "[Daily] Message sent.")
	})
}
