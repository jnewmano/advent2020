package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jnewmano/advent2020/input"
)

// var things = input.LoadRaw()
var things = input.LoadSliceString("")

/*[]string{
	`1-3 a: abcde`,
	`1-3 b: cdefg`,
	`2-9 c: ccccccccc`,
}*/

func main() {
	sum := parta()
	fmt.Println(sum)
}

// find the two entries that sum to 2020 and then multiply those two numbers together.
func parta() int {

	count := 0
	for _, v := range things {
		low, high, letter, password := parseLine(v)
		if (string(password[low-1]) == letter) != (string(password[high-1]) == letter) {
			count++
		}
	}
	return count
}

func countLetters(password string, letter string) int {
	i := 0
	for _, v := range password {
		if string(v) == letter {
			i++
		}
	}

	return i
}

func parseLine(v string) (int, int, string, string) {
	parts := strings.Split(v, ":")
	password := strings.TrimSpace(parts[1])

	parts = strings.Split(parts[0], " ")
	letter := strings.TrimSpace(parts[1])

	parts = strings.Split(parts[0], "-")
	low, _ := strconv.Atoi(parts[0])
	high, _ := strconv.Atoi(parts[1])

	return low, high, letter, password
}
