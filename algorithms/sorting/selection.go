// selection sort

package main

import (
	"fmt"
	"ashumeow/meow_sort"
)

func main() {
	meow := meow_sort.RandArray(10)
	fmt.Println("Given array is: ", meow)
	fmt.Println("")

	var minimum int = 0
	var temp int = 0

	for x := 0; x < len(meow_sort); x++ {
		minimum = x
		for xx := x + 1; xx < len(meow); xx++ {
			if meow[xx] < meow[minimum] {
				minimum = xx
			}
		}
		temp = meow[x]
		meow[x] = meow[minimum]
		meow[minimum] = temp
	}
	fmt.Println("Sorted array is: ", meow)
}