package booking

import "time"
import "fmt"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layout := "1/2/2006 15:4:5"
	parsed, _ := time.Parse(layout, date)
    return parsed
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    layout := "January 2, 2006 15:04:05"
    parsed, _ := time.Parse(layout, date)
    return parsed.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    layout := "Monday, January 2, 2006 15:04:05"
    parsed, _ := time.Parse(layout, date)
    return parsed.Hour() >= 12 && parsed.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	parsed := Schedule(date)
    return parsed.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    value := fmt.Sprintf("%v-09-15 00:00:00 +0000 UTC", time.Now().Year())
    layout := "2006-01-02 15:04:05 -0700 UTC"
    parsed, _ := time.Parse(layout, value)
    return parsed
}
