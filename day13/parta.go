package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jnewmano/advent2020/input"
	"github.com/jnewmano/advent2020/output"
)

func main() {

	//	input.SetRaw(raw)
	// var things = input.Load()
	// var things = input.LoadSliceSliceString("")
	var things = input.LoadSliceString("")

	earliest, _ := strconv.Atoi(things[0])
	buses := strings.Split(things[1], ",")

	var a, id int
	for _, v := range buses {
		if v == "x" {
			continue
		}
		j, _ := strconv.Atoi(v)
		m := ((earliest / j) + 1) * j
		if m < a || a == 0 {
			a = m
			id, _ = strconv.Atoi(v)
		}
	}

	fmt.Println(id, a-earliest, id*(a-earliest))
	// output.High(list)
	// output.Sum(list)

}

func process(s string) {

}

var _ = output.High(nil)

var raw = `939
7,13,x,x,59,x,31,19`
