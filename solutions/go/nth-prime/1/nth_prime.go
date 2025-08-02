package prime

import "errors"
// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
        return 0, errors.New("n must be 1 or higher")
    }
	primes := []int{2, 3, 5, 7, 11}
    candidate := 12

    for len(primes) < n {
        isPrime := true
        for _, prime := range primes {
            if candidate % prime == 0 {
                isPrime = false
                break
            }
        }

        if isPrime {
            primes = append (primes, candidate)
        }

        candidate++
    }
    return primes[n-1], nil
}
