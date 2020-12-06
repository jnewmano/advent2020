package main

import (
	"fmt"
	"unicode"

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
		answers := process(v)
		counts = append(counts, len(answers))
	}

	return output.Sum(counts)
}

func process(s string) map[rune]rune {

	u := make(map[rune]rune)
	for _, v := range s {
		if unicode.IsSpace(v) {
			continue
		}

		u[v] = v
	}
	return u
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
