package main

import (
	"fmt"
)

func roughness(sea [][seaWidth]string) int {
	c := 0
	for _, v := range sea {
		for _, r := range v {
			if r == "#" {
				c++
			}
		}
	}
	return c
}

func stitchTogether(images []*Image) [][seaWidth]string {

	minX, _, minY, _ := findCorners(images)
	shiftImages(images, -minX, -minY)

	_, maxX, _, maxY := findCorners(images)

	sea := make([][seaWidth]string, maxY*8+8)

	for y := 0; y < maxY+1; y++ {
		for x := 0; x < maxX+1; x++ {
			img := imageAtPosition(images, x, y)
			// copy the image into the sea
			img.RemoveBorder()

			for dy := 0; dy < 8; dy++ {
				for dx := 0; dx < 8; dx++ {
					sea[y*8+dy][x*8+dx] = img.NoBorders[dy][dx]
				}
			}
		}
	}

	return sea
}

func displaySea(sea [][seaWidth]string) {
	for i, _ := range sea {
		d := sea[len(sea)-1-i]
		for _, v := range d {
			fmt.Print(v)
		}
		fmt.Print("|\n")
	}
}

func removeBorders(images []*Image) {
	for _, v := range images {
		v.RemoveBorder()
	}
}

func shiftImages(images []*Image, dx int, dy int) {
	for _, v := range images {
		v.Point.X += dx
		v.Point.Y += dy
	}
}

func copySea(sea [][seaWidth]string) [][seaWidth]string {
	newSea := make([][seaWidth]string, len(sea))
	for i, v := range sea {
		newSea[i] = v
	}
	return newSea
}
