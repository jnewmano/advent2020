package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

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
	var things = input.LoadSliceString("")

	for i := 0; i < len(things); i++ {

		ins := []instruction{}
		for _, v := range things {
			i := parseLine(v)
			ins = append(ins, i)
		}

		v := ins[i]
		switch v.code {
		case "nop":
			v.code = "jmp"
			ins[i] = v
		case "jmp":
			v.code = "nop"
			ins[i] = v
		default:
			continue
		}

		err := run(ins)
		if err == nil {
			continue
		}
	}
	return 0
}

func run(ins []instruction) error {
	// Immediately before any instruction is executed a second time, what value is in the accumulator?

	runCount := make(map[int]bool)
	pc := 0
	acc := 0

	for {
		if pc == len(ins)-1 {
			fmt.Println("done")
			fmt.Println(acc)
			os.Exit(1)
		}

		i := ins[pc]

		fmt.Println(pc, i.code, i.a)

		if runCount[pc] {
			return fmt.Errorf("seen already")
		}
		runCount[pc] = true

		switch i.code {
		case "nop":
			pc++
		case "acc":
			pc++
			acc += i.a
		case "jmp":
			pc += i.a
		default:
			panic("unknown code")
		}

	}

	// output.High(list)
	// output.Sum(list)
	return nil
}

type instruction struct {
	code string
	a    int
}

func parseLine(s string) instruction {
	parts := strings.Split(s, " ")
	v, _ := strconv.Atoi(parts[1])
	ins := instruction{
		code: parts[0],
		a:    v,
	}
	return ins
}

func process(s string) {

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
