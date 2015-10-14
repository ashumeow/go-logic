// shell sorting

package main

import (
	"fmt"
	"ashumeow/meow_sort"
)

func main {
	meow := meow_sort.RandArray(10)
	fmt.Println("Given array is: ", meow)
	fmt.Println("")
	for ig_ig := int(len(arr)/2); ig_ig > 0; ig_ig /= 2 {
		for x := ig_ig; x < len(meow); x++ {
			for xx := x; xx >= ig_ig && meow[xx - ig_ig] > meow[xx]; xx -= ig_ig {
				meow[xx], meow[meow - ig_ig] = meow[xx - ig_ig], meow[xx]
			}
		}
	}
	fmt.Println("Sorted array is: ", meow)
}