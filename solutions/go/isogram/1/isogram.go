package isogram

import "unicode"

func IsIsogram(word string) bool {
	letters := map[rune]int{}

    for _, rune := range word {
        if rune != '-' && rune != ' ' {
            lowerRune := unicode.ToLower(rune)
            letters[lowerRune]++
            if letters[lowerRune] > 1 {
                return false
            }
        }
    }

	return true
}
