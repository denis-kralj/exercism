package parsinglogfiles

import (
    "strings"
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
	exp := regexp.MustCompile(`"(?i).*(password).*"`)
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

func TagWithUserName(lines []string) []string {
	exp := regexp.MustCompile(`User\s+\w+`)
    result := []string{}
    for _, line := range lines {
        if exp.MatchString(line) {
            match := exp.FindString(line)
            split := strings.Split(match, " ")
            result = append(result, "[USR] " + split[len(split)-1] + " " + line)
        } else {
        	result = append(result, line)
        }
    }

    return result
}
