package phonenumber

import (
    "errors"
    "fmt"
    "unicode"
)

func Number(phoneNumber string) (string, error) {
    output := ""
    for _, ele := range phoneNumber {
        if unicode.IsDigit(ele) {
            output += string(ele)
        }
    }

    outputLength := len(output)

    if outputLength < 10 || outputLength > 11 {
        return "", errors.New("Invalid number")
    }

    if outputLength == 10 {
        output = "1" + output
    }

    if output[0] != '1' || output[1] == '0' ||output[1] == '1' || output[4] == '0' ||output[4] == '1' {
        return "", errors.New("Invalid country, exchange or number plan area code")
    }
    return output[1:], nil
}

func AreaCode(phoneNumber string) (string, error) {
	if value, err := Number(phoneNumber); err == nil {
        return fmt.Sprintf("%v", value[:3]), nil
    } else {
    	return "", err
    }
}

func Format(phoneNumber string) (string, error) {
		if value, err := Number(phoneNumber); err == nil {
        return fmt.Sprintf("(%v) %v-%v", value[:3], value[3:6], value[6:10]), nil
    } else {
    	return "", err
    }
}
