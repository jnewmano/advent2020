package main

import (
	"fmt"

	"github.com/jnewmano/advent2020/input"
)

func main() {
	// input.SetRaw(raw)

	// var things = input.Load()
	// var things = input.LoadSliceSliceString("")

	sum := parta()
	fmt.Println(sum)
}

// find the two entries that sum to 2020 and then multiply those two numbers together.
func parta() interface{} {
	var things = input.LoadSliceString("")

	// start by counting all the trees you would encounter for the slope right 3, down 1:
	width := len(things[0])
	height := len(things)
	fmt.Println(width)
	fmt.Println(height)

	dx := 3
	dy := 1
	x := 0
	y := 0

	treeCount := 0
	for {
		if string(things[y][x]) == "#" {
			treeCount++
		}
		y += dy
		x += dx
		if y == height {
			break
		}
		if x >= width {
			x -= width
		}
	}

	return treeCount
}

var raw = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`
