package calendar

import "time"

// Get all ongoing activity (i.e. current & future activity) from database.
func (c *Calendar) GetOngoingActivity(t time.Time) ([]CalendarItem, error) {
	return c.getCalendarItems(
		"SELECT scheulde_name, start_at, end_at, is_predicted FROM `calendar_ranged` WHERE start_at < ? ORDER BY start_at",
		ConvertTimeToInt(t))
}

// Get all current activity from database.
func (c *Calendar) GetCurrentActivity(t time.Time) ([]CalendarItem, error) {
	return c.getCalendarItems(
		"SELECT scheulde_name, start_at, end_at, is_predicted FROM `calendar_ranged` WHERE start_at <= ? AND end_at >= ? ORDER BY start_at",
		ConvertTimeToInt(t), ConvertTimeToInt(t))
}
