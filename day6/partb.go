package main

import (
	"fmt"

	"github.com/jnewmano/advent2020/input"
	"github.com/jnewmano/advent2020/output"
)

func main() {

	sum := parta()
	fmt.Println(sum)
}

func parta() interface{} {
	//input.SetRaw(raw)
	// var things = input.Load()
	// var things = input.LoadSliceSliceString("")
	var things = input.LoadSliceString("\n\n")

	var counts []int
	for _, v := range things {
		people, answers := process(v)
		score := 0
		for _, v := range answers {
			if v == people {
				score++
			}
		}
		counts = append(counts, score)
	}
	// var list = make([]int)
	count := 0
	for _, v := range counts {
		count += v
	}

	// output.High(list)
	return count
}

func process(s string) (int, map[string]int) {

	u := make(map[string]int)
	people := 1
	for _, v := range s {
		v := string(v)
		if v == "\n" {
			people++
		}
		if v == "" || v == "\n" {

			continue
		}

		u[v] = u[v] + 1
	}
	return people, u
}

var _ = output.High(nil)

var raw = `abc

a
b
c

ab
ac

a
a
a
a

b`
