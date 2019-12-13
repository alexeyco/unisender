package test

import "math"

const eps float64 = 0.00000001

func EqualFloat64(expected, given float64) bool {
	if math.Abs(expected-given) < eps {
		return true
	}

	return false
}
