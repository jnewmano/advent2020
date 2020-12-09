package main

import (
	"fmt"

	"github.com/jnewmano/advent2020/input"
	"github.com/jnewmano/advent2020/output"
)

func main() {

	sum := parta()
	fmt.Println(sum)
}

func parta() interface{} {
	// input.SetRaw(raw)
	// var things = input.Load()
	// var things = input.LoadSliceSliceString("")
	var things = input.LoadSliceInt("")

	fmt.Println(len(things))
	// var list = make([]int)

	m := contiguous(things, 32321523)
	if len(m) != 0 {
		fmt.Println(m)
		return output.High(m) + output.Low(m)
	}

	return nil
}

func contiguous(nums []int, match int) []int {
	for v, _ := range nums {
		for k, _ := range nums {
			sum := 0
			for i := v; i < k; i++ {
				sum += nums[i]
				if sum == match {
					return nums[v:k]
				}
			}
		}
	}
	return nil
}

var _ = output.High(nil)

var raw = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`
