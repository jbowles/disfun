/*
minmax define functions for determining a Min and/or Max int32
*/
package disfun

import "math"

func MinFloat64(a ...float64) float64 {
	min := math.MaxFloat64
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

func MinInt32(a ...int) int {
	min := math.MaxInt32
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

func MaxInt32(a ...int) int {
	max := math.MinInt32 //int(0)
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}
