// Package disfun implements various distance functions.
package disfun

import (
	"github.com/gonum/matrix/mat64"
)

func FloatsToMatrix(floats []float64) *mat64.Dense {
	return mat64.NewDense(1, len(floats), floats)
}

func ZeroDense(rows, columns int) *mat64.Dense {
	var matrixPoints []float64
	for i := 0; i < (rows * columns); i++ {
		matrixPoints = append(matrixPoints, 0.0)
	}
	return mat64.NewDense(rows, columns, matrixPoints)
}
