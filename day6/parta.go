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
		answers := process(v)
		counts = append(counts, len(answers))
	}
	// var list = make([]int)
	count := 0
	for _, v := range counts {
		count += v
	}

	// output.High(list)
	return count
}

func process(s string) map[string]string {

	u := make(map[string]string)
	for _, v := range s {
		v := string(v)
		if v == "" || v == "\n" {
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
