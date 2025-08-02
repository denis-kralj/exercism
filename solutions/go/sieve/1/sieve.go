package sieve

func Sieve(limit int) []int {
    numbers := make([]bool, limit + 1)
    numbers[0], numbers[1] = true, true

    for idx := 2; idx < limit + 1; idx++ {
        if numbers[idx] {
            continue
        }

        for idxNonPrime := idx * 2; idxNonPrime < limit + 1; idxNonPrime+=idx {
            numbers[idxNonPrime] = true
        }
    }

    result := []int{}

    for prime, isNotPrime := range numbers {
        if !isNotPrime {
            result = append(result, prime)
        }
    }

    return result
}
