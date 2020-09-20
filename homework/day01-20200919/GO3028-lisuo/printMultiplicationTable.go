package main

/*
a short code to print 9 X 9 mutiplication chart
*/

import (
	"fmt"
)

// print 9 X 9 mutiplication chart
func main() {
	// print vertical
	for i := 1; i <= 9; i++ {
		// print horizontal line
		for j := 1; j <= i; j++ {
			fmt.Printf("%v X %v = %-4v", j, i, j*i)
		}
		// return at the end of each line
		fmt.Println("")
	}
}
