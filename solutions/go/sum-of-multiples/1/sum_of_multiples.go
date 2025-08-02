package summultiples

// import "fmt"
type nothing struct{}

func SumMultiples(limit int, divisors ...int) (sum int) {
    for i := 1; i < limit; i++ {
    	for _, ele := range divisors {
        	if ele != 0 && i % ele == 0 {
                sum+= i
                break
            }
        }
	}

    return
}
