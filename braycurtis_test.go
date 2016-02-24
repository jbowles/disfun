package disfun

import "testing"

func TestBrayCurtisDistanceStructZero(t *testing.T) {
	x := floatsToMatrix([]float64{12, 45, 78, 23, 45, 97})

	distance := BrayCurtis(x, x)
	expected := 0.0
	if distance != expected {
		t.Log("Expected: ", expected, "but got ", distance)
		t.Fail()
	}
}
func TestBrayCurtisDistanceStruct(t *testing.T) {
	var expected float64
	x := floatsToMatrix([]float64{1, 2, 3})
	y := floatsToMatrix([]float64{2, 4, 6})

	distance := BrayCurtis(x, y)
	expected = (1.0 / 3.0)
	if distance != expected {
		t.Log("Expected: ", expected, "but got ", distance)
		t.Fail()
	}
}
