package main

import (
	"fmt"
	"sort"

	"github.com/jnewmano/advent2020/input"
	"github.com/jnewmano/advent2020/output"
)

func main() {

	sum := parta()
	fmt.Println(sum)
}

func parta() interface{} {
	// input.SetRaw(raw1)
	// var things = input.Load()
	// var things = input.LoadSliceSliceString("")
	var things = input.LoadSliceInt("")

	sort.Ints(things)
	things = append(things, things[len(things)-1]+3)

	// var list = make([]int)
	diff1 := 0
	diff3 := 0
	last := 0
	for _, v := range things {
		diff := v - last

		if diff == 1 {
			diff1++
		} else if diff == 3 {
			diff3++
		}
		last = v
	}

	fmt.Println(diff1)
	fmt.Println(diff3)

	// output.High(list)
	// output.Low(list)
	// output.Sum(list)
	return nil
}

func process(s string) {

}

var _ = output.High(nil)

var raw1 = `16
10
15
5
1
11
7
19
6
12
4
`

var raw = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`
