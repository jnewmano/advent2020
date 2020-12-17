package main

import (
	"fmt"

	"github.com/jnewmano/advent2020/input"
)

type Dim struct {
	x int
	y int
	z int
}

var grid = make(map[Dim]bool)

func main() {

	// input.SetRaw(raw)

	var things = input.LoadSliceString("")
	for i, v := range things {
		for j, k := range v {
			if k == '#' {
				d := Dim{
					x: j,
					y: i,
					z: 0,
				}
				grid[d] = true
			}
		}
	}

	for i := 0; i < 6; i++ {
		dest := make(map[Dim]bool)

		for k, _ := range grid {
			process(dest, k)
			ns := neighbors(k)
			for _, v := range ns {
				process(dest, v)
			}
		}

		grid = dest
	}

	active := countActive()
	fmt.Println(active)
}

func process(dest map[Dim]bool, node Dim) {

	activeCount := 0
	for _, v := range neighbors(node) {
		v, ok := grid[v]
		if !ok {
			continue
		}
		if !v {
			continue
		}
		activeCount++
	}

	// If a cube is active and exactly 2 or 3
	// of its neighbors are also active,
	// the cube remains active
	if grid[node] {
		if activeCount == 2 || activeCount == 3 {
			dest[node] = true
		}

	} else {
		// If a cube is inactive but exactly 3 of its
		//  neighbors are active, the cube becomes active
		if activeCount == 3 {
			dest[node] = true
		}
	}
}

func neighbors(node Dim) []Dim {
	deltas := [][]int{}

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				if i == 0 && j == 0 && k == 0 {
					continue
				}
				deltas = append(deltas, []int{i, j, k})
			}
		}
	}

	dims := []Dim{}
	for _, v := range deltas {
		d := Dim{
			x: v[0] + node.x,
			y: v[1] + node.y,
			z: v[2] + node.z,
		}

		dims = append(dims, d)
	}
	return dims
}

func countActive() int {
	c := 0
	for _, v := range grid {
		if v {
			c++
		}
	}
	return c
}

var raw = `.#.
..#
###`
