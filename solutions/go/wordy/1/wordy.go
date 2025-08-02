package wordy

import (
    "strings"
    "strconv"
    "fmt"
)

func Answer(question string) (int, bool) {
	question = strings.ReplaceAll(question, "What is ", "")
    question = strings.ReplaceAll(question, "?", "")
    question = strings.ReplaceAll(question, "plus", "+")
    question = strings.ReplaceAll(question, "minus", "-")
    question = strings.ReplaceAll(question, "divided by", "/")
    question = strings.ReplaceAll(question, "multiplied by", "*")

    if !isValidQuestion(question) {
        return 0, false
    }

    return calculate(question), true
}

func calculate(question string) int {
    split := strings.Split(question, " ")

    if len(split) == 1 {
        first, _ := strconv.Atoi(split[0])
        return first
    }

    first, _ := strconv.Atoi(split[0])
    second, _ := strconv.Atoi(split[2])
    result := 0
    switch split[1] {
        case "+": result = first + second
        case "-": result = first - second
        case "*": result = first * second
        default : result = first / second
            }
    if len(split) == 3 {
        return calculate(fmt.Sprint(result))
    } else {
        return calculate(fmt.Sprintf("%v %v %v", fmt.Sprint(result), split[3], split[4]))
    }
}
func isValidQuestion(question string) bool {
    split := strings.Split(question, " ")

    if len(split) % 2 == 0 {
        return false
    }

    for idx, val := range split {
        if idx % 2 == 0 && !isNumber(val) {
            return false
        } else if idx % 2 == 1 && !isOperation(val) {
        	return false
        }
    }

    return true
}

func isNumber(value string) bool {
    if _, err := strconv.Atoi(value); err == nil {
        return true
    }
	return false
}

func isOperation(value string) bool {
    return value == "+" || value == "-" || value == "*" || value == "/"
} 