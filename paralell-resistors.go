package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

type rpValue struct {
	numComponents int
	value         float64
}

func parallellResistors(c *cli.Context) error {
	eSeries := c.Int("component-series")
	//values := getValues(eSeries, defaultResistenceStart, defaultCapacitanceStop)
	values := getValues(eSeries, 1, 1000)
	count := len(values)
	fmt.Println(values)
	maxNum := c.Int("max-components")
	tolerance := c.Float64("tolerance") / 100.0
	desiredR, err := parseFloat64(c.String("desired-resistance"))
	if err != nil {
		return err
	}
	var foundValues []rpValue
	for i := 0; i < count; i++ {
		v := values[i]
		for n := 1; n <= maxNum; n++ {
			r := (1.0 / v) * float64(n)
			rp := 1.0 / r
			if equalByTolerance(desiredR, rp, tolerance) {
				foundValues = append(foundValues, rpValue{
					numComponents: n,
					value:         v,
				})
			}
		}
	}

	fmt.Println("foundValues", foundValues)
	return nil
}
