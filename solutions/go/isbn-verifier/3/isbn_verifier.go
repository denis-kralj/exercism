package isbn

import (
    "strings"
    "strconv"
)

func IsValidISBN(isbn string) bool {
    isbn = strings.ReplaceAll(isbn, "-", "")

    sum, coef := 0, 10

    for index, ltr := range isbn {
        if index == 9 && ltr == 'X' {
            sum+= 10
            coef--
        } else if number, ok := strconv.Atoi(string(ltr)); ok == nil {
            sum += coef * number
            coef--
        } else {
        	return false
        }
    }
	return sum % 11 == 0 && coef == 0
}
