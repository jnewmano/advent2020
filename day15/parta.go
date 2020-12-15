package main

import (
	"fmt"

	"github.com/jnewmano/advent2020/input"
)

// number, turn it was said on
var nums = make(map[int]int)

func main() {

	// input.SetRaw(raw)
	var things = input.LoadSliceInt(",")
	last := 0
	for i, v := range things {
		last = sayNumber(i, v)
	}

	var prev int
	for i := len(things); i < 2020; i++ {
		prev = last
		last = sayNumber(i, last)

	}

	fmt.Println(prev)
}

func sayNumber(i int, num int) int {

	cur, ok := nums[num]
	nums[num] = i + 1
	if ok {
		return i + 1 - cur
	}

	return 0

}

var raw = `0,3,6`
