package wordcount

import (
    "unicode"
    "strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
    result := make(Frequency)
	split := strings.FieldsFunc(phrase, Splitters)

    for _, word := range split {
        word = strings.ToLower(word)
		trimmed := strings.Trim(word, "'")
        if trimmed == "" {
            continue
        }
        if _, ok := result[trimmed]; ok {
        	result[trimmed]++    
        } else {
        	result[trimmed] = 1
        }
        
    }

    return result
}

func Splitters(r rune) bool {
    return !unicode.IsLetter(r)  && !unicode.IsNumber(r) && r != '\''
}
