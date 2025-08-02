package sieve

func Sieve(limit int) []int {
    numbers := make([]bool, limit + 1)
    result := []int{}

    for idx := 2; idx < limit + 1; idx++ {
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
