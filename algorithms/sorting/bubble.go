// bubble sorting

package main

import (
	"fmt"
	"ashumeow/meow_sort"
)

func main() {
	meow := meow_sort.RandArray(10)
	fmt.Println("Given array is: ",meow)
	fmt.Println("")
	temp := 0

	for x := 0; x < len(meow); x++ {
		for xx := 0; xx < len(meow) - 1; xx++ {
			if meow[xx] > meow[xx + 1] {
				temp = meow[xx]
				meow[xx] = meow[xx + 1]
				meow[xx + 1] = temp
			}
		}
	}
	fmt.Println("Sorted array is: ",meow)
}