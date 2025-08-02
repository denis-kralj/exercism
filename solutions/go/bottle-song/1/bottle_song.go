package bottlesong

import (
    "strings"
    "fmt"
)

const lineOneTwoTemplate string = "%v green bottle%v hanging on the wall,"
const lineThreeTemplate string = "And if %v green bottle%v should accidentally fall,"
const lineFourTemplate string = "There'll be %v green bottle%v hanging on the wall."

var bottleCountStrings [11]string = [11]string{"no", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

func Recite(startBottles, takeDown int) []string {
    result := []string{}

    for i := startBottles; i > startBottles-takeDown; i-- {
        result = append(result, fmt.Sprintf(lineOneTwoTemplate, capFirst(bottleCountStrings[i]), getSuffix(i)))
        result = append(result, fmt.Sprintf(lineOneTwoTemplate, capFirst(bottleCountStrings[i]), getSuffix(i)))
        result = append(result, fmt.Sprintf(lineThreeTemplate, bottleCountStrings[1], getSuffix(1)))
        result = append(result, fmt.Sprintf(lineFourTemplate, bottleCountStrings[i-1], getSuffix(i-1)))
        if i-1 > startBottles-takeDown {
            result = append(result, "")
		}
    }

	return result
}

func capFirst(s string) string {
    return strings.ToUpper(s[:1]) + s[1:]    
}

func getSuffix(count int) string {
    if count == 1 {
        return ""
    }
	return "s"
}