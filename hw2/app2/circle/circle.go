package circle

import "math"

func CircleParams(area float64) (float64, float64) {
	diameter := math.Sqrt(area/math.Pi) * 2
	circleLen := math.Pi * diameter
	return diameter, circleLen
}
