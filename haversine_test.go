package disfun

import "testing"

func TestHaversinePointZeroDistance(t *testing.T) {
	pa1 := 203.9
	pa2 := 203.9
	pb1 := 203.9
	pb2 := 203.9
	radius := 8098.8
	distance := HaversinePoint(pa1, pa2, pb1, pb2, radius)
	expected := 0.0

	if distance != expected {
		t.Log("Expected: ", expected, "but got ", distance)
		t.Fail()
	}
}

func TestHaversinePointDistanceRadiusZero(t *testing.T) {
	pa1 := 78543.03
	pa2 := 00.143567
	pb1 := 233057.87
	pb2 := 00.18756783
	radius := 0.0
	distance := HaversinePoint(pa1, pa2, pb1, pb2, radius)
	expected := 0.0

	if distance != expected {
		t.Log("Expected: ", expected, "but got ", distance)
		t.Fail()
	}
}

func TestHaversinePointDistanceRadiusNegOne(t *testing.T) {
	pa1 := 78543.03
	pa2 := 00.143567
	pb1 := 233057.87
	pb2 := 00.18756783
	radius := -1.0
	distance := HaversinePoint(pa1, pa2, pb1, pb2, radius)
	expected := -1.3062043094350282

	if distance != expected {
		t.Log("Expected: ", expected, "but got ", distance)
		t.Fail()
	}
}

func TestHaversinePointDistanceRadiusOne(t *testing.T) {
	pa1 := 78543.03
	pa2 := 00.143567
	pb1 := 233057.87
	pb2 := 00.18756783
	radius := 2.0
	distance := HaversinePoint(pa1, pa2, pb1, pb2, radius)
	expected := 2.6124086188700564

	if distance != expected {
		t.Log("Expected: ", expected, "but got ", distance)
		t.Fail()
	}
}

func TestHaversinePointDistance(t *testing.T) {
	pa1 := 40.3
	pa2 := 112.0
	pb1 := 82.4
	pb2 := 67.19
	radius := 1345.09
	distance := HaversinePoint(pa1, pa2, pb1, pb2, radius)
	expected := 1045.8094451034126

	if distance != expected {
		t.Log("Expected: ", expected, "but got ", distance)
		t.Fail()
	}
}

func TestHaversinePointDistEagleToShastaMountain(t *testing.T) {
	pa1 := 40.3061
	pa2 := 112.0097
	pb1 := 41.4092
	pb2 := 122.1949
	distance := HaversinePoint(pa1, pa2, pb1, pb2, EarthRadius)
	expected := 864.8005613102791

	if distance != expected {
		t.Log("Expected: ", expected, "but got ", distance)
		t.Fail()
	}
}

func TestHaversineLatLonZeroDistance(t *testing.T) {
	lat1 := 40.3061
	lon1 := 112.0097
	lat2 := 40.3061
	lon2 := 112.0097
	distance := HaversineLatLon(lat1, lon1, lat2, lon2)
	expected := 0.0

	if distance != expected {
		t.Log("Expected: ", expected, "but got ", distance)
		t.Fail()
	}
}

func TestHaversineLatLonDistEagleToShastaMountain(t *testing.T) {
	lat1 := 40.3061
	lon1 := 112.0097
	lat2 := 41.4092
	lon2 := 122.1949
	distance := HaversineLatLon(lat1, lon1, lat2, lon2)
	expected := 864.8005613102791

	if distance != expected {
		t.Log("Expected: ", expected, "but got ", distance)
		t.Fail()
	}
}
