// Package bob 
package bob

import (
    "unicode"
    "strings"
)


const sure string = "Sure."
const whoa string = "Whoa, chill out!"
const relax string = "Calm down, I know what I'm doing!"
const fine string = "Fine. Be that way!"
const whatever string = "Whatever."

// Hey should have a comment documenting it.
func Hey(remark string) string {
	switch {
        case IsQuestion(remark) && !IsYelling(remark) && !IsSilence(remark): return sure
    	case !IsQuestion(remark) && IsYelling(remark) && !IsSilence(remark): return whoa
    	case IsQuestion(remark) && IsYelling(remark) && !IsSilence(remark): return relax
    	case IsSilence(remark): return fine
    	default: return whatever
    }
}

func IsQuestion(input string) bool {
    trimmed := strings.TrimSpace(input)
    return len(trimmed) > 0 && trimmed[len(trimmed)-1] == '?'
}

func IsYelling(input string) bool {
    hadLetter := false
    for _, ltr := range input {
        if unicode.IsLetter(ltr) {
            hadLetter = true
        }
        if unicode.IsLetter(ltr) && unicode.IsLower(ltr) {
            return false
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