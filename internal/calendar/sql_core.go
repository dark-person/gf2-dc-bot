package calendar

// Core function to get calendar items from database.
// This function assume all argument passed to this function is valid, no checking will be performed.
func (c *Calendar) getCalendarItems(query string, intTime ...any) ([]CalendarItem, error) {
	// Query calendar item by given SQL, trust the intTime parameter
	rows, err := c.db.Query(query, intTime...)

	if err != nil {
		return nil, err
	}

	// Loop results
	var items []CalendarItem

	for rows.Next() {
		var scheduleName string
		var startDate, endDate int
		var isPredicted bool

		err := rows.Scan(&scheduleName, &startDate, &endDate, &isPredicted)
		if err != nil {
			return nil, err
		}

		// To item
		item := CalendarItem{
			Name:        scheduleName,
			StartAt:     startDate,
			EndAt:       endDate,
			IsPredicted: isPredicted,
		}

		items = append(items, item)
	}

	return items, nil
}
