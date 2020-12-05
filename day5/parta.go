package main

import (
	"fmt"

	"github.com/jnewmano/advent2020/input"
	"github.com/jnewmano/advent2020/output"
)

func main() {
	//	fmt.Println(processPass("BFFFBBFRRR"))
	//os.Exit(1)

	sum := parta()
	fmt.Println(sum)
}

// where F means "front", B means "back", L means "left", and R means "right"
/*
Start by considering the whole range, rows 0 through 127.
F means to take the lower half, keeping rows 0 through 63.
B means to take the upper half, keeping rows 32 through 63.
F means to take the lower half, keeping rows 32 through 47.
B means to take the upper half, keeping rows 40 through 47.
B keeps rows 44 through 47.
F keeps rows 44 through 45.
The final F keeps the lower of the two, row 44.
*/
/*
The last three characters will be either L or R; these specify exactly one of the 8 columns of seats on the plane (numbered 0 through 7). The same process as above proceeds again, this time with only three steps. L means to keep the lower half, while R means to keep the upper half.

*/

func parta() interface{} {
	// input.SetRaw(raw)
	// var things = input.Load()
	// var things = input.LoadSliceSliceString("")
	var things = input.LoadSliceString("\n")
	scores := []int{}
	for _, v := range things {
		s := processPass(v)
		scores = append(scores, s)
	}

	return output.High(scores)
}

func processPass(s string) int {
	// we don't actually care about the row and column
	//row := 0
	//column := 0
	score := 0

	for _, v := range s {
		j := 0
		if v == 'B' || v == 'R' {
			j = 1
		}

		score = score<<1 + j
	}

	return score
}

var raw = ``
