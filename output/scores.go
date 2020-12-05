package output

import "sort"

func High(scores []int) int {
	if len(scores) == 0 {
		return 0
	}

	n := Duplicate(scores)
	sort.Ints(n)
	return n[len(n)-1]
}

func Low(scores []int) int {
	n := Duplicate(scores)
	sort.Ints(n)
	return n[0]
}

func Duplicate(list []int) []int {
	n := make([]int, len(list))
	copy(n, list)
	return n
}
