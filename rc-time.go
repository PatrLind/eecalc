package main

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/urfave/cli/v2"
)

func rcTimeConstant(c *cli.Context) error {
	time, _, err := parseSI(c.String("time"))
	if err != nil {
		return fmt.Errorf("unable to parse time value: %w", err)
	}
	capacitance, _, err := parseSI(c.String("capacitance"))
	if err != nil {
		return fmt.Errorf("unable to parse time value: %w", err)
	}
	resistance, _, err := parseSI(c.String("resistance"))
	if err != nil {
		return fmt.Errorf("unable to parse time value: %w", err)
	}
	voltage, _, err := parseSI(c.String("voltage"))
	if err != nil {
		return fmt.Errorf("unable to parse voltage value: %w", err)
	}
	eSeries := c.Int("component-series")
	tolerance := getTolerance(c, eSeries)
	if (time == 0 && (capacitance == 0 || resistance == 0)) ||
		(capacitance == 0 && (time == 0 || resistance == 0)) ||
		(resistance == 0 && (capacitance == 0 || time == 0)) {
		return fmt.Errorf("you need to give at least two values to work with (t, C or R)")
	}
	fmt.Println("τ = time constant = approximately 63% of the charge time")
	if time == 0 {
		time = capacitance * resistance
		fmt.Printf("τ = CR = %s\n", humanize.SI(time, timeUnit))
	} else if resistance == 0 {
		resistance = time / capacitance
		fmt.Printf("R = τ/C = %s\n", humanize.SI(resistance, resistanceUnit))
		printSuggestedValues(eSeries, tolerance, resistance, "R", resistanceUnit)
	} else {
		capacitance = time / resistance
		fmt.Printf("C = τ/R = %s\n", humanize.SI(capacitance, capacitanceUnit))
	}
	if voltage != 0 {
		energy := voltage * voltage * (capacitance / 2)
		fmt.Printf("E = V²×R/2 = %s\n", humanize.SI(energy, energyUnit))
	}

	fmt.Printf("Time to fully charged capacitor = 5τ = %s\n", humanize.SI(time*5, "s"))
	return nil
}
