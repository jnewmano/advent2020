package main

import (
	"fmt"
	"strconv"

	"github.com/jnewmano/advent2020/input"
	"github.com/jnewmano/advent2020/simplemath"
)

func main() {

	//input.SetRaw(raw)
	var things = input.LoadSliceString("")

	dir := 'E'
	x, y := 0, 0

	for _, v := range things {
		d, _ := strconv.Atoi(v[1:])

		w := rune(v[0])
		switch w {
		case 'N', 'S', 'E', 'W':
			x, y = move(x, y, w, d)
		case 'L', 'R':
			dir = rotate(dir, w, d)
		case 'F':
			x, y = move(x, y, dir, d)
		default:
			panic("unknown")
		}
	}

	fmt.Println(x, y, simplemath.Abs(x)+simplemath.Abs(y))
}

func rotate(dir rune, w rune, d int) rune {
	if d == 0 {
		return dir
	}

	if w == 'L' {
		d = 360 - d
		w = 'R'
	}

	switch dir {
	case 'N':
		dir = 'E'
	case 'E':
		dir = 'S'
	case 'S':
		dir = 'W'
	case 'W':
		dir = 'N'
	}

	return rotate(dir, w, d-90)
}
func move(x, y int, a rune, d int) (int, int) {
	switch a {
	case 'N':
		y += d
	case 'S':
		y -= d
	case 'E':
		x += d
	case 'W':
		x -= d
	default:
		panic("unknown")
	}
	return x, y
}

var raw = `F10
N3
F7
R90
F11`
