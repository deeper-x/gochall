package minimaxsum

import (
	"log"
	"testing"
)

func TestMinMaxSum(t *testing.T) {
	min, max := Puzzle([]int{1, 2, 3, 4, 5})
	log.Println(min, max)
	if min != 10 || max != 14 {
		t.Errorf("min %d should be 10 and max %d should be 14", min, max)
	}
}
