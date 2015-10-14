// merge sorting

package main

import (
	"fmt"
	"ashumeow/meow_sort"
)

func merge(left, right []int) []int {
	out := make([]int,0,len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(out, right...)
		}
		if len(right) == 0 {
			return append(out, left...)
		}
		if left[0] <= right[0] {
			out = append(out, left[0])
			left = left[1:]
		} else {
			out = append(out, right[0])
			right = right[1:]
		}
	}
	return out
}

func merge_ig(meow []int) []int {
	if len(meow) <= 1 {
		return meow
	}
	middle := len(meow) / 2
	left := merge_ig(meow[:middle])
	right := merge_ig(meow[:middle])
	return merge(left, right)
}

func main() {
	meow := meow_sort.RandArray(10)
	fmt.Println("Given array is: ",meow)
	fmt.Println("")
	fmt.Println("Sorted array is: ",merge_ig(meow))
}