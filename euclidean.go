/*
Euclidean finds the "ordinary" distance between vectors [u,v] given by the Pythagorean formula (sum of squares of all points):

	dist((u, v) = dist(v, u)) = √(v_1 - u_1)² + (v_2 - u_2)² + ... + (v_n - u_n)²

http://reference.wolfram.com/language/ref/EuclideanDistance.html
*/

package disfun

import (
	"github.com/gonum/matrix/mat64"
	"math"
)

type Euclidean struct{}

func NewEuclidean() *Euclidean {
	return &Euclidean{}
}

// InnerProduct computes a Eucledian inner product.
func (e *Euclidean) InnerProduct(u, v *mat64.Dense) (result float64) {
	result = u.Dot(v)
	return
}

func (e *Euclidean) Distance(u, v *mat64.Dense) float64 {
	subVec := mat64.NewDense(0, 0, nil)
	subVec.Sub(u, v)

	result := e.InnerProduct(subVec, subVec)
	return math.Sqrt(result)
}

/*
 Not sure why this doesn't work!!
func EuclideanDistance(u, v *mat64.Dense) float64 {
	subVec := mat64.NewDense(0, 0, nil)
	subVec.Sub(u, v)

	result := u.Dot(v)
	//fmt.Printf("%v\n", result)
	return math.Sqrt(result)
}
*/
