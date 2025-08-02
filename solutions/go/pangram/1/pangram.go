package pangram



import "unicode"
func IsPangram(input string) bool {
    letterCounter := map[rune]struct{}{}
    for _, ltr := range input {
        if unicode.IsLetter(ltr) {
            letterCounter[unicode.ToLower(ltr)] = struct{}{}
        }
    }

    return len(letterCounter) == 26
}
