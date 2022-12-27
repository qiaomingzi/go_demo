package simplemath

import "math"

func Sqrt(v int) int {
	r := math.Sqrt(float64(v))
	return int(r)
}
