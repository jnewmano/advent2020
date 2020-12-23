package main

import (
	"fmt"

	"github.com/jnewmano/advent2020/input"
)

func main() {

	// input.SetRaw(raw)
	var things = input.Load()
	ref, current := process(things, true)

	runGame(ref, current, 10000000)

	idx := ref[1]

	a := idx.Next.Value
	b := idx.Next.Next.Value

	fmt.Println(a, b, a*b)
}

var raw = `389125467`
