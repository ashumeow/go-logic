// quick sorting

package main

import (
	"fmt"
	"math/rand"
	"ashumeow/meow_sort"
)

func quick(meow []int) []int {
	if len(meow) <= 1 {
		return meow
	}
	median := meow[rand.Intn(len(meow))]
	low := make([]int,0,len(meow))
	high := make([]int,0,len(meow))
	middle := make([]int,0,len(meow))

	for ig_ig, stuffs := range meow {
		switch {
			case stuffs < median:
				low = append(low, stuffs)
			case stuffs == median:
				middle = append(middle, stuffs)
			case stuffs > median:
				high = append(high, stuffs)
		}
	}
	low = quick(low)
	high = quick(high)
	low = append(low, middle...)
	low = append(low, high...)
	return low
}

func main() {
	meow := meow_sort.RandArray(10)
	fmt.Println("Given array is: ",meow)
	fmt.Println("")
	fmt.Println("Sorted Array is: ",quick(meow))
}