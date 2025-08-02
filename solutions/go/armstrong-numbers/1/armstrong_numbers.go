package armstrong

import (
 	"fmt"
    "math"
    "strconv"
)
func IsNumber(n int) bool {
    asString := fmt.Sprintf("%v", n)
    sum := 0.0
    for _, digit := range asString {
        if num, ok := strconv.Atoi(string(digit)); ok == nil {
        	sum += math.Pow(float64(num), float64(len(asString)))
        }
    }

    return float64(n) == sum
}
