// Package leap contains leap year related logis
package leap

// IsLeapYear returns if the provided year is a leap year or not
func IsLeapYear(year int) bool {
	isDivisibleByFour := year % 4 == 0
    isDivisibleByHundred := year % 100 == 0
    isDivisibleByFourHundred := year % 400 == 0

    if isDivisibleByFour && !isDivisibleByHundred {
        return true
    }

    if isDivisibleByFour && isDivisibleByHundred && isDivisibleByFourHundred {
        return true
    }
	
    return false
}
