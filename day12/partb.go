package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/jnewmano/advent2020/input"
)

func main() {

	//	input.SetRaw(raw)
	var things = input.LoadSliceString("")

	x, y := 0, 0
	wx, wy := 10, 1 // waypoint position, always relative to the ship's position

	for _, v := range things {
		d, _ := strconv.Atoi(v[1:])

		w := rune(v[0])
		switch w {
		case 'N', 'S', 'E', 'W':
			wx, wy = move(wx, wy, w, d)
		case 'L', 'R':
			wx, wy = rotateWaypoint(wx, wy, w, d)
		case 'F':
			// moves the ship wx, wy, d times
			for i := 0; i < d; i++ {
				x += wx
				y += wy
			}
		default:
			panic("unknown")
		}
	}

	fmt.Println(x, y, math.Abs(float64(x))+math.Abs(float64(y)))
}

func rotateWaypoint(wx, wy int, w rune, d int) (int, int) {
	if d <= 0 {
		return wx, wy
	}

	if w == 'L' {
		w = 'R'
		d = 360 - d
	}
	wx, wy = wy, -wx

	return rotateWaypoint(wx, wy, w, d-90)
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
