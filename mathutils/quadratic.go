package mathutils

import "math"

func SolveQuadratic(a, b, c float64) (roots []float64, hasSolution bool) {
	D := b*b - 4*a*c
	if D < 0 {
		return nil, false
	}
	if D == 0 {
		return []float64{-b / (2 * a)}, true
	}
	sqrtD := math.Sqrt(D)
	return []float64{
		(-b + sqrtD) / (2 * a),
		(-b - sqrtD) / (2 * a),
	}, true
}
