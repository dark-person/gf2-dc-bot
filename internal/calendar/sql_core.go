package calendar

import "time"

// Core function to check if given time is within range of gun-smoke frontline.
func (c *Calendar) IsGunSmokeFrontline(t time.Time) (bool, error) {
	row, err := c.db.QueryRow("SELECT COUNT(*) FROM `calendar_ranged` WHERE start_at <= ? AND end_at >= ? AND schedule_name=? ORDER BY start_at", ConvertTimeToInt(t), ConvertTimeToInt(t), "塵煙")

	if err != nil {
		return false, err
	}

	var count int
	err = row.Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

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

// Core function to add calendar items to database.
// This function assume all argument passed to this function is valid, no checking will be performed.
func (c *Calendar) addCalendarItems(item CalendarItem) error {
	// Insert record back to database
	_, err := c.db.Exec(
		"INSERT OR IGNORE INTO `calendar_ranged` (schedule_name, start_at, end_at, is_predicted) VALUES (?, ?, ?, ?)",
		item.Name, item.StartAt, item.EndAt, item.IsPredicted,
	)

	return err
}

// Add given cycled event item to database.
func (c *Calendar) addCycledEventItem(item CycledEventItem) error {
	// Insert record back to database
	_, err := c.db.Exec(
		"INSERT OR IGNORE INTO `calendar_cycled_event` (cycle_event_id, deadline_at, is_auto_calc) VALUES (?, ?, ?)",
		item.EventID, item.Deadline, item.IsAutoCalc,
	)

	return err
}

// Core function to get cycled event from database.
// This function assume all argument passed to this function is valid, no checking will be performed.
func (c *Calendar) getCycledEventItem(query string, intTime ...any) ([]CycledEventItem, error) {
	// Query calendar item by given SQL, trust the intTime parameter
	rows, err := c.db.Query(query, intTime...)

	if err != nil {
		return nil, err
	}

	// Loop results
	var items []CycledEventItem

	for rows.Next() {
		var id uint
		var scheduleName string
		var deadlineDate int
		var isPredicted bool
		var cycleDay int
		var remindBefore int

		err := rows.Scan(&id, &scheduleName, &deadlineDate, &isPredicted, &cycleDay, &remindBefore)
		if err != nil {
			return nil, err
		}

		// To item
		item := CycledEventItem{
			EventID:      id,
			Name:         scheduleName,
			Deadline:     deadlineDate,
			IsAutoCalc:   isPredicted,
			CycleDay:     cycleDay,
			RemindBefore: remindBefore,
		}

		items = append(items, item)
	}

	return items, nil
}
