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
	acc := 0

	var things = input.LoadSliceString("")

	for i := 0; i < len(things); i++ {
		c := LoadComputer(true)

		fmt.Println(i)

		v := c.ins[i]
		switch v.Code {
		case "nop":
			v.Code = "jmp"
		case "jmp":
			v.Code = "nop"
		default:
			continue
		}

		err := c.Run()
		if err != nil {
			continue
		}
		acc = c.acc
		break
	}
	return acc
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
