package main

import (
	"fmt"
	"strconv"

	"github.com/jnewmano/advent2020/input"
)

func main() {

	//input.SetRaw(raw)
	var things = input.Load()
	ref, current := process(things, false)

	runGame(ref, current, 100)

	idx := ref[1]
	next := idx.Next
	s := ""
	for {
		s += strconv.Itoa(next.Value)
		next = next.Next
		if next == idx {
			break
		}
	}
	fmt.Println(s)
}

var raw = `389125467`
