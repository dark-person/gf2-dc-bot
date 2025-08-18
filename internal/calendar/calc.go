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
	raw, err := c.getLatestRangedActivity()
	if err != nil {
		return fmt.Errorf("error when get confirmed ranged activity: %v", err)
	}

	// Convert raw data to map for easier manipulation and comparison
	m := LatestScheduleMap(raw)

	// ---- Gun-smoke frontline calculation (Add three weeks)----
	item := m["塵煙"]

	// Skip calculation if it is not expired
	endAt := ConvertIntToTime(item.EndAt)
	if !endAt.Before(time.Now()) {
		return nil
	}

	next := item.Copy()
	next.StartAt = ConvertTimeToInt(ConvertIntToTime(next.StartAt).AddDate(0, 0, 21))
	next.EndAt = ConvertTimeToInt(ConvertIntToTime(next.StartAt).AddDate(0, 0, 6))
	next.IsPredicted = false

	// Insert latest record to database
	err = c.addCalendarItems(next)
	if err != nil {
		return fmt.Errorf("error when insert next gun-smoke frontline date: %v", err)
	}

	return nil
}

// Calculate the next deadline by latest
// (i.e. can be calculated by program)
// data in database.
func (c *Calendar) CalcNextDeadline() error {
	// Get Latest cycled event
	raw, err := c.getLatestDeadline()
	if err != nil {
		return fmt.Errorf("error when get latest deadline: %v", err)
	}

	t := time.Now()

	// Loop every items
	for _, item := range raw {
		// Skip if next deadline is ready
		itemDeadline := ConvertIntToTime(item.Deadline)
		if t.Before(itemDeadline) {
			continue
		}

		// Calculate
		next := item.Copy()

		nextDeadline := ConvertIntToTime(next.Deadline).AddDate(0, 0, next.CycleDay)
		next.Deadline = ConvertTimeToInt(nextDeadline)
		next.IsAutoCalc = true

		// Insert to database
		err := c.addCycledEventItem(next)
		if err != nil {
			return fmt.Errorf("error when insert next deadline: %v", err)
		}
	}

	return nil
}
