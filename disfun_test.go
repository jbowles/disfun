package disfun

import "testing"

func TestMinInt32(t *testing.T) {
	max := int(483264732)
	min := int(4387430)
	val := MinInt32(max, min)
	if val != min {
		t.Error("Expected MinInt to return: ", min, "but got: ", val)
	}
}

func TestMaxInt32(t *testing.T) {
	max := int(483264732)
	min := int(-456)
	val := MaxInt32(max, min)
	if val != max {
		t.Error("Expected MinInt to return: ", max, "but got: ", val)
	}
}

func TestMinInt64(t *testing.T) {
	max := float64(483264732)
	min := float64(4387430)
	val := MinFloat64(max, min)
	if val != min {
		t.Error("Expected MinInt to return: ", min, "but got: ", val)
	}
}

func TestMaxInt64(t *testing.T) {
	max := float64(483264732)
	min := float64(-456)
	val := MaxFloat64(max, min)
	if val != max {
		t.Error("Expected MinInt to return: ", max, "but got: ", val)
	}
}
