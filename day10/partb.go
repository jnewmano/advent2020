package main

import (
	"fmt"
	"sort"

	"github.com/jnewmano/advent2020/input"
	"github.com/jnewmano/advent2020/output"
)

func main() {

	answer := partb()
	fmt.Println(answer)
}

func partb() interface{} {
	// input.SetRaw((raw2))
	var things = input.LoadSliceInt("")
	things = append(things, 0) // add the starting node
	sort.Ints(things)

	paths := map[int]int{
		0: 1, // seed the starting node with 1
	}

	for _, v := range things {
		for _, k := range things {
			if k <= v {
				continue
			}
			if k-v > 3 { // if it's too far to the next one, don't add it as a valid path
				continue
			}

			// add number of existing paths to get to the neighbor node
			paths[k] += paths[v]
		}
	}

	// return the number of paths to the last adapter
	last := output.High(things)
	return paths[last]
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
