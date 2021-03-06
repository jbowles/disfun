package disfun

import (
	"github.com/gonum/matrix/mat64"
)

//Levenshtein (edit distance) gives similarity metric by calcuating number of positions for substitution, insertion, and deletion.
//
//
//Currently bemchmarking 4 different implementations. One is a simple pass through (Lev), another (Leven) uses a struct and separate insertion, deletion, substitution functions with a handmade matrix called VectorCell. The third one (Levenshtein) uses the mat64 matrix package. The last one (EditDistance) is the fastest but I haven't checked it accuracy yet.
//
//	Row     == Height
//	Column  == Width
//
//	Demo here: http://andrew.hedges.name/experiments/levenshtein/
//
// Levenshtein is the most accurate but also the most costly due to building matrices. I can get this down but using dense matrices for this is not a good idea. The other two functions are about the same speed but they are not very readable and not the most accurate... see the tests for differences in accuracy between Leven, Lev, and Levenshtein.
type Levenshtein struct {
	Source       []rune
	Target       []rune
	SourceString string
	TargeString  string
	RowHeight    int
	ColWidth     int
	M            *mat64.Dense
}

func NewLevenshtein(source, target string) *Levenshtein {
	rsource := []rune(source)
	rtarget := []rune(target)
	rows := len(rsource) + 1
	columns := len(rtarget) + 1
	return &Levenshtein{
		Source:       rsource,
		Target:       rtarget,
		SourceString: source,
		TargeString:  target,
		RowHeight:    rows,
		ColWidth:     columns,
		M:            zeroDense(rows, columns),
	}
}

func (l *Levenshtein) ComputeMatrix() *mat64.Dense {

	// cells for i of source string
	for i := 0; i < l.RowHeight; i++ {
		l.M.Set(i, 0, float64(i))
	}

	// cell for j of target string
	for j := 1; j < l.ColWidth; j++ {
		l.M.Set(0, j, float64(j))
	}

	for i := 1; i < l.RowHeight; i++ {
		for j := 1; j < l.ColWidth; j++ {
			if l.Source[i-1] == l.Target[j-1] {
				l.M.Set(i, j, (l.M.At(i-1, j-1)))
			} else {
				delCost := l.M.At(i-1, j) + Deletion
				subCost := l.M.At(i-1, j-1) + Substitution
				insCost := l.M.At(i, j-1) + Insertion
				l.M.Set(i, j, minFloat64(delCost, subCost, insCost))
			}
		}
	}
	return l.M
}

func (l *Levenshtein) Similarity() float64 {
	l.ComputeMatrix()
	return l.M.At(len(l.Source), len(l.Target))
}

//Leven (edit distance) gives similarity metric by calcuating number of positions for substitution, insertion, and deletion.
//
//Currently bemchmarking 4 different implementations. One is a simple pass through (Lev), another (Leven) uses a struct and separate insertion, deletion, substitution functions with a handmade matrix called VectorCell. The third one (Levenshtein) uses the mat64 matrix package. The last one (EditDistance) is the fastest but I haven't checked it accuracy yet.
//
// Levenshtein is the most accurate but also the most costly due to building matrices. I can get this down but using dense matrices for this is not a good idea. The other two functions are about the same speed but they are not very readable and not the most accurate... see the tests for differences in accuracy between Leven, Lev, and Levenshtein.
type Leven struct {
	S1         string
	S2         string
	Lens1      int
	Lens2      int
	Width      int
	VectorCell []int
}

func NewLeven(s1, s2 string) *Leven {
	lens1 := len(s1)
	lens2 := len(s2)
	return &Leven{
		S1:         s1,
		S2:         s2,
		Lens1:      lens1,
		Lens2:      lens2,
		Width:      (lens2 - 1),
		VectorCell: make([]int, lens1*lens2),
	}
}

func (l *Leven) deletion(idx1, idx2 int) int {
	return l.VectorCell[(idx1-1)*l.Width+idx2] + 1
}
func (l *Leven) insertion(idx1, idx2 int) int {
	return l.VectorCell[(idx1*l.Width+(idx2-1))] + 1
}
func (l *Leven) substitution(idx1, idx2 int) int {
	return l.VectorCell[((idx1-1)*l.Width+(idx2-1))] + 1
}

