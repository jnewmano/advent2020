package main

import (
	"fmt"

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
	c := LoadComputer(true)
	_ = c.Run()
	fmt.Println(c.acc)

	return nil
}

var _ = output.High(nil)

var raw = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`
