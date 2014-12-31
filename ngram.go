package disfun

import (
	"github.com/juju/utils/set"
)

type Ngram struct {
	Set1 set.Strings
	Set2 set.Strings
	S1   string
	S2   string
	N    int
}

func NewNgram(n int, s1, s2 string) *Ngram {
	set1 := set.NewStrings()
	set2 := set.NewStrings()
	return &Ngram{
		Set1: set1,
		Set2: set2,
		S1:   s1,
		S2:   s2,
		N:    n,
	}
}

func (n *Ngram) JaccardCoEfficient() float64 {
	return float64(n.Set1.Intersection(n.Set2).Size()) / float64(n.Set1.Union(n.Set2).Size())
}

func (n *Ngram) Build() {
	for i := 0; i < (len(n.S1) - (n.N) + 1); i++ {
		n.Set1.Add(n.S1[i : i+n.N])
	}
	for i := 0; i < (len(n.S2) - (n.N) + 1); i++ {
		n.Set2.Add(n.S2[i : i+n.N])
	}
}

func (n *Ngram) Similarity() float64 {
	n.Build()
	return n.JaccardCoEfficient()
}
