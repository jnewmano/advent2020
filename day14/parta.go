package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jnewmano/advent2020/input"
	"github.com/jnewmano/advent2020/output"
)

var (
	data = make(map[int]int)
	mask string
)

func main() {

	//input.SetRaw(raw)
	var things = input.LoadSliceString("")

	for _, v := range things {
		process(v)
	}

	sum := 0
	for _, v := range data {
		sum += v
	}
	fmt.Println(sum)
}

// The current bitmask is applied to values immediately before they are written to memory

func process(s string) {
	parts := strings.Split(s, " = ")
	if strings.HasPrefix(parts[0], "mask") {
		mask = parts[1]
		return
	}

	newValue, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	d := parts[0][4 : len(parts[0])-1]
	address, err := strconv.Atoi(d)
	if err != nil {
		panic(err)
	}

	// apply
	for i, v := range mask {
		if v == 'X' {
			continue
		}
		// set current bit to zero
		if v == '0' {
			tmp := ^(1 << (35 - i))
			newValue &= tmp
		} else {
			tmp := 1 << (35 - i)
			newValue |= tmp
		}
	}

	data[address] = newValue
}

var _ = output.High(nil)

var raw = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`
