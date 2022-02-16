package algorithms

import (
	"fmt"
	"sort"
)

//BinarySearch search a given number in a slice and return results
func BinarySearch(num int, si []int) bool {
	sort.Ints(si)
	start := 0
	end := len(si)

	fmt.Println(si)

	for start < end {
		middle := (start + end) / 2

		if si[middle] < num {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	if start == len(si) || si[start] != num {
		return false
	}

	return true
}
