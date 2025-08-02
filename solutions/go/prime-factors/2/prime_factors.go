package prime

func Factors(n int64) []int64 {
    factors := []int64{}
	divisor := int64(2)
    for n != 1 {
        if n % divisor == 0 {
            n /= divisor
            factors = append(factors, divisor)
        } else {
        	divisor++
        }
    }

    return factors
}
