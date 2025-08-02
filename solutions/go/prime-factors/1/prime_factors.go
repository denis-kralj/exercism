package prime

func Factors(n int64) []int64 {
	primes := Sieve(n)
    factors := []int64{}

    pIdx := 0
    for n != 1 && pIdx < len(primes) {
        for n % primes[pIdx] != 0 {
            pIdx++
        }

        n /= primes[pIdx]
        factors = append(factors, primes[pIdx])
    }

    return factors
}

func Sieve(limit int64) []int64 {
    if limit > 1000000 {
        limit = 1000000
    }
    result := []int64{}
    if limit < 2 {
        return result
    }

    numbers := make([]bool, limit + 1)

    for idx := int64(2); idx < limit + int64(1); idx++ {
        if numbers[idx] {
            continue
        }

        result = append(result, idx)

        for idxNonPrime := idx * 2; idxNonPrime < limit + 1; idxNonPrime+=idx {
            numbers[idxNonPrime] = true
        }
    }

    return result
}