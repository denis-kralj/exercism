package diffsquares

func SquareOfSum(n int) (sum int) {
	for i := 1; i <= n; i++ {
        sum += i
    }

    sum *= sum

    return
}

func SumOfSquares(n int) (sum int) {
	for i := 1; i <= n; i++ {
        sum += i*i
    }

    return
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
