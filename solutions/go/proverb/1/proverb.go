// Package proverb generates proverbs based on input
package proverb

import "fmt"

const template_line string = "For want of a %v the %v was lost."
const template_end string = "And all for the want of a %v."

// Proverb generates a list of proverbs based on inputs
func Proverb(rhyme []string) (output []string) {
    if len(rhyme) == 0 {
        return
    }
	for i := 1; i < len(rhyme); i++ {
    	output = append(output, fmt.Sprintf(template_line, rhyme[i-1], rhyme[i]))    
    }

    output = append(output, fmt.Sprintf(template_end, rhyme[0]))
    return
}
