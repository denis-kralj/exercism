package proverb

import "fmt"

func Proverb(rhyme []string) (output []string) {
    if len(rhyme) != 0 {
    	for i := 1; i < len(rhyme); i++ {
        	output = append(output, fmt.Sprintf("For want of a %v the %v was lost.", rhyme[i-1], rhyme[i]))    
        }

	    output = append(output, fmt.Sprintf("And all for the want of a %v.", rhyme[0]))
    }
    return
}