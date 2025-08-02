package parsinglogfiles

import (
    "fmt"
    "regexp"
)
func IsValidLine(text string) bool {
    exp := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return exp.MatchString(text)
}

func SplitLogLine(text string) []string {
	exp := regexp.MustCompile(`<(~|\*|=|-)*>`)

    return exp.Split(text, -1)
}

func CountQuotedPasswords(lines []string) (count int) {
	exp := regexp.MustCompile(`"(?i).*password.*"`)
    for _, line := range lines {
        if exp.MatchString(line) {
            count++
        }
    }

    return
}

func RemoveEndOfLineText(text string) string {
	exp := regexp.MustCompile(`end-of-line\d+`)
    return exp.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) (result []string) {
	exp := regexp.MustCompile(`User\s+(\w+)`)
    for _, line := range lines {
        if exp.MatchString(line) {
            line = fmt.Sprintf("[USR] %v %v", exp.FindStringSubmatch(line)[1], line)
        } 
        
        result = append(result, line)
    }

    return 
}
