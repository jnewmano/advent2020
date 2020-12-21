package main

func intersection(a []string, b []string) []string {
	list := []string{}
	for _, v := range a {
		for _, w := range b {
			if v == w {
				list = append(list, w)
			}
		}
	}
	return list
}

// remove items in b from a
func remove(a []string, b []string) []string {

	list := []string{}
	for _, v := range a {
		ignore := false
		for _, w := range b {
			if v == w {
				ignore = true
				break
			}
		}
		if ignore == false {
			list = append(list, v)
		}
	}
	return list
}
