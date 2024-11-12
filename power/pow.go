package power

import "math"

func GetPower(a *float64, b float64) float64 {
	return math.Pow(*a, b)
}
