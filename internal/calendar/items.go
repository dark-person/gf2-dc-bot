package calendar

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
