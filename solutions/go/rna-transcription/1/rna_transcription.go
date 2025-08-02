package strand

var conversionMap map[rune]rune = map[rune]rune{
    'G':'C',
    'C':'G',
    'T':'A',
    'A':'U',
}

func ToRNA(dna string) (result string) {
    for _, ltr := range dna {
        result += string(conversionMap[ltr])
    }
	return
}
