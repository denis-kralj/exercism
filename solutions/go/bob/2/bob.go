package bob

import (
    "unicode"
    "strings"
)

func Hey(remark string) string {
	switch {
    	case IsSilence(remark): return "Fine. Be that way!"
    	case IsQuestion(remark) && IsYelling(remark): return "Calm down, I know what I'm doing!"
        case IsQuestion(remark): return "Sure."
    	case IsYelling(remark): return "Whoa, chill out!"
    	default: return "Whatever."
    }
}

func IsQuestion(input string) bool {
    trimmed := strings.TrimSpace(input)
    return trimmed[len(trimmed)-1] == '?'
}

func IsYelling(input string) bool {
    hadLetter := false
    for _, ltr := range input {
        if unicode.IsLetter(ltr) {
            hadLetter = true
            if unicode.IsLower(ltr) {
            	return false
        	}
        }
    }
    return true && hadLetter
}

func IsSilence(input string) bool {
    for _, ltr := range input {
        if !unicode.IsSpace(ltr) {
            return false
        }
    }

    return true
}