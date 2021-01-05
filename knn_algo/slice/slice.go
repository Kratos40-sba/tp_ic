package slice

import "sort"

type Slice struct {
	sort.Interface
	Idx []int
}

func (s Slice) Swap(i, j int) {
	s.Interface.Swap(i, j)
	s.Idx[i], s.Idx[j] = s.Idx[j], s.Idx[i]
}
func NewSlice(n sort.Interface) *Slice {
	s := &Slice{n, make([]int, n.Len())}
	for i := range s.Idx {
		s.Idx[i] = i
	}
	return s
}
func NewFlag64Slice(n []float64) *Slice {
	return NewSlice(sort.Float64Slice(n))
}
