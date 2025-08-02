package luhn

import (
    "strings"
    "strconv"
)

func Valid(id string) bool {
    id = strings.ReplaceAll(id, " ", "")

    if len(id) < 2 {
        return false
    }

	sum := 0
    process := false
    for i := len(id) ; i > 0 ; i-- {
        numericValue, err := strconv.Atoi(id[i-1:i])
		
        if  err != nil {
            return false
        }

        if process {
            numericValue+=numericValue
            if numericValue > 9 {
                numericValue -= 9
            }
    	}
    
        sum += numericValue
        process = !process
	}

    return sum % 10 == 0
}
