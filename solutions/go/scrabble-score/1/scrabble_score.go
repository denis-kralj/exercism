package scrabble

import "strings"

func getScore(letter string) int {
    letterUpper := strings.ToUpper(letter)
    switch letterUpper {
    	case "D", "G": return 2
    	case "B", "C", "M", "P": return 3
    	case "F", "H", "V", "W", "Y" : return 4
        case "K" : return 5
        case "J", "X": return 8
        case "Q", "Z": return 10
    	default: return 1
    }
}

func Score(word string) (sum int) {
	for i :=0; i < len(word); i++ {
        sum += getScore(word[i:i+1])
    }
	return
}
