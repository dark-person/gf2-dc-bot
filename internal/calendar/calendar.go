// Package to manage all scheduler event/activity.
package calendar

import "github.com/dark-person/lazydb"

// Calendar struct for manage all scheduled events.
type Calendar struct {
	db *lazydb.LazyDB // Database connection, should be init & opened before passed to this struct
}

// Create a new Calendar instance.
func NewCalendar(db *lazydb.LazyDB) *Calendar {
	return &Calendar{db: db}
}

// Data container for one scheduled event.
// This struct is designed based on database structure.
type CalendarItem struct {
	Name        string // Name of scheduled event
	StartAt     int    // Start time of scheduled event
	EndAt       int    // End time of scheduled event
	IsPredicted bool   // This scheduled event is predicted or not
}

// Create a empty calendar item.
func Empty() CalendarItem {
	return CalendarItem{
		Name:        "",
		StartAt:     -1,
		EndAt:       -1,
		IsPredicted: false,
	}
}

// Check if the calendar item is empty.
func (c *CalendarItem) IsEmpty() bool {
	return c.Name == "" && c.StartAt == -1 && c.EndAt == -1 && !c.IsPredicted
}
