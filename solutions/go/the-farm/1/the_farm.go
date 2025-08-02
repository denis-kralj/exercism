package thefarm

import (
    "errors"
    "fmt"
)

func DivideFood(fodderCalculator FodderCalculator, numberOfCows int) (float64, error) {
    fodderAmount, err := fodderCalculator.FodderAmount(numberOfCows)
    if err != nil {
        return 0, err
    }
    fatteningFactor, err := fodderCalculator.FatteningFactor()
    if err != nil {
        return 0, err
    }

    return fodderAmount / float64(numberOfCows) * fatteningFactor, nil
}

func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, numberOfCows int) (float64, error) {
 	if numberOfCows < 1 {
        return 0, errors.New("invalid number of cows")
    }
	return DivideFood(fodderCalculator, numberOfCows)
}

func ValidateNumberOfCows(numberOfCows int) error {
    var message string
    switch {
        case numberOfCows < 0: message = "there are no negative cows"
    	case numberOfCows == 0: message = "no cows don't need food"
    }

    if message != "" {
        return &InvalidCowsError{
            numberOfCows,
            message,
        }
    }

    return nil
}

type InvalidCowsError struct {
    numberOfCows int
    customMessage string
}

func (err *InvalidCowsError) Error() string {
    return fmt.Sprintf("%v cows are invalid: %v", err.numberOfCows, err.customMessage)
}