// Package to manage all scheduler event/activity.
package calendar

import (
	"fmt"
	"time"

	"github.com/dark-person/lazydb"
)

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

// Create a deep copy of calendar item.
func (c CalendarItem) Copy() CalendarItem {
	return CalendarItem{
		Name:        c.Name,
		StartAt:     c.StartAt,
		EndAt:       c.EndAt,
		IsPredicted: c.IsPredicted,
	}
}

// Data container for cycled event. This struct is designed based on database structure.
type CycledEventItem struct {
	Name         string // Name of cycled event
	Deadline     int    // Deadline date of cycled event
	IsAutoCalc   bool   // This event date is auto calculated or not
	CycleDay     int    // Day amount for entire cycle
	RemindBefore int    // When to remind user before deadline
}

// Check if given time require a reminder on current cycled event.
func (c *CycledEventItem) IsNeedReminder(t time.Time) bool {
	// Calculate the reminder date by subtraction
	remindDate := ConvertIntToTime(c.Deadline).AddDate(0, 0, c.RemindBefore*-1)

	// Check if given time is larger than current deadline
	return t.After(remindDate)
}

// Get remaining day of this cycled event.
// If this event is expired, then a warning log will appear.
func (c *CycledEventItem) RemainDays(t time.Time) int {
	temp := ConvertTimeToInt(t)
	days := c.Deadline - temp

	if days < 0 {
		fmt.Println("[WARN] cycled event expired: ", c)
	}

	return days
}
