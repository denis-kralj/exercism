package collatzconjecture
import "errors"
func CollatzConjecture(n int) (int, error) {
    if n < 1 {
        return 0, errors.New("Number must be equal to or greater than 1")
    }
    steps := 0

    for n != 1 {
        if n % 2 == 0 {
            n /= 2
        } else {
        	n = 1 + n * 3
        }
    	steps++
    }

    return steps, nil
}
