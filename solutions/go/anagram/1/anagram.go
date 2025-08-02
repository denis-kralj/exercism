package anagram

import "strings"

type charMap map[rune]int

func Detect(subject string, candidates []string) (result []string) {
	subjectMap := getCharMap(strings.ToLower(subject))

    for _, ele := range candidates {
        if strings.ToLower(subject) == strings.ToLower(ele) {
            continue
        }
        candidateMap := getCharMap(strings.ToLower(ele))

        if areSame(subjectMap, candidateMap) {
            result = append(result, ele)
        }
    }

    return
}

func getCharMap(input string) charMap {
    result := charMap{}
    for _, ltr := range input {
        if _, ok := result[ltr]; ok {
            result[ltr]++
        } else {
            result[ltr] = 1
        }
    }

    return result
}

func areSame(map1 charMap, map2 charMap) bool {
    if len(map1) != len(map2) {
        return false
    }

    for key, value := range map1 {
        if map2[key] != value {
            return false
        }
    }

    return true
}