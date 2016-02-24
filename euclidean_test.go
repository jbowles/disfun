package disfun

import (
	"github.com/jbowles/disfun/Godeps/_workspace/src/github.com/gonum/matrix/mat64"
	"math"
	"testing"
)

var m = mat64.NewDense(3, 1, []float64{1, 2, 3})
var n = mat64.NewDense(3, 1, []float64{2, 4, 6})

func BenchmarkEuclideanDistance(b *testing.B) {
	euclidean := NewEuclidean()
	for i := 0; i < b.N; i++ {
		euclidean.Distance(n, m)
	}
}

func TestEuclideanDistanceStruct(t *testing.T) {
	euclidean := NewEuclidean()
	var expected float64
	x := mat64.NewDense(3, 1, []float64{1, 2, 3})
	y := mat64.NewDense(3, 1, []float64{2, 4, 6})

	innerResult := euclidean.InnerProduct(x, y)
	expected = 28.0
	if innerResult != expected {
		t.Log("Expected: ", expected, "but got ", innerResult)
		t.Fail()
	}

	distance := euclidean.Distance(x, y)
	expected = math.Sqrt(14)
	if distance != expected {
		t.Log("Expected: ", expected, "but got ", distance)
		t.Fail()
	}
}

func TestEuclideanDistanceStructZero(t *testing.T) {
	euclidean := NewEuclidean()
	var expected float64
	x := floatsToMatrix([]float64{12, 45, 78, 23, 45, 97})

	innerResult := euclidean.InnerProduct(x, x)
	expected = 20216.0
	if innerResult != expected {
		t.Log("Expected: ", expected, "but got ", innerResult)
		t.Fail()
	}

	distance := euclidean.Distance(x, x)
	expected = 0.0
	if distance != expected {
		t.Log("Expected: ", expected, "but got ", distance)
		t.Fail()
	}
}

/*
func TestEuclideanDistanceZero(t *testing.T) {
	x := FloatsToMatrix([]float64{12, 45, 78, 23, 45, 97})
	y := FloatsToMatrix([]float64{12, 45, 78, 23, 45, 97})

	distance := EuclideanDistance(x, y)
	expected := 0.0
	DistanceTestCheck(t, distance, expected)
}
*/
