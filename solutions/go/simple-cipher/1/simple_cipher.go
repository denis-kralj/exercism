package cipher

import "strings"
import "unicode"
var letters = []rune {'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z'}

type shift struct{
    shiftCount int
}
type vigenere struct{
    mask string
}

func NewCaesar() Cipher {
    return shift{3}
}

func NewShift(distance int) Cipher {
    if distance == 0 || distance > 25 || distance < -25 {
        return nil
    }
    if distance < 0 {
        distance = 26 + distance
    }

	return shift{distance % 26}
}

func (c shift) Encode(input string) string {
    result := ""
    input = normalizeInput(input)
    for _, ltr := range input {
        originalIndex := getIndex(ltr)
        newIndex := (originalIndex+c.shiftCount) % 26
        result += string(letters[newIndex])
    }
	return result
}

func getIndex(ltr rune) int {
    for i, v := range letters {
        if v == ltr {
            return i
        }
    }

    panic("this shouldn't happen, letter was [" + string(ltr) + "]")
}

func (c shift) Decode(input string) string {
	result := ""
    input = normalizeInput(input)
    for _, ltr := range input {
        originalIndex := getIndex(ltr)
        newIndex := (originalIndex + 26 - c.shiftCount) % 26
        result += string(letters[newIndex])
    }
	return result
}

func NewVigenere(key string) Cipher {
    if key != normalizeInput(key) || strings.ReplaceAll(key, "a", "") == "" {
        return nil
    }
	return vigenere{
        mask: key,
    }
}

func (v vigenere) Encode(input string) string {
    result := ""
    input = normalizeInput(input)

    for idx, ltr := range input {
        maskIndex := idx % len(v.mask)
        maskLtr := ([]rune(v.mask))[maskIndex]
        maskLtrIndex := getIndex(maskLtr)
        exchangeIndex := (getIndex(ltr) + maskLtrIndex) % 26
        result += string(letters[exchangeIndex])
    }
	return result
}

func (v vigenere) Decode(input string) string {
	result := ""
    input = normalizeInput(input)

    for idx, ltr := range input {
        maskIndex := idx % len(v.mask)
        maskLtr := ([]rune(v.mask))[maskIndex]
        maskLtrIndex := getIndex(maskLtr)
        exchangeIndex := (getIndex(ltr) + 26 - maskLtrIndex) % 26
        result += string(letters[exchangeIndex])
    }
	return result
}

func normalizeInput(input string) string {
    output := ""

    for _, ltr := range input {
        if unicode.IsLetter(ltr) {
            output += string(unicode.ToLower(ltr))
        }
    }
	return output
}