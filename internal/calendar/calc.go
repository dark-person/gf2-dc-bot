package calendar

import (
	"fmt"
	"time"
)

// Get current time in opinionated formatted string.
func currentTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Calculate the next ranged date by non-predicted data in database.
func (c *Calendar) CalcNextRangedDate() error {
	// Get Latest schedule item that is not predicted
	raw, err := c.GetLatestConfirmedRangedActivity()
	if err != nil {
		return fmt.Errorf("error when get confirmed ranged activity: %v", err)
	}

	// Convert raw data to map for easier manipulation and comparison
	m := LatestScheduleMap(raw)

	// Gun-smoke frontline calculation (Add three weeks)
	item := m["塵煙"]
	next := item.Copy()

	next.StartAt = ConvertTimeToInt(ConvertIntToTime(next.StartAt).AddDate(0, 0, 21))
	next.EndAt = ConvertTimeToInt(ConvertIntToTime(next.StartAt).AddDate(0, 0, 6))
	next.IsPredicted = true

	// Insert latest record to database
	err = c.addCalendarItems(next)
	if err != nil {
		return fmt.Errorf("error when insert next gun-smoke frontline date: %v", err)
	}

	return nil
}
