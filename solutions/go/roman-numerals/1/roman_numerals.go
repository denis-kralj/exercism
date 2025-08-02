package romannumerals

import (
    "errors"
    "strings"
)
func ToRomanNumeral(input int) (string, error) {
    if input > 3999 || input < 1 {
        return "", errors.New("Invalid number")
    }
	output := ""
    remainder := input % 10
    output = getOnes(remainder) + output
    input /= 10
    remainder = input % 10
    output = getTens(remainder) + output
    input /= 10
    remainder = input % 10
    output = getHundreds(remainder) + output
    input /= 10
    remainder = input % 10
    output = getThousands(remainder) + output
    return output, nil
}

func getOnes(number int) string {
    switch number {
        case 0,1,2,3: return strings.Repeat("I", number)
    	case 4: return "IV"
    	case 5,6,7,8: return "V" + strings.Repeat("I", number - 5)
    	default: return "IX"
    }
}

func getTens(number int) string {
    switch number {
        case 0,1,2,3: return strings.Repeat("X", number)
    	case 4: return "XL"
    	case 5,6,7,8: return "L" + strings.Repeat("X", number - 5)
    	default: return "XC"
    }
}

func getHundreds(number int) string {
    switch number {
        case 0,1,2,3: return strings.Repeat("C", number)
    	case 4: return "CD"
    	case 5,6,7,8: return "D" + strings.Repeat("C", number - 5)
    	default: return "CM"
    }
}

func getThousands(number int) string {
    return strings.Repeat("M", number)
}