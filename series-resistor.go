package main

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/urfave/cli/v2"
)

func seriesResistor(c *cli.Context) error {
	vs, _, err := parseSI(c.String("supply-voltage"))
	if err != nil {
		return fmt.Errorf("unable to parse supply voltage: %w", err)
	}
	vd, _, err := parseSI(c.String("voltage-drop"))
	if err != nil {
		return fmt.Errorf("unable to parse voltage drop: %w", err)
	}
	i, _, err := parseSI(c.String("current"))
	if err != nil {
		return fmt.Errorf("unable to parse current: %w", err)
	}
	eSeries := c.Int("component-series")
	tolerance := getTolerance(c, eSeries)
	r := (vs - vd) / i
	p := (vs - vd) * i
	fmt.Printf("R = %s\n", humanize.SI(r, resistanceUnit))
	fmt.Printf("P = %s\n", humanize.SI(p, powerUnit))
	fmt.Printf("Suggested min power rating: %s to %s\n", humanize.SI(p*2, powerUnit), humanize.SI(p*10, powerUnit))
	fmt.Printf("Suggested component values (%.0f%%):\n", tolerance*100.0)
	printSuggestedValues(eSeries, tolerance, r, "  ", resistanceUnit)
	return nil
}
