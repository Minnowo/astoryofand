package util

import (
	"math"
	"strconv"
)

const float64EqualityThreshold = 1e-5

func AlmostEqual32(a, b float32) bool {
	return math.Abs(float64(a)-float64(b)) <= float64EqualityThreshold
}

func AlmostEqual64(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func F32TS(value float32) string {
	return strconv.FormatFloat(float64(value), 'f', 2, 32)
}
