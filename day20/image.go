package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	Y int
	X int
}

func imageAtPosition(images []*Image, x int, y int) *Image {
	for _, v := range images {
		if v.Point.X == x && v.Point.Y == y {
			return v
		}
	}
	panic("couldn't find image")
	return nil
}

func findCorners(images []*Image) (int, int, int, int) {
	minX, maxX, minY, maxY := 0, 0, 0, 0
	for _, v := range images {
		if v.Point.X > maxX {
			maxX = v.Point.X
		}
		if v.Point.Y > maxY {
			maxY = v.Point.Y
		}
		if v.Point.X < minX {
			minX = v.Point.X
		}
		if v.Point.Y < minY {
			minY = v.Point.Y
		}
	}

	return minX, maxX, minY, maxY
}

func alignImage(referenceImage *Image, attemptImage *Image) bool {

	rotations := []int{0, 90, 180, 270}
	flips := []int{FLIP_NONE, FLIP_VERTICAL}
	relativePositions := []Point{
		Point{1, 0},
		Point{-1, 0},
		Point{0, -1},
		Point{0, 1},
	}
	for _, dr := range rotations {
		for _, f := range flips {
			for _, dp := range relativePositions {
				// capture the original image
				original := attemptImage.Data
				originalRotate := attemptImage.Rotate
				originalFlipV := attemptImage.FlipVertical
				originalPoint := attemptImage.Point

				rotate(attemptImage, dr)
				flip(attemptImage, f)
				shift(attemptImage, referenceImage.Point, dp)

				if isAligned(referenceImage, attemptImage) {
					// set the attempt image in place
					attemptImage.Aligned = true
					return true
				}
				// reset the image
				attemptImage.Data = original
				attemptImage.Rotate = originalRotate
				attemptImage.FlipVertical = originalFlipV
				attemptImage.Point = originalPoint
			}
		}
	}
	return false
}

func isAligned(imageA, imageB *Image) bool {
	dy := imageB.Point.Y - imageA.Point.Y
	dx := imageB.Point.X - imageA.Point.X

	switch {
	case dy == 0 && dx == 1:
		for y := 0; y < 10; y++ {
			if imageA.Data[y][9] != imageB.Data[y][0] {
				return false
			}
		}

	case dy == 0 && dx == -1:
		for y := 0; y < 10; y++ {
			if imageA.Data[y][0] != imageB.Data[y][9] {
				return false
			}
		}

	case dy == 1 && dx == 0:
		for x := 0; x < 10; x++ {
			if imageA.Data[9][x] != imageB.Data[0][x] {
				return false
			}
		}

	case dy == -1 && dx == 0:
		for x := 0; x < 10; x++ {
			if imageA.Data[0][x] != imageB.Data[9][x] {
				return false
			}
		}
	default:
		panic(fmt.Sprintf("unhandled shift (%d, %d)", dy, dx))
	}

	return true
}

func shift(attemptImage *Image, p Point, dp Point) {
	attemptImage.Point.Y = p.Y + dp.Y
	attemptImage.Point.X = p.X + dp.X
}

const (
	FLIP_NONE     = 0
	FLIP_VERTICAL = 1
)

func flip(image *Image, t int) {
	switch t {
	case FLIP_NONE:
		// do nothing
	case FLIP_VERTICAL:
		image.FlipVertical = !image.FlipVertical
		flipVertically(image)
	default:
		panic("unhandled flip request")
	}
}

func rotate(image *Image, degrees int) {
	image.Rotate = (image.Rotate + degrees) % 360

	for i := 0; i < degrees; i += 90 {
		rotate90(image)
	}
}

func (img *Image) displayImage(showBorders bool) {

	fmt.Println("Tile:", img.ID)
	if showBorders == false {
		img.RemoveBorder()
		data := img.NoBorders

		for i, _ := range data {
			d := data[len(data)-1-i]
			for _, v := range d {
				fmt.Print(v)
			}
			fmt.Print("\n")
		}
		return
	}

	data := img.Data
	for i, _ := range data {
		fmt.Println(data[len(data)-1-i])
	}
}

type Image struct {
	ID        int
	Data      [10][10]string
	NoBorders [8][8]string

	Aligned bool
	Point   Point

	Rotate       int // (0, -90, -180, -270) - counterclockwise rotation
	FlipVertical bool
}

func (i Image) String() string {
	return fmt.Sprintf("Aligned: %v, Point (%d, %d), Rotate: %d, FlipV %v", i.Aligned, i.Point.X, i.Point.Y, i.Rotate, i.FlipVertical)
}

func (img *Image) RemoveBorder() {
	for i, v := range img.Data[1 : len(img.Data)-1] {
		for j, k := range v[1 : len(v)-1] {
			img.NoBorders[i][j] = k
		}
	}
}

func parseImage(s string) *Image {
	data := strings.Split(s, "\n")
	ids := strings.Trim(data[0], "Tile ")
	ids = strings.Trim(ids, ":")
	id, err := strconv.Atoi(ids)
	if err != nil {
		panic(err)
	}

	var rows [10][10]string
	for i, v := range data[1:] {
		var row [10]string
		for j, r := range v {
			row[j] = string(r)
		}
		rows[i] = row
	}
	img := Image{
		ID:   id,
		Data: rows,
	}
	return &img
}

func rotate90(image *Image) {
	//	[[0 -1][1 0]
	img := image.Data
	var newImg [10][10]string

	for y := range img {
		for x := range img {

			newImg[y][x] = img[9-x][y]
		}

	}
	image.Data = newImg
}

func flipVertically(image *Image) {
	img := image.Data
	var newImg [10][10]string
	for y := range img {
		for x := range img {
			newImg[y][x] = img[9-y][x]
		}
	}
	image.Data = newImg
}

var raw2 = `Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..`

var raw = `Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...`
