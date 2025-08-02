package hamming

import "errors"
func Distance(a, b string) (count int, err error) {
	if len(a) != len(b) {
        err = errors.New("must be equal length")
        return
    }

    for i := range a {
        if a[i] != b[i] {
            count++
        }
    }

    return
}
