// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package gigasecond provideds a method allowing the user to add a gigasecond to an arbitrary time.Time value and returns the resulting time.Time
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond adds a gigasecond (1_000_000_000 seconds) to an arbitrary time.Time value and returns the resulting time.Time
func AddGigasecond(t time.Time) time.Time {
    seconds := 1_000_000_000.0
	oneGigaSecond := time.Duration(seconds * float64(time.Second))
    
	return t.Add(oneGigaSecond)
}
