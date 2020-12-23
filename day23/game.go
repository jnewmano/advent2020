package main

import (
	"strconv"
)

var max = 0

func runGame(ref map[int]*Node, current *Node, turns int) {

	for i := 1; i < turns+1; i++ {

		cups := take3(current)

		insertCups(current, cups, ref)

		current = current.Next
	}
}

func insertCups(current *Node, cups *Node, ref map[int]*Node) {
	v := current.Value
	for {
		v--
		if v == 0 {
			v = max
		}

		if cups.Value != v && cups.Next.Value != v && cups.Next.Next.Value != v {
			break
		}

	}

	insert := ref[v]
	next := insert.Next
	insert.Next = cups
	cups.Next.Next.Next = next
}

func take3(current *Node) *Node {
	a := current.Next
	b := a.Next
	c := b.Next

	current.Next = c.Next

	return a
}

func addNode(ref map[int]*Node, last *Node, i int) *Node {

	n := Node{
		Value: i,
		Next:  nil,
	}
	ref[i] = &n

	if last != nil {
		last.Next = &n
	}

	return &n
}

type Node struct {
	Value int
	Next  *Node
}

func process(s string, fullGame bool) (map[int]*Node, *Node) {
	ref := make(map[int]*Node)

	var last *Node
	var first *Node

	for i, sv := range s {
		v, err := strconv.Atoi(string(sv))
		if err != nil {
			panic(err)
		}

		n := addNode(ref, last, v)
		if i == 0 {
			first = n
		}
		last = n

		if v > max {
			max = v
		}
	}

	if fullGame {
		for v := 10; v < 1000000+1; v++ {
			n := addNode(ref, last, v)
			last = n

			if v > max {
				max = v
			}
		}
	}

	// loop it all together
	last.Next = first

	return ref, first
}
