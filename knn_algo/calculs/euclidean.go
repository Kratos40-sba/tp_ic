package calculs

import "math"

func EuclideanDist(source, dist []float64) float64 {
	v := 0.0
	for i := range dist {
		v += math.Pow(source[i]-dist[i], 2)
	}
	return math.Sqrt(v)
}
