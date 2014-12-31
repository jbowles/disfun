package disfun

import "testing"

func TestManhattanDistanceStructZero(t *testing.T) {
	man := NewManhattan()
	x := FloatsToMatrix([]float64{12, 45, 78, 23, 45, 97})

	distance := man.Distance(x, x)
	expected := 0.0
	if distance != expected {
		t.Log("Expected: ", expected, "but got ", distance)
		t.Fail()
	}
}
func TestManhattanDistanceStruct(t *testing.T) {
	man := NewManhattan()
	x := FloatsToMatrix([]float64{1, 2, 3})
	y := FloatsToMatrix([]float64{2, 4, 6})

	distance := man.Distance(x, y)
	expected := 6.0
	if distance != expected {
		t.Log("Expected: ", expected, "but got ", distance)
		t.Fail()
	}
}
