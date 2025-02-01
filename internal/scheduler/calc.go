package scheduler

import (
	"fmt"
	"time"
)

// Calculate the next ranged date by non-predicted data in database.
func (s *Scheduler) calcNextRangedDate() {
	// Gun-smoke frontline day, add two weeks
	row, err := s.db.QueryRow("SELECT start_at, end_at FROM `calendar_ranged` WHERE scheulde_name = '塵煙' AND is_predicted = 0")
	if err != nil {
		fmt.Println(currentTimeStr(), "Error getting next gun-smoke frontline date:", err)
		return
	}

	var startDate, endDate string
	err = row.Scan(&startDate, &endDate)
	if err != nil {
		fmt.Println(currentTimeStr(), "Error scanning next gun-smoke frontline date:", err)
		return
	}

	// Parse day
	lastStartDate, err := time.Parse("20060102", startDate)
	if err != nil {
		fmt.Println(currentTimeStr(), "Error parsing next gun-smoke frontline date:", err)
		return
	}

	// ========================================================

	// Add three weeks
	nextStartDate := lastStartDate.AddDate(0, 0, 21)
	nextEndDate := nextStartDate.AddDate(0, 0, 6)

	// Loop three times, which result 3 entry in database
	for i := 0; i < 3; i++ {
		// Insert record back to database
		_, err = s.db.Exec("INSERT OR IGNORE INTO `calendar_ranged` (scheulde_name, start_at, end_at, is_predicted) VALUES ('塵煙', ?, ?, 1)",
			nextStartDate.Format("20060102"), nextEndDate.Format("20060102"))
		if err != nil {
			fmt.Println(currentTimeStr(), "Error inserting next gun-smoke frontline date:", err)
			return
		}

		// Add more three weeks
		nextStartDate = nextStartDate.AddDate(0, 0, 21)
		nextEndDate = nextStartDate.AddDate(0, 0, 6)
	}
}
