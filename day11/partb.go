package main

import (
	"fmt"

	"github.com/jnewmano/advent2020/input"
)

func main() {

	layout := loadMap(true)

	updated := 0

	for {
		updated, layout = process(layout)
		if updated == 0 {
			break
		}
	}
	fmt.Println(countOccupied(layout))
}

func process(l [][]string) (int, [][]string) {
	update := loadMap(false)
	updated := 0
	for i, v := range l {
		for j, x := range v {
			c := seeOccupied(i, j, l)

			setTo := x
			switch x {
			case "L":
				if c == 0 {
					setTo = "#"
				}
			case "#":
				if c >= 5 {
					setTo = "L"
				}
			}

			if setTo != x {
				updated++
			}
			update[i][j] = setTo
		}
	}
	return updated, update
}

func countOccupied(l [][]string) int {
	c := 0
	for _, v := range l {
		for _, x := range v {
			if x == "#" {
				c++
			}
		}
	}
	return c
}

func seeOccupied(a int, b int, l [][]string) int {
	c := 0
	c += see(a, b, 1, 0, l)
	c += see(a, b, -1, 0, l)
	c += see(a, b, 0, 1, l)
	c += see(a, b, 0, -1, l)
	c += see(a, b, -1, 1, l)
	c += see(a, b, -1, -1, l)
	c += see(a, b, 1, -1, l)
	c += see(a, b, 1, 1, l)
	return c
}

func see(x, y, dx, dy int, l [][]string) int {
	width := len(l[0])
	height := len(l)

	for {
		x += dx
		y += dy
		if x < 0 || x >= height || y < 0 || y >= width {
			return 0
		}
		v := l[x][y]
		if v == "#" {
			return 1
		}
		if v == "L" {
			return 0
		}
	}
}

func loadMap(set bool) [][]string {

	var things = input.LoadSliceString("")

	// convert
	layout := make([][]string, len(things))
	for i, _ := range layout {
		layout[i] = make([]string, len(things[0]))
		for j, k := range things[i] {
			layout[i][j] = string(k)
		}
	}
	return layout
}
