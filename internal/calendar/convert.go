package calendar

import "time"

// Convert time to database format.
func ConvertTimeToInt(t time.Time) int {
	year := t.Year()
	month := t.Month()
	day := t.Day()

	return year*10000 + int(month)*100 + day
}

// Convert database format to time, with system timezone.
func ConvertIntToTime(date int) time.Time {
	year := date / 10000
	month := (date % 10000) / 100
	day := date % 100

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}
