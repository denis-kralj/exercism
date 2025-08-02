package atbash

import (
    "unicode"
	"strings"
)

var letterMap = map[rune]rune{
    'a':'z',
    'b':'y',
    'c':'x',
    'd':'w',
    'e':'v',
    'f':'u',
    'g':'t',
    'h':'s',
    'i':'r',
    'j':'q',
    'k':'p',
    'l':'o',
    'm':'n',
    'n':'m',
    'o':'l',
    'p':'k',
    'q':'j',
    'r':'i',
    's':'h',
    't':'g',
    'u':'f',
    'v':'e',
    'w':'d',
    'x':'c',
    'y':'b',
    'z':'a',    
}

func Atbash(s string) string {
    result := ""
    charCount := 0
	for _, ltr := range s {
        ltr := unicode.ToLower(ltr)
        if unicode.IsLetter(ltr) {
            result += string(letterMap[ltr])
            charCount++
        }

        if unicode.IsNumber(ltr) {
            result += string(ltr)
            charCount++
        }

        if charCount == 5 {
            result += " "
            charCount = 0
        }
    }

    return strings.TrimSpace(result)
}
