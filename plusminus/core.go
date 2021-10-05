package plusminus

import (
	"errors"
	"fmt"
)

// Puzzle challenge - source: https://tinyurl.com/yzsyscrx
func Puzzle(arr []int64) (string, error) {
	var pos, neg, zero int
	var aLen = len(arr)
	var res = ""

	if aLen > 100 {
		return res, errors.New("array length must be <= 100")
	}

	for _, v := range arr {
		if v < -100 || v > 100 {
			return res, errors.New("array elements must be between -100 and 100")
		}

		if v > 0 {
			pos++
		} else if v == 0 {
			zero++
		} else {
			neg++
		}
	}

	tPos := float64(pos) / float64(aLen)
	tNeg := float64(neg) / float64(aLen)
	tZero := float64(zero) / float64(aLen)

	res = fmt.Sprintf("%.6f\n%.6f\n%.6f\n", tPos, tNeg, tZero)

	return res, nil
}
