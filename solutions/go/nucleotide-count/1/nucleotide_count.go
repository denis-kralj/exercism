package dna

import (
    "fmt"
    "errors"
)
// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

func newHistogram() Histogram {
    return map[rune]int{
        'A':0,
        'C':0,
        'G':0,
        'T':0,
    }
}
// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (h Histogram, err error) {
    result := newHistogram()
    for _, ltr := range d {
        if _, ok := result[ltr]; !ok {
            err = errors.New(fmt.Sprintf("Letter '%v' isn't a valid DNA element", ltr))
            return
        }
    	result[ltr]++
    }
	h = result
	return
}
