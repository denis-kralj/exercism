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
    symbols := [4][3]string{
        { "I","V","X"},
        { "X","L","C"},
        { "C","D","M"},
        { "M","⚡","🤛"},// don't use the last 2 anyway
    }
	symbolIndex := 0

    for input > 0 {
    	remainder := input % 10
        output = getRomanNumeral(remainder, symbols[symbolIndex][0],symbols[symbolIndex][1],symbols[symbolIndex][2]) + output
        symbolIndex++
        input /= 10
    }
    return output, nil
}

func getRomanNumeral(number int, ones string, fives string, tens string) string {
    switch number {
        case 0,1,2,3: return strings.Repeat(ones, number)
    	case 4: return ones + fives
    	case 9: return ones + tens
    	default: return fives + strings.Repeat(ones, number - 5)
    }
}
