// binary search

package main

import "fmt"

func main() {
	value := 0
	meow := [10]int{1, 5, 100, 0, -100, 15, 4, 102, 30, 1000}
	fmt.Println(meow)

	// sorting the numbers
	temp := 0
	for x := 0; x < len(meow); x++ {
		for xx := 0; xx < len(meow)-1; xx++ {
			if meow[xx] > meow[xx+1] {
				temp = meow[xx]
				meow[xx] = meow[xx+1]
				meow[xx+1] = temp
			}
		}
	}
	fmt.Println(meow)
	left := 0
	right := len(meow) - 1

	if right < left {
		fmt.Println("Not found")
		return
	}
	// lookup
	for left <= right {
		middle := (left + right) / 2
		if meow[middle] == value {
			fmt.Println("Found at position: ", middle)
			return
		} else if meow[middle] < value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	fmt.Println("Not found")
}