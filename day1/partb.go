package main

import (
	"fmt"

	"github.com/jnewmano/advent2020/input"
)

var things = input.LoadSliceInt("")

func main() {
	sum := partb()
	fmt.Println(sum)
}

// find the two entries that sum to 2020 and then multiply those two numbers together.
func partb() int {

	for _, a := range things {

		for _, b := range things {

			for _, c := range things {

				if a+b+c == 2020 {
					return a * b * c
				}
			}
		}
	}
	return 0
}
