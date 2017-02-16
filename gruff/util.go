package gruff

import (
	"math"
)

func Round(f float64) float64 {
	return math.Floor(f + .5)
}

func RoundToDecimal(f float64, decimals int) float64 {
	factor := math.Pow(10, float64(decimals))
	res := f * factor
	res = Round(res)
	res = res / factor
	return res
}
