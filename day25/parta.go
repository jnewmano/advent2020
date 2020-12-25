package main

import (
	"fmt"
	"strconv"

	"github.com/jnewmano/advent2020/input"
)

func main() {

	//	input.SetRaw(raw)
	cardPK, doorPK := loadKeys()

	cardLoops := 0

	sn := 7
	cardLoops = loop(sn, cardPK)

	doorLoops := 0
	doorLoops = loop(sn, doorPK)

	fmt.Println("card loops:", cardLoops)
	fmt.Println("door loops:", doorLoops)

	fmt.Println(privateKey(cardLoops, doorPK))

}

func loop(subjectNumber int, key int) int {
	v := 1
	for i := 1; i < 100000000; i++ {
		v *= subjectNumber
		v %= 20201227
		if v == key {
			return i
		}
	}
	return -1
}

func privateKey(loops int, subjectNumber int) int {

	v := 1
	for i := 0; i < loops; i++ {
		v *= subjectNumber
		v %= 20201227
	}
	return v
}

func loadKeys() (int, int) {
	things := input.LoadSliceString("")
	a, err := strconv.Atoi(things[0])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(things[1])
	if err != nil {
		panic(err)
	}

	return a, b
}

var raw = `5764801
17807724`
