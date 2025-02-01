package scheduler

import (
	"fmt"
	"strconv"
	"time"
)

// Container to store flags for notification.
type scheduleDetails struct {
	isGunSmokeFrontline bool // Indicate is today in gun-smoke frontline
	isActivityBattle    bool // Indicate is today in activity, and can battle
}

// Get schedule details from given date.
func (s *Scheduler) getScheduleDetails(t time.Time) (*scheduleDetails, error) {
	currentDate, err := strconv.Atoi(t.Format("20060102"))
	if err != nil {
		fmt.Println(currentTimeStr(), "Error getting current date:", err)
		return nil, err
	}

	rows, err := s.db.Query("SELECT scheulde_name FROM `calendar_ranged` WHERE start_at <= ? AND end_at >= ?",
		currentDate, currentDate)
	if err != nil {
		fmt.Println(currentTimeStr(), "Error getting details:", err)
		return nil, err
	}

	// Init result
	result := &scheduleDetails{
		isGunSmokeFrontline: false,
		isActivityBattle:    false,
	}

	// Loop through all rows
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println(currentTimeStr(), "Error scanning name:", err)
			return nil, err
		}

		switch name {
		case "塵煙":
			result.isGunSmokeFrontline = true

		case "活動物資":
			result.isActivityBattle = true

		default:
			fmt.Println(currentTimeStr(), "Unknown schedule name:", name)
		}
	}

	return result, nil
}
