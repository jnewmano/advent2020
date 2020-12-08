package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jnewmano/advent2020/input"
)

type Computer struct {
	pc  int
	acc int

	ins []*Instruction

	DetectLoops bool
}

func LoadComputer(detectLoops bool) *Computer {
	var things = input.LoadSliceString("")
	ins := []*Instruction{}
	for _, v := range things {
		i := parseLine(v)
		ins = append(ins, &i)
	}

	c := Computer{
		ins:         ins,
		DetectLoops: detectLoops,
	}

	return &c
}

func (c *Computer) Run() error {

	runCount := make(map[int]int)

	for {
		if c.pc == len(c.ins)-1 {
			return nil
		}

		i := c.ins[c.pc]

		if c.DetectLoops && runCount[c.pc] > 0 {
			return fmt.Errorf("seen already")
		}
		runCount[c.pc] = runCount[c.pc] + 1

		newPC := c.pc + 1
		switch i.Code {
		case "nop":
		case "acc":
			c.acc += i.A
		case "jmp":
			newPC = c.pc + i.A
		default:
			panic("unknown code")
		}
		c.pc = newPC

	}

	return nil
}

type Instruction struct {
	Code string
	A    int
}

func parseLine(s string) Instruction {
	parts := strings.Split(s, " ")
	v, _ := strconv.Atoi(parts[1])
	ins := Instruction{
		Code: parts[0],
		A:    v,
	}
	return ins
}
