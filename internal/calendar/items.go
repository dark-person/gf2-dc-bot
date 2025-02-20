package calendar

import "time"

// Check if a list of calendars item contains a event that has given schedule name.
func HasScheduleName(c []CalendarItem, scheduleName string) bool {
	for _, item := range c {
		if item.Name == scheduleName {
			return true
		}
	}
	return false
}

// Get 1st event to be appear in calendar item list.
func GetFirstByScheduleName(c []CalendarItem, scheduleName string) CalendarItem {
	for _, item := range c {
		if item.Name == scheduleName {
			return item
		}
	}
	return Empty()
}

// Get map of calendar items, which only contains latest event only.
// Key is the scheduled event name, value is the actual calendar item itself.
func LatestScheduleMap(c []CalendarItem) map[string]CalendarItem {
	m := make(map[string]CalendarItem)

	for _, item := range c {
		// Check if item name already exists
		record, existed := m[item.Name]

		// Add entry if not exist
		if !existed {
			m[item.Name] = item
			continue
		}

		// Check if current item is later than existed record
		if item.StartAt > record.StartAt {
			m[item.Name] = item
			continue
		}
	}

	return m
}

// Filter the list of calendars item that has given schedule name.
func FilterByScheduleName(c []CalendarItem, scheduleName string) []CalendarItem {
	result := []CalendarItem{}

	for _, item := range c {
		if item.Name == scheduleName {
			result = append(result, item)
		}
	}

	return result
}

// Filter the list of calendars item that has given predicted flag value.
func FilterByPredictFlag(c []CalendarItem, isPredicted bool) []CalendarItem {
	result := []CalendarItem{}

	for _, item := range c {
		if item.IsPredicted == isPredicted {
			result = append(result, item)
		}
	}

	return result
}

// Filter the cycled event that require reminder in given time.
func FilterNoRemindCycledEvents(c []CycledEventItem, t time.Time) []CycledEventItem {
	result := []CycledEventItem{}

	for _, item := range c {
		if item.IsNeedReminder(t) {
			result = append(result, item)
		}
	}

	return result
}
