package cipher

import "strings"
import "unicode"
var letters = []rune {'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z'}

type vigenere struct{
    mask string
}

type indexGetter func(int, int) int

func NewCaesar() Cipher {
    return NewVigenere("d")
}

func NewShift(distance int) Cipher {    
    if distance == 0 || distance > 25 || distance < -25 {
        return nil
    }
    if distance < 0 {
        distance = 26 + distance
    }

    return NewVigenere(string(letters[distance % 26]))
}

func getIndex(ltr rune) int {
    for i, v := range letters {
        if v == ltr {
            return i
        }
    }

    panic("this shouldn't happen, letter was [" + string(ltr) + "]")
}

func NewVigenere(key string) Cipher {
    if key != normalizeInput(key) || strings.ReplaceAll(key, "a", "") == "" {
        return nil
    }
	return vigenere{
        mask: key,
    }
}

func shift(input string, mask string, getter indexGetter) string {
    result := ""
    input = normalizeInput(input)

    for idx, ltr := range input {
        maskIndex := idx % len(mask)
        maskLtr := ([]rune(mask))[maskIndex]
        maskLtrIndex := getIndex(maskLtr)
        exchangeIndex := getter(getIndex(ltr), maskLtrIndex)
        result += string(letters[exchangeIndex])
    }
	return result
}

func (v vigenere) Encode(input string) string {
    getter := func(orgLtrIx int, maskLtrIx int) int {
        return (orgLtrIx + maskLtrIx) % 26
    }

    return shift(input, v.mask, getter)
}

func (v vigenere) Decode(input string) string {
    getter := func(orgLtrIx int, maskLtrIx int) int {
        return (orgLtrIx + 26 - maskLtrIx) % 26
    }
	return shift(input, v.mask, getter)
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