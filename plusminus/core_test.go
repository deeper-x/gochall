package plusminus

import (
	"errors"
	"testing"
)

func TestPlusMinus(t *testing.T) {
	arr1 := []int64{1, -1, 0, 2, -3, 2}
	res, err := Puzzle(arr1)

	expected := `0.500000
0.333333
0.166667
`
	if err != nil {
		t.Error(err)
	}

	if res != expected {
		t.Errorf("result %v is not as expected", res)
	}
}

func TestPlusMinusOutRange(t *testing.T) {
	arr1 := []int64{1, -101, 0, 2, -3, 2}
	_, err := Puzzle(arr1)

	if err == nil {
		t.Error(errors.New("one element is -101, it's out of range and it cannot be admitted"))
	}

}
