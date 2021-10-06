package minimaxsum

import (
	"sort"
)

// Puzzle ref: https://tinyurl.com/yg9zqgxt
func Puzzle(arr []int) (int, int) {
	sort.Ints(arr)

	// example: [1,2,3,4,5] -> 10, 14
	min := sumsElems(arr[0:4])
	max := sumsElems(arr[1:5])

	return min, max
}

func sumsElems(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}

	return sum
}