func (l *Leven) Similarity() float64 {

	// cells for i of s1
	for idxS1 := 1; idxS1 < l.Lens1; idxS1++ {
		l.VectorCell[idxS1*l.Width] = idxS1
	}

	// cell for j of n(s2)
	for idxS2 := 1; idxS2 < l.Lens2; idxS2++ {
		l.VectorCell[l.Width+idxS2] = idxS2
	}

	for idxS2 := 1; idxS2 < l.Lens2; idxS2++ {
		for idxS1 := 1; idxS1 < l.Lens1; idxS1++ {
			if l.S1[idxS1] == l.S2[idxS2] {
				l.VectorCell[idxS1*l.Width+idxS2] = l.VectorCell[(idxS1-1)*l.Width+(idxS2-1)]
			} else {
				l.VectorCell[idxS1*l.Width+idxS2] = minInt32(
					l.deletion(idxS1, idxS2),
					l.insertion(idxS1, idxS2),
					l.substitution(idxS1, idxS2))
			}
		}
	}
	return float64(l.VectorCell[l.Lens1*l.Width])
}

//Lev (edit distance) gives similarity metric by calcuating number of positions for substitution, insertion, and deletion.
//
//Currently bemchmarking 4 different implementations. One is a simple pass through (Lev), another (Leven) uses a struct and separate insertion, deletion, substitution functions with a handmade matrix called VectorCell. The third one (Levenshtein) uses the mat64 matrix package. The last one (EditDistance) is the fastest but I haven't checked it accuracy yet.
//
// Levenshtein is the most accurate but also the most costly due to building matrices. I can get this down but using dense matrices for this is not a good idea. The other two functions are about the same speed but they are not very readable and not the most accurate... see the tests for differences in accuracy between Leven, Lev, and Levenshtein.
func Lev(s1, s2 string) int {
	m1 := len(s1)
	n2 := len(s2)
	width := n2 - 1
	vcell := make([]int, m1*n2)

	// cells for i of m(s1)
	for i1 := 1; i1 < m1; i1++ {
		vcell[i1*width+0] = i1
	}

	// cell for j of n(s2)
	for j2 := 1; j2 < n2; j2++ {
		vcell[0*width+j2] = j2
	}

	for j2 := 1; j2 < n2; j2++ {
		for i1 := 1; i1 < m1; i1++ {
			if s1[i1] == s2[j2] {
				vcell[i1*width+j2] = vcell[(i1-1)*width+(j2-1)]
			} else {
				deletion := vcell[(i1-1)*width+j2] + 1
				insertion := vcell[(i1*width+(j2-1))] + 1
				substitution := vcell[((i1-1)*width+(j2-1))] + 1
				vcell[i1*width+j2] = minInt32(deletion, insertion, substitution)
			}
		}
	}
	return vcell[m1*width]
}

// LevEditDistance computes the Levenshtein distance between two strings. The returned value - distance - is the number of insertions, deletions, and substitutions it takes to transform one string (s1) into another (s2). Each step in the transformation "costs" one distance point.
//
//
// //Currently bemchmarking 4 different implementations. One is a simple pass through (Lev), another (Leven) uses a struct and separate insertion, deletion, substitution functions with a handmade matrix called VectorCell. The third one (Levenshtein) uses the mat64 matrix package. The last one (EditDistance) is the fastest but I haven't checked it accuracy yet.
func LevEditDistance(s1, s2 string) (distance int) {
	// index by code point, not byte
	r1 := []rune(s1)
	r2 := []rune(s2)

	rows := len(r1) + 1
	cols := len(r2) + 1

	var d1 int
	var d2 int
	var d3 int
	var i int
	var j int
	dist := make([]int, rows*cols)

	for i = 0; i < rows; i++ {
		dist[i*cols] = i
	}

	for j = 0; j < cols; j++ {
		dist[j] = j
	}

	for j = 1; j < cols; j++ {
		for i = 1; i < rows; i++ {
			if r1[i-1] == r2[j-1] {
				dist[(i*cols)+j] = dist[((i-1)*cols)+(j-1)]
			} else {
				d1 = dist[((i-1)*cols)+j] + 1
				d2 = dist[(i*cols)+(j-1)] + 1
				d3 = dist[((i-1)*cols)+(j-1)] + 1

				dist[(i*cols)+j] = minInt32(d1, minInt32(d2, d3))
			}
		}
	}

	distance = dist[(cols*rows)-1]

	return
}
