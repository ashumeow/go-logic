// D:\go\src\ashumeow\meow_sort
package meow_sort

import "math/rand"

func meowArray(ig int) []int {
	meow := make([]int, ig)

	for x := 0; x <= ig - 1; x++ {
		meow[x] = rand.Intn(ig)
	}
	return meow;
}