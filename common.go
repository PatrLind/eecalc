package main

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/urfave/cli/v2"
)

func getTolerance(c *cli.Context, eSeries int) float64 {
	t := c.Float64("tolerance") / 100.0
	if t == 0 {
		t = toleranceFromESeries(eSeries)
	}
	return t
}

func printSuggestedValues(eSeries int, tolerance, value float64, name, unit string) {
	fmt.Printf("Suggested component values for %s (%.0f%%):\n", name, tolerance*100.0)
	vs := getSuggestedComponentValues(eSeries, tolerance, value)
	if len(vs) == 0 {
		fmt.Println("  no suggestions within tolerance")
	}
	for _, v := range vs {
		fmt.Printf("  %s\n", humanize.SI(v, unit))
	}
}

func getSuggestedComponentValues(eSeries int, tolerance, value float64) []float64 {
	start := findNearestBase10(value)
	stop := findNearestBase10(value * 10)
	values := getValues(eSeries, start, stop)
	var ret []float64
	for _, v := range values {
		if equalByTolerance(value, v, tolerance) {
			ret = append(ret, v)
		}
	}
	return ret
}
