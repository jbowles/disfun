package disfun

import (
	"github.com/jbowles/disfun/Godeps/_workspace/src/github.com/juju/utils/set"
)

//Ngram is continuous sequence of n-items from a given sequence. The distance is the relative number of items between these two sequences.
//
// References:
//
//	https://webdocs.cs.ualberta.ca/~kondrak/papers/spire05.pdf
//	https://en.wikipedia.org/wiki/N-gram
//	http://m.wolframalpha.com/input/?i=n-grams+%22n-gram+example+of+n-grams+in+wolfram+alpha%22&x=0&y=0
type Ngram struct {
	Set1 set.Strings
	Set2 set.Strings
	S1   string
	S2   string
	N    int
}

// NewNgram initializes and creates data structures for the Ngram struct, which you can then call Similarity() on.
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

// JaccardCoEfficient calculates the similarity of two sets as the intersection divided by the union of the two sets.
func (n *Ngram) JaccardCoEfficient() float64 {
	return float64(n.Set1.Intersection(n.Set2).Size()) / float64(n.Set1.Union(n.Set2).Size())
}

// Build creates Set1 and Set2 of the Ngram.
func (n *Ngram) Build() {
	for i := 0; i < (len(n.S1) - (n.N) + 1); i++ {
		n.Set1.Add(n.S1[i : i+n.N])
	}
	for i := 0; i < (len(n.S2) - (n.N) + 1); i++ {
		n.Set2.Add(n.S2[i : i+n.N])
	}
}

// Similarity bulds the two ngram sets and returns their similarity.
func (n *Ngram) Similarity() float64 {
	n.Build()
	return n.JaccardCoEfficient()
}
