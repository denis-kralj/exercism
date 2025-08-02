// Package weather contains tools for relaying weather
// conditions and forcasting.
package weather

// CurrentCondition represents the current weather condition for the CurrentLocation.
var CurrentCondition string
// CurrentLocation represents the current location that we want to report conditions on.
var CurrentLocation string

// Forecast returns a string  describing the current weather conditions for the provided
// city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
