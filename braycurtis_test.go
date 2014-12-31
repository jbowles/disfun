package disfun

import "testing"

func TestBrayCurtisDistanceStructZero(t *testing.T) {
	man := NewBrayCurtis()
	x := FloatsToMatrix([]float64{12, 45, 78, 23, 45, 97})

	distance := man.Distance(x, x)
	expected := 0.0
	if distance != expected {
		t.Log("Expected: ", expected, "but got ", distance)
		t.Fail()
	}
}
func TestBrayCurtisDistanceStruct(t *testing.T) {
	var expected float64
	b := NewBrayCurtis()
	x := FloatsToMatrix([]float64{1, 2, 3})
	y := FloatsToMatrix([]float64{2, 4, 6})

	distance := b.Distance(x, y)
	expected = (1.0 / 3.0)
	if distance != expected {
		t.Log("Expected: ", expected, "but got ", distance)
		t.Fail()
	}
}
