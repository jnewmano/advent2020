package main

import (
	"github.com/jnewmano/advent2020/input"
	"github.com/jnewmano/advent2020/output"
)

func main() {

	// input.SetRaw(raw)
	// var things = input.Load()
	// var things = input.LoadSliceSliceString("")
	var things = input.LoadSliceString("")

	// var list = make([]int)
	for _, v := range things {
		process(v)
	}

	// output.High(list)
	// output.Sum(list)

}

func process(s string) {

}

var _ = output.High(nil)

var raw = ``
