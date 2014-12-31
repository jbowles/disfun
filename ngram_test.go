package disfun

import "testing"

func TestNgramDistanceZero(t *testing.T) {
	a := "Hallo und guten Morgen"
	b := "Hallo und guten Morgen"
	n := NewNgram(3, a, b)
	similar := n.Similarity()
	expected := 1.0
	if similar != expected {
		t.Log("Expected: ", expected, "but got ", similar)
		t.Fail()
	}
}

func TestNgramDistance(t *testing.T) {
	a := "Hallo und guten Morgen"
	b := "Hallo und guten Tag"
	n := NewNgram(3, a, b)
	similar := n.Similarity()
	expected := 0.6086956521739131
	if similar != expected {
		t.Log("Expected: ", expected, "but got ", similar)
		t.Fail()
	}
}
