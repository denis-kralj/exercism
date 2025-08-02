package isbn

import (
    "strings"
    "strconv"
)

func IsValidISBN(isbn string) bool {
    isbn = strings.ReplaceAll(isbn, "-", "")
    if len(isbn) != 10 { // assuming no funky characters
        return false
    }

    sum := 0

    for index, ltr := range isbn {
        if index == 9 && ltr == 'X' {
            sum+= 10
            continue
        } else if number, ok := strconv.Atoi(string(ltr)); ok == nil {
            sum += (10 - index) * number
        } else {
        	return false
        }
    }
	return sum % 11 == 0
}
