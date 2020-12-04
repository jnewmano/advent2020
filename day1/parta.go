package main

import (
	"fmt"

	"github.com/jnewmano/advent2020/input"
)

var check = false

var things = input.LoadSliceInt("")

func main() {
	sum := parta()
	fmt.Println(sum)
}

// find the two entries that sum to 2020 and then multiply those two numbers together.
func parta() int {

	for _, a := range things {
		for _, b := range things {
			if a+b == 2020 {
				return a * b
			}
		}
	}
	return 0
}
