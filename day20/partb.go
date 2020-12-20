package main

import (
	"fmt"
	"os"

	"github.com/jnewmano/advent2020/input"
)

const seaWidth = 12 * 8

func main() {

	//	input.SetRaw(raw)
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

	rotate(c, 270)
	flip(c, FLIP_VERTICAL)

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

	// make sure everything lines up correctly
	for _, v := range images {
		if v.Aligned == false {
			fmt.Printf("ID: %d is not aligned\n", v.ID)
			os.Exit(1)
		}
	}

	sea := stitchTogether(images)

	markedSea, found := findMonsters(sea)
	if found == 0 {
		fmt.Println("no monsters found")
		os.Exit(1)
	}

	fmt.Printf("Found %d monsters\n", found)
	fmt.Println("Sea roughness is:", roughness(markedSea))
}
