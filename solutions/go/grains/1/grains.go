package grains

import "errors"

func actualSquare(number int) (uint64, error) {
    var count uint64 = 1
	for i := 2; i <= number; i++ {
        count += count
    }

    return count, nil
}

func Square(number int) (uint64, error) {
    if number < 1 || number > 64 {
        return 0, errors.New("must be between 1 and 64")
    }
    return actualSquare(number)
}

func Total() uint64 {
    val, _ := actualSquare(65)
	return val - 1
}
