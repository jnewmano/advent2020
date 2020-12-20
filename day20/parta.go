package main

import (
	"fmt"
	"os"

	"github.com/jnewmano/advent2020/input"
)

func main() {

	// input.SetRaw(raw)
	var rawImages = input.LoadSliceString("\n\n")

	var images = make([]*Image, len(rawImages))
	for i, v := range rawImages {
		images[i] = parseImage(v)
	}

	// assume the first image is in the desired transformation and position
	// say rotate:0, flip: 0, (dy,dx) = (0,0)
	// once a transformation and position are set, they are fixed

	currentImage := 0
	c := images[0]
	c.Aligned = true

	referenceImages := make(map[int]bool)
	for {
		// find and set orientation for all images that line up for where image at position i currently is
		for i, img := range images {
			if img.Aligned || i == currentImage {
				continue
			}
			current := images[currentImage]
			if alignImage(current, img) {
				img.Aligned = true
			}
		}
		referenceImages[currentImage] = true
		// find the next image to align off of

		isDone := false
		for _, v := range images {
			if v.Aligned == false {
				break
			}
		}
		if isDone {
			break
		}

		// look at more images
		for i := 0; i < len(images); i++ {
			next := (currentImage + i) % len(images)
			if next == currentImage {
				continue
			}

			if images[next].Aligned {
				currentImage = (currentImage + i) % len(images)
				break
			}
		}
		if len(referenceImages) >= len(images) {
			break
		}

	}

	for _, v := range images {
		if v.Aligned == false {
			fmt.Printf("ID: %d is not aligned\n", v.ID)
			os.Exit(1)
		}
	}

	// find min x, max x, min y, may y
	minX, maxX, minY, maxY := findCorners(images)

	{
		a := imageAtPosition(images, minX, minY)
		b := imageAtPosition(images, minX, maxY)
		c := imageAtPosition(images, maxX, minY)
		d := imageAtPosition(images, maxX, maxY)

		fmt.Println(a.ID * b.ID * c.ID * d.ID)
	}
}
