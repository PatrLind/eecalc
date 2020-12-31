package main

import (
	"math"
)

const defaultResistenceStart = 0.1
const defaultResistenceStop = 100000000000.0   // 100 Gohm
const defaultCapacitanceStart = 0.000000000001 // 1pF
const defaultCapacitanceStop = 10.0
const defaultInductanceStart = 0.000000000001 // 1pH
const defaultInductanceStop = 1000.0

// eSeries returns the E-series value where m is the E series group number
// and n is between 0 and m-1
// special cases for E24 are handled to give the industry standard values
func eSeries(m, n int) float64 {
	if m == 24 {
		switch n {
		case 10:
			return 2.7
		case 11:
			return 3.0
		case 12:
			return 3.3
		case 13:
			return 3.6
		case 14:
			return 3.9
		case 15:
			return 4.3
		case 16:
			return 4.7
		case 22:
			return 8.2
		}
	}
	v := root(math.Pow10(n), m)
	// round to one digit by default
	d := 10.0
	if m > 24 {
		// round to two digits
		d = 100.0
	}
	return math.Round(v*d) / d
}

func getValues(m int, start, stop float64) []float64 {
	orders := 0
	eSeriesList := make([]float64, m)
	for i := 0; i < m; i++ {
		eSeriesList[i] = eSeries(m, i)
	}
	for i := start; i < stop; i *= 10 {
		orders++
	}
	curVal := start
	count := m*orders + 1
	values := make([]float64, count)
	for i := 0; i < orders; i++ {
		orderStart := i * m
		for j := 0; j < m; j++ {
			curVal = math.Round(eSeriesList[j]*math.Pow10(i)*100.0) / 100.0
			values[orderStart+j] = curVal
		}
	}
	// also add the last value
	values[count-1] = math.Round(eSeriesList[0]*math.Pow10(orders+1)*100.0) / 100.0
	return values
}
