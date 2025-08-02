package reverse

func Reverse(input string) string {
    asRunes := []rune(input)

    for i := 0; i < len(asRunes) / 2; i++ {
        asRunes[i], asRunes[len(asRunes) - 1 - i] = asRunes[len(asRunes) - 1 - i], asRunes[i]
    }

    return string(asRunes)
}
