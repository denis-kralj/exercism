package rotationalcipher

import (
    "unicode"
)
var letters [26]rune = [26]rune{'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z'}
func RotationalCipher(plain string, shiftKey int) string {
	result := ""
    for _, ltr := range plain {
        if unicode.IsLetter(ltr) {
            lower := unicode.ToLower(ltr);
            curIndex := getCurrentIndex(lower)
            newIndex := (curIndex + shiftKey) % len(letters)
            substitute := letters[newIndex]
            if unicode.IsLower(ltr) {
                result += string(substitute)
            } else {
            	result += string(unicode.ToUpper(substitute))
            }
        } else {
        	result += string(ltr)
        }
    }

    return result
}

func getCurrentIndex(ltr rune) int {
    for index, ele := range letters {
        if ele == ltr {
            return index
        }
    }

    panic("element not found")
}