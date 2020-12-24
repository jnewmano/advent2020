package main

import (
	"fmt"

	"github.com/jnewmano/advent2020/input"
)

func main() {

	//	input.SetRaw(raw)
	var things = input.LoadSliceString("")

	floor := make(map[Point]int)

	for _, v := range things {
		rules := process(v)

		start := Point{
			Array:  0,
			Row:    0,
			Column: 0,
		}
		for _, v := range rules {
			start = move(start, v)
		}

		floor[start] = (floor[start] + 1) % 2
	}

	fmt.Printf("Day %d: %d\n", 0, CountBlackTiles(floor))
	var i int
	for i = 1; i <= 100; i++ {
		floor = Day(floor)
	}
	fmt.Printf("Day %d: %d\n", i-1, CountBlackTiles(floor))
}

func Day(floor map[Point]int) map[Point]int {
	newFloor := make(map[Point]int)

	for p := range floor {
		points := getNeighbors(p)
		points = append(points, p)
		for _, p := range points {
			c := countSetNeighbors(floor, p)
			if floor[p] == 1 {
				//
				if c == 0 || c > 2 {
					// flip to white
				} else {
					// keep as a black tile
					newFloor[p] = 1
				}
			} else {
				if c == 2 {
					newFloor[p] = 1
				}
				// otherwise it should be white and ignored
			}
		}
	}
	return newFloor
}

func countSetNeighbors(floor map[Point]int, p Point) int {
	neighbors := getNeighbors(p)
	count := 0
	for _, v := range neighbors {
		if floor[v] == 1 {
			count++
		}
	}
	return count
}

func move(start Point, dir string) Point {
	switch dir {
	case "e":
		start.Column++
	case "se":
		start.Array++
		if start.Array > 1 {
			start.Array = 0
			start.Row++
			start.Column++
		}

	case "sw":
		start.Array++
		if start.Array > 1 {
			start.Array = 0
			start.Row++
		} else {
			start.Column--
		}
	case "w":
		start.Column--
	case "ne":
		start.Array--
		if start.Array < 0 {
			start.Array = 1
			start.Row--
		} else {
			start.Column++
		}
	case "nw":
		start.Array--
		if start.Array < 0 {
			start.Array = 1
			start.Row--
			start.Column--
		}

	default:
		panic("unknown direction")
	}
	return start
}

func CountBlackTiles(floor map[Point]int) int {
	count := 0
	count2 := 0
	for _, v := range floor {
		if v == 1 {
			count++
		} else {
			count2++
		}
	}
	return count
}

func getNeighbors(point Point) []Point {
	dirs := []string{"e", "se", "sw", "w", "nw", "ne"}
	points := []Point{}
	for _, v := range dirs {
		points = append(points, move(point, v))
	}
	return points
}

// Use Hexagonal Coordinate System
type Point struct {
	Array  int
	Row    int
	Column int
}

func process(s string) []string {
	last := ""
	rules := []string{}

	for _, v := range s {
		if (last == "n" || last == "s") && (v == 'e' || v == 'w') {
			last = last + string(v)
			rules = append(rules, last)
			last = ""
			continue
		}
		if v == 'e' || v == 'w' {
			rules = append(rules, string(v))
			last = ""
			continue
		}

		if last != "" {
			rules = append(rules, last)
		}

		last = string(v)

	}

	if last != "" {
		rules = append(rules, last)
	}

	for _, v := range rules {
		switch v {
		case "e", "se", "sw", "w", "nw", "ne":
		default:
			panic("unknown direction: " + v)
		}
	}
	return rules
}

var raw = `sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`
