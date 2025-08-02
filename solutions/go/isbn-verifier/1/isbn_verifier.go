package isbn

import (
    "strings"
    "unicode/utf8"
    "strconv"
)

func IsValidISBN(isbn string) bool {
    isbn = strings.ReplaceAll(isbn, "-", "")
    if utf8.RuneCountInString(isbn) != 10 {
        return false
    }

    sum := 0
    coef := 10

    for index, ltr := range isbn {
        if index == 9 && ltr == 'X' {
            sum+= 10
            continue
        }

        if number, ok := strconv.Atoi(string(ltr)); ok == nil {
            sum += coef * number
            coef--
        } else {
        	return false
        }
    }
	return sum % 11 == 0
}
