package knn

import (
	"github.com/Kratos40-sba/tp_ic/knn_algo/calculs"
	"github.com/Kratos40-sba/tp_ic/knn_algo/slice"
	"github.com/Kratos40-sba/tp_ic/knn_algo/sortmap"
	"sort"
)

type Knn struct {
	K      int
	Res    []string
	Donnee [][]float64
}

func (knn *Knn) Fit(x [][]float64, y []string) {
	knn.Donnee = x
	knn.Res = y
}
func (knn *Knn) Predict(x [][]float64) []string {
	var res []string
	for _, src := range x {
		var (
			voisins      []string
			distanceList []float64
		)
		for _, dist := range knn.Donnee {
			distanceList = append(distanceList, calculs.EuclideanDist(src, dist))
		}
		s := slice.NewFlag64Slice(distanceList)
		sort.Sort(s)
		index := s.Idx[:knn.K]
		for _, i := range index {
			voisins = append(voisins, knn.Res[i])
		}
		freqVoisinage := calculs.Counter(voisins)
		a := sortmap.List{}
		for key, value := range freqVoisinage {
			e := sortmap.Map{Name: key, Value: value}
			a = append(a, e)
		}
		sort.Sort(a)
		res = append(res, a[0].Name)

	}
	return res

}
