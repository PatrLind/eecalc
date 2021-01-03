package main

import (
	"fmt"
	"math"

	"github.com/dustin/go-humanize"
	"github.com/urfave/cli/v2"
)

func voltageDivider(c *cli.Context) error {
	vs, _, err := parseSI(c.String("supply-voltage"))
	if err != nil {
		return fmt.Errorf("unable to parse supply voltage: %w", err)
	}
	vo, _, err := parseSI(c.String("output-voltage"))
	if err != nil {
		return fmt.Errorf("unable to parse output voltage: %w", err)
	}
	r1, _, err := parseSI(c.String("r1"))
	if err != nil {
		return fmt.Errorf("unable to parse r1: %w", err)
	}
	r2, _, err := parseSI(c.String("r2"))
	if err != nil {
		return fmt.Errorf("unable to parse r2: %w", err)
	}
	rLoad, _, err := parseSI(c.String("rl"))
	if err != nil {
		return fmt.Errorf("unable to parse rl: %w", err)
	}
	eSeries := c.Int("component-series")
	tolerance := getTolerance(c, eSeries)
	if ((vs == 0 && (r1 == 0 || r2 == 0 || vo == 0)) ||
		(vo == 0 && (vs == 0 || r1 == 0 || r2 == 0)) ||
		(r1 == 0 && (vs == 0 || r2 == 0 || vo == 0)) ||
		(r2 == 0 && (r1 == 0 || vs == 0 || vo == 0))) && !(r1 == 0 && r2 == 0 && rLoad != 0) {
		return fmt.Errorf("you need to give at least three values to work with (s, o, r1 or r2)")
	}
	r2Orig := r2
	if rLoad != 0 {
		if r1 == 0 && r2 == 0 {
			r1 = (rLoad * (vs - vo)) / vo
			fmt.Printf("R1: %s\n", humanize.SI(r1, resistanceUnit))
			fmt.Printf("R2: not used (R2 is RL)\n")
			r2 = rLoad
			r2Orig = math.Inf(1)
		} else {
			r2 = 1 / ((1 / r2) + (1 / rLoad))
		}
	}
	if vs == 0 {
		vs = (vo * r2) / (r1 + r2)
		fmt.Printf("Supply voltage: %s\n", humanize.SI(vs, voltageUnit))
	} else if vo == 0 {
		vo = (vs * r2) / (r1 + r2)
		fmt.Printf("Output voltage: %s\n", humanize.SI(vo, voltageUnit))
	} else if r1 == 0 {
		r1 = (r2 * (vs - vo)) / vo
		fmt.Printf("R1: %s\n", humanize.SI(r1, resistanceUnit))
	} else if r2 == 0 {
		r2 = (vo * r1) / (vs - vo)
		if r2Orig == 0 {
			r2Orig = r2
		}
		fmt.Printf("R2: %s\n", humanize.SI(r2, resistanceUnit))
	}
	is := vs / (r1 + r2)
	pTot := vs * is
	fmt.Printf("Power supply power delivery: %s (%s)\n", humanize.SI(pTot, powerUnit), humanize.SI(is, currentUnit))
	if rLoad != 0 {
		i2 := vo / r2Orig
		iLoad := vo / rLoad
		pRL := vo * iLoad
		pR2 := vo * i2
		fmt.Printf("Load resistor power: %s (%s)\n", humanize.SI(pRL, powerUnit), humanize.SI(iLoad, currentUnit))
		fmt.Printf("Load resistor %% of total power: %.01f%%\n", pRL/pTot*100)
		fmt.Printf("R2 power: %s (%s)\n", humanize.SI(pR2, powerUnit), humanize.SI(i2, currentUnit))
	} else {
		i2 := vo / r2Orig
		pR2 := vo * i2
		fmt.Printf("R2 resistor %% of total power: %.01f%%\n", pR2/pTot*100)
		fmt.Printf("R2 power: %s (%s)\n", humanize.SI(pR2, powerUnit), humanize.SI(i2, currentUnit))
	}
	printSuggestedValues(eSeries, tolerance, r1, "R1", resistanceUnit)
	if r2Orig != math.Inf(1) {
		printSuggestedValues(eSeries, tolerance, r2, "R2", resistanceUnit)
	}

	return nil
}
