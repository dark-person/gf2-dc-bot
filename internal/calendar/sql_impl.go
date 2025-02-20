package calendar

import "time"

// Get all ongoing activity (i.e. current & future activity) from database.
func (c *Calendar) GetOngoingActivity(t time.Time) ([]CalendarItem, error) {
	return c.getCalendarItems(
		"SELECT schedule_name, start_at, end_at, is_predicted FROM `calendar_ranged` WHERE start_at < ? ORDER BY start_at",
		ConvertTimeToInt(t))
}

// Get all current activity from database.
func (c *Calendar) GetCurrentActivity(t time.Time) ([]CalendarItem, error) {
	return c.getCalendarItems(
		"SELECT schedule_name, start_at, end_at, is_predicted FROM `calendar_ranged` WHERE start_at <= ? AND end_at >= ? ORDER BY start_at",
		ConvertTimeToInt(t), ConvertTimeToInt(t))
}

// Get all upcoming deadline (which is nearly expired) from database.
// This method will only return the latest event for each event type.
func (c *Calendar) GetUpcomingDeadline(t time.Time) ([]CycledEventItem, error) {
	return c.getCycledEventItem(`SELECT cycle_event_name, deadline_at, is_auto_calc, cycle_day, remind_before FROM calendar_cycled_event evt 
LEFT JOIN const_cycle_event const ON evt.cycle_event_id = const.cycle_event_id 
WHERE deadline_at >= ? ORDER BY deadline_at`,
		ConvertTimeToInt(t))
}
