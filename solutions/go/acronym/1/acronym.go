package acronym

import (
    "unicode"
    "strings"
)
func Abbreviate(s string) string {
    s = cleanString(s)

    split := strings.Split(s, " ")
	output := ""
	for _, word := range split {
        if word == "" {
            continue
        }
        output += strings.ToUpper(string(word[0]))
    }

    return output
}

func cleanString(s string) string {
    output := ""
    for _, ltr := range s {
        if unicode.IsLetter(ltr) || ltr == ' ' || ltr == '-' {
            if ltr == '-' {
                output += " "
            } else {
            	output += string(ltr)
            }
        }
    }
	return output
}