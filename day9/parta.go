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
	//	input.SetRaw(raw)
	// var things = input.Load()
	// var things = input.LoadSliceSliceString("")
	var things = input.LoadSliceInt("")

	p := 25
	fmt.Println(len(things))
	// var list = make([]int)
	for i, _ := range things {
		if i < p {
			continue
		}
		fmt.Println(things[i])
		if verify(things, p, i) == false {
			return things[i]
		}
	}

	// output.High(list)
	// output.Sum(list)
	return nil
}

func verify(nums []int, p int, idx int) bool {

	for _, v := range nums[idx-p : idx] {
		for _, k := range nums[idx-p : idx] {
			if v == k {
				continue
			}

			fmt.Println(v, k, nums[idx])
			if v+k == nums[idx] {
				return true
			}
		}
	}

	return false
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
