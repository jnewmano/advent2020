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

	//	input.SetRaw(raw)
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

	// memory address decoder.
	// mask applies to address, not to the value

	// apply mask to address
	addresses := []int{address}

	for i, v := range mask {
		if v == 'X' {
			// expand out the possible addresses to write to
			// once with the address set and the other with the address cleared
			addresses = expandAddresses(addresses, 35-i)
		} else if v == '0' {
			continue
		} else {
			for ai, av := range addresses {
				addresses[ai] = setBit(av, 35-i)
			}
		}
	}

	for _, address := range addresses {
		data[address] = newValue
	}

}

func expandAddresses(addr []int, bit int) []int {
	var newAddrs []int

	for _, v := range addr {
		newAddrs = append(newAddrs, clearBit(v, bit))
		newAddrs = append(newAddrs, setBit(v, bit))
	}

	return newAddrs
}

func setBit(v int, bit int) int {

	tmp := 1 << bit
	v |= tmp

	return v
}

func clearBit(v int, bit int) int {

	tmp := ^(1 << bit)
	v &= tmp
	return v
}

var _ = output.High(nil)

var raw = `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`
