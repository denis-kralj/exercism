package reverse

func Reverse(input string) string {
	result := ""
    asRunes := []rune(input)

    for i := len(asRunes) - 1; i >= 0; i-- {
        result += string(asRunes[i])
    }

    return result
}
