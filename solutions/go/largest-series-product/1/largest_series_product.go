package lsproduct

import (
    "errors"
    "strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 1 {
        return 0, errors.New("span must be 1 or more")
    }

    if span > len(digits) {
        return 0, errors.New("span must be smaller or equal to the lenght of digits")
    }

    var max int64 = 0

    for startIndex := range digits {
        if startIndex+span > len(digits) {
            break
        }
    
        var product int64 = 1
        for i := 0; i < span; i++ {
			if v, ok := strconv.Atoi(string(rune(digits[startIndex+i]))); ok != nil {
                return 0, errors.New("input must only be digits")
            } else {
            	product *= int64(v)
            }
        }
    	if product > max {
            max = product
        }
    }

    return max, nil
}
