package disfun

import (
	"github.com/jbowles/disfun/Godeps/_workspace/src/github.com/gonum/matrix"
	"github.com/jbowles/disfun/Godeps/_workspace/src/github.com/gonum/matrix/mat64"
	"math"
)

// Manhattan (also called TaxiCab) finds the distance by the absolute sum of all differences of points in vectors [u,v]:
//
//	Example:
//	 dist((u, v) = dist(v, u)) = |v_1 - u_1| + |v_2 - u_2| + ... + |v_n - u_n|
//
//	References:
//	http://mathworld.wolfram.com/TaxicabMetric.html
//	http://demonstrations.wolfram.com/TaxicabGeometry/
func Manhattan(x, y *mat64.Dense) (result float64) {
	r1, c1 := x.Dims()
	r2, c2 := y.Dims()
	if r1 != r2 || c1 != c2 {
		panic(matrix.ErrShape)
	}

	for i := 0; i < r1; i++ {
		for j := 0; j < c1; j++ {
			result += math.Abs(x.At(i, j) - y.At(i, j))
		}
	}

	return
}
