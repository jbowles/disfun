package disfun

import "testing"

func TestManhattanDistanceStructZero(t *testing.T) {
	x := floatsToMatrix([]float64{12, 45, 78, 23, 45, 97})

	distance := Manhattan(x, x)
	expected := 0.0
	if distance != expected {
		t.Log("Expected: ", expected, "but got ", distance)
		t.Fail()
	}
}
func TestManhattanDistanceStruct(t *testing.T) {
	x := floatsToMatrix([]float64{1, 2, 3})
	y := floatsToMatrix([]float64{2, 4, 6})

	distance := Manhattan(x, y)
	expected := 6.0
	if distance != expected {
		t.Log("Expected: ", expected, "but got ", distance)
		t.Fail()
	}
}
