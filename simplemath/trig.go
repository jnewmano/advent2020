package simplemath

import "fmt"

func Sin(theta int) int {
	return Cos(theta - 90)
}

func Cos(theta int) int {
	if theta < 0 {
		theta = theta + 360
	}

	theta = theta % 360

	switch theta {
	case 0:
		return 1
	case 90:
		return 0
	case 180:
		return -1
	case 270:
		return 0
	}
	panic(fmt.Sprintf("cos(%d) not handled", theta))
}
