// heap sorting
package main

import (
	"fmt"
	"ashumeow/meow_sort"
)

func heap(meow []int, x int, meowArrLen int) []int {
	end := false
	temp := 0
	maxBaby := 0

	for(x * 2 + 1 < meowArrLen) && (!end) {
		if x * 2 + 1 == meowArrLen - 1 {
			maxBaby = x * 2 + 1
		} else if (meow[x * 2 + 1] > meow[x * 2 + 2]) {
			maxBaby = x * 2 + 1
		} else {
			maxBaby = x * 2 + 2
		}
		if meow[x] < meow[maxBaby] {
			temp = meow[x]
			meow[x] = meow[maxBaby]
			meow[maxBaby] = temp
			x = maxBaby
		} else {
			end = true
		}
	}
	return meow
}

func main() {
	meow := meow_sort.RandArray(10)
	fmt.Println("Given array is: ", meow)
	fmt.Println("")
	x := 0
	temp := 0
	for x = len(meow) / 2 - 1; x >= 0; x-- {
		meow = heap(meow, x, len(meow))
	}
	for x = len(meow) - 1; x >= 1; x-- {
		temp = meow[0];
		meow[0] = meow[x];
		meow[x] = temp;
		meow = heap(meow, 0, x);
	}
	fmt.Println("Sorted array is: ", meow)
}