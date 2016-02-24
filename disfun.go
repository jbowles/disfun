// Package disfun implements various distance functions.
//
package disfun

import (
	"github.com/jbowles/disfun/Godeps/_workspace/src/github.com/gonum/matrix/mat64"
	"math"
)

const (
	EARTH_RADIUS = float64(6371)
	Substitution = float64(1)
	Insertion    = float64(1)
	Deletion     = float64(1)
)

// FloatsToMatrix creates a new mat64.Dense matrix from a slice of float64.
func FloatsToMatrix(floats []float64) *mat64.Dense {
	return mat64.NewDense(1, len(floats), floats)
}

// ZeroDense creates a new mat64.Dense matrix of zero values with dimenions of r and c.
func ZeroDense(r, c int) *mat64.Dense {
	var matrixPoints []float64
	for i := 0; i < (r * c); i++ {
		matrixPoints = append(matrixPoints, 0.0)
	}
	return mat64.NewDense(r, c, matrixPoints)
}

// MinFloat64 finds the minimum float64 value of a range of float64.
func MinFloat64(a ...float64) float64 {
	min := math.MaxFloat64
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

// MinInt32 finds the minimum int32 value of a range of int.
func MinInt32(a ...int) int {
	min := math.MaxInt32
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

// MaxINt32 finds the maximum int32 value of a range of int.
func MaxInt32(a ...int) int {
	max := math.MinInt32
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}
