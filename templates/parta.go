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
	// input.SetRaw(raw)
	// var things = input.Load()
	// var things = input.LoadSliceSliceString("")
	var things = input.LoadSliceString("")

	for _, v := range things {
		process(v)
	}
	// var list = make([]int)

	// output.High(list)
	return nil
}

func process(s string) {

}

var _ = output.High(nil)

var raw = ``
