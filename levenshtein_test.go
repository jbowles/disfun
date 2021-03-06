package disfun

/*
RUN BENCHMARKS: go test -bench=.
*/

import (
	"testing"
)

var s1 = "supercalibro"
var s2 = "suprcalbro"
var s3 = "levenshtein"
var s4 = "hamming"
var s5 = "jfkldhfjdkhsayruiewfjiodnmcopjdiohvidhnsuvjhdioshvbudsjkl;ieiryewfbcklsjmcmdkvhewioufsauifgbeuncoejkwgfoijewjfdklajfkheklwhfkhdjskhfjkldshfjdsajfkldhsfk"
var s6 = "fkdhsavhdjslfikjfdkshfjjdslkjfdsjfkldshfewijinewfnioenvd.hsavhnjdnvkdjsvjndvndksnvkdsjfovdsjfopwqjfihgjhnvioewhfkjdskfhdjsjflds;jfljdkshfidlsaofjdksfiods"

func BenchmarkLevenshteinDistance(b *testing.B) {
	for n := 0; n < b.N; n++ {
		l0 := NewLevenshtein(s5, s6)
		//l1 := NewLevenshtein(s3, s4)
		l0.Similarity()
		//l1.Similarity()
	}
}

func BenchmarkLevenDistance(b *testing.B) {
	for n := 0; n < b.N; n++ {
		l0 := NewLeven(s5, s6)
		//l1 := NewLeven(s3, s4)
		l0.Similarity()
		//l1.Similarity()
	}
}

func BenchmarkLevDistance(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Lev(s5, s6)
		//Lev(s3, s4)
	}
}

func BenchmarkLevEditDistance(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LevEditDistance(s5, s6)
		//Lev(s3, s4)
	}
}

func TestLevenshteinDistanceTwo(t *testing.T) {
	l := NewLevenshtein(s1, s2)
	distance := l.Similarity()
	expected := 2.0

	if distance != expected {
		t.Log("Expected: ", expected, "but got ", distance)
		t.Fail()
	}
}

func TestLevenshteinDistanceMucho(t *testing.T) {
	l := NewLevenshtein(s3, s4)
	distance := l.Similarity()
	expected := 10.0

	if distance != expected {
		t.Log("Expected: ", expected, "but got ", distance)
		t.Fail()
	}
}

func TestLevenDistanceTwo(t *testing.T) {
	l1 := NewLeven(s1, s2)
	distance := l1.Similarity()
	expected := 2.0

	if distance != expected {
		t.Log("Expected: ", expected, "but got ", distance)
		t.Fail()
	}
}

func TestLevenDistanceMucho(t *testing.T) {
	l1 := NewLeven(s3, s4)
	distance := l1.Similarity()
	expected := 9.0

	if distance != expected {
		t.Log("Expected: ", expected, "but got ", distance)
		t.Fail()
	}
}

func TestLevDistanceTwo(t *testing.T) {
	distance := Lev(s1, s2)
	expected := 2

	if distance != expected {
		t.Log("Expected: ", expected, "but got ", distance)
		t.Fail()
	}
}

func TestLevDistanceMucho(t *testing.T) {
	distance := Lev(s3, s4)
	expected := 9

	if distance != expected {
		t.Log("Expected: ", expected, "but got ", distance)
		t.Fail()
	}
}

var levtests = []struct {
	s1   string
	s2   string
	dist int
}{
	// insertion
	{"car", "cars", 1},
	// substitution
	{"library", "librari", 1},
	// deletion
	{"library", "librar", 1},
	// one empty, left
	{"", "library", 7},
	// one empty, right
	{"library", "", 7},
	// two empties
	{"", "", 0},
	// unicode stuff!
	{"Schüßler", "Schübler", 1},
	{"Schüßler", "Schußler", 1},
	{"Schüßler", "Schüßler", 0},
	{"Schüßler", "Schüler", 1},
	{"Schüßler", "Schüßlers", 1},
}

// Regular Levenshtein
func TestLevEditDistance(t *testing.T) {
	for _, tt := range levtests {
		dist := LevEditDistance(tt.s1, tt.s2)
		if dist != tt.dist {
			t.Errorf("Levenshtein('%s', '%s') = %v, want %v", tt.s1, tt.s2, dist, tt.dist)
		}
	}
}
