package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/dustin/go-humanize"
	"github.com/urfave/cli/v2"
)

type valueT struct {
	numP       int
	numS       int
	value      float64
	totalValue float64
}

func equivalentValues(c *cli.Context, invert bool, unit string) error {
	// Collect user input
	eSeries := c.Int("component-series")
	maxNum := c.Int("max-components")
	tolerance := c.Float64("tolerance") / 100.0
	desiredValue, _, err := humanize.ParseSI(c.String("desired-value"))
	if err != nil {
		return err
	}

	// Find suitable values based on the desired value
	var start, stop float64
	start = findNearestBase10(desiredValue / float64(maxNum))
	stop = findNearestBase10(desiredValue * float64(maxNum))
	values := getValues(eSeries, start, stop)
	count := len(values)

	// Find value combinations
	var foundValues []valueT
	for i := 0; i < count; i++ {
		v := values[i]
		for n1 := 1; n1 <= maxNum; n1++ {
			var rp float64
			if invert {
				rp = v * float64(n1)
			} else {
				rp = (1.0 / v) * float64(n1)
			}
			if !invert {
				rp = 1.0 / rp
			}
			for n2 := 1; n2 <= maxNum && n2*n1 <= maxNum; n2++ {
				var rs float64
				if invert {
					rs = (1.0 / rp) * float64(n2)
				} else {
					rs = rp * float64(n2)
				}
				if invert {
					rs = 1.0 / rs
				}
				if equalByTolerance(desiredValue, rs, tolerance) {
					foundValues = append(foundValues, valueT{
						numP:       n1,
						numS:       n2,
						value:      v,
						totalValue: rs,
					})
				}
			}
		}
	}

	// Sort the found values
	sort.SliceStable(foundValues, func(i, j int) bool {
		v1 := foundValues[i]
		v2 := foundValues[j]
		// Sort by the least number of components first
		if v1.numP+v1.numS < v2.numP+v2.numS {
			return true
		}
		// Sort by the best value second
		d1 := math.Abs(desiredValue - v1.totalValue)
		d2 := math.Abs(desiredValue - v2.totalValue)
		return d1 < d2
	})

	// Print found values
	for _, v := range foundValues {
		vStr := humanize.SI(v.value, unit)
		vTotStr := humanize.SI(v.totalValue, unit)
		if v.numP == 1 && v.numS == 1 {
			continue
		} else if v.numP == 1 {
			fmt.Println(fmt.Sprintf("%dx %s in series = %s", v.numS, vStr, vTotStr))
		} else if v.numS == 1 {
			fmt.Println(fmt.Sprintf("%dx %s in parallel = %s", v.numP, vStr, vTotStr))
		} else {
			fmt.Println(fmt.Sprintf("%dx (%dx %s in parallel) in series = %s", v.numS, v.numP, vStr, vTotStr))
		}
	}
	return nil
}
