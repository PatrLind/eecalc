package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	resistanceUnit  = "Ω"
	capacitanceUnit = "F"
	inductanceUnit  = "H"
	voltageUnit     = "V"
	powerUnit       = "W"
	currentUnit     = "A"
	timeUnit        = "s"
	energyUnit      = "J"
)

func main() {
	app := &cli.App{
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:        "rc-time-constant",
				Aliases:     []string{"rct"},
				Usage:       "Calculate RC time constant",
				Description: "The function will solve for a missing value (t, C or R)",
				Action:      rcTimeConstant,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "time",
						Aliases: []string{"time-constant"},
						Usage:   "time constant",
					},
					&cli.StringFlag{
						Name:    "capacitance",
						Aliases: []string{"c"},
						Usage:   "capacitance value",
					},
					&cli.StringFlag{
						Name:    "resistance",
						Aliases: []string{"r"},
						Usage:   "resistance value",
					},
					&cli.StringFlag{
						Name:    "voltage",
						Aliases: []string{"v", "volt"},
						Usage:   "charge voltage to calculate energy",
					},
					&cli.IntFlag{
						Name:    "component-series",
						Aliases: []string{"s"},
						Usage:   "E-series (example: 3, 6, 12, 24, 48, 96, 192) 12=10%, 24=5%, 96=1%",
						Value:   24,
					},
					&cli.Float64Flag{
						Name:    "tolerance",
						Aliases: []string{"t"},
						Usage:   "tolerance of the desired resistance (%)",
					},
				},
			},
			{
				Name:    "series-resistor",
				Aliases: []string{"sr"},
				Usage:   "Calculate series resistor (ex: for LED)",
				Action:  seriesResistor,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "supply-voltage",
						Aliases:  []string{"v"},
						Usage:    "voltage (ex: 12 V)",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "current",
						Aliases:  []string{"i"},
						Usage:    "desired current (ex 10 mA)",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "voltage-drop",
						Aliases:  []string{"d"},
						Usage:    "voltage drop over the component (ex: 2 V)",
						Required: true,
					},
					&cli.IntFlag{
						Name:    "component-series",
						Aliases: []string{"s"},
						Usage:   "E-series (example: 3, 6, 12, 24, 48, 96, 192) 12=10%, 24=5%, 96=1%",
						Value:   24,
					},
					&cli.Float64Flag{
						Name:    "tolerance",
						Aliases: []string{"t"},
						Usage:   "tolerance of the desired resistance (%)",
					},
				},
			},
			{
				Name:    "voltage-divider",
				Aliases: []string{"vd"},
				Usage:   "Calculate voltage devider values",
				Action:  voltageDivider,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "supply-voltage",
						Aliases: []string{"vs"},
						Usage:   "input voltage (ex: 12 V)",
					},
					&cli.StringFlag{
						Name:    "output-voltage",
						Aliases: []string{"vo"},
						Usage:   "output voltage (ex: 6 V)",
					},
					&cli.StringFlag{
						Name:  "r1",
						Usage: "first resistor of the voltage divider",
					},
					&cli.StringFlag{
						Name:  "r2",
						Usage: "second resistor of the voltage divider",
					},
					&cli.StringFlag{
						Name:  "rl",
						Usage: "resistance of the load (if applicable)",
					},
					&cli.IntFlag{
						Name:    "component-series",
						Aliases: []string{"s"},
						Usage:   "E-series (example: 3, 6, 12, 24, 48, 96, 192) 12=10%, 24=5%, 96=1%",
						Value:   24,
					},
					&cli.Float64Flag{
						Name:    "tolerance",
						Aliases: []string{"t"},
						Usage:   "tolerance of the desired resistance (%)",
					},
				},
			},
			{
				Name:    "equivalent",
				Aliases: []string{"eq"},
				Usage:   "calculate equivalent components",
				Subcommands: []*cli.Command{
					{
						Name:    "resistors",
						Aliases: []string{"resistance"},
						Action: func(c *cli.Context) error {
							return equivalentValues(c, false, resistanceUnit)
						},
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "desired-value",
								Aliases:  []string{"v"},
								Required: true,
							},
							&cli.Float64Flag{
								Name:    "max-components",
								Aliases: []string{"m"},
								Value:   10.0,
							},
							&cli.IntFlag{
								Name:    "component-series",
								Aliases: []string{"s"},
								Usage:   "E-series (example: 3, 6, 12, 24, 48, 96, 192) 12=10%, 24=5%, 96=1%",
								Value:   24,
							},
							&cli.Float64Flag{
								Name:    "tolerance",
								Aliases: []string{"t"},
								Usage:   "tolerance of the desired resistance (%)",
							},
							&cli.StringFlag{
								Name:    "power-handling",
								Aliases: []string{"p"},
								Usage:   "target power handling capability (W)",
							},
						},
					},
					{
						Name: "capacitors", Aliases: []string{"capacitance"},
						Action: func(c *cli.Context) error {
							return equivalentValues(c, true, capacitanceUnit)
						},
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "desired-value",
								Aliases:  []string{"v"},
								Required: true,
							},
							&cli.Float64Flag{
								Name:    "max-components",
								Aliases: []string{"m"},
								Value:   10.0,
							},
							&cli.IntFlag{
								Name:    "component-series",
								Aliases: []string{"s"},
								Usage:   "E-series (example: 3, 6, 12, 24, 48, 96, 192) 12=10%, 24=5%, 96=1%",
								Value:   24,
							},
							&cli.Float64Flag{
								Name:    "tolerance",
								Aliases: []string{"t"},
								Usage:   "tolerance of the desired resistance (%)",
							},
							&cli.StringFlag{
								Name:    "voltage-handling",
								Aliases: []string{"u"},
								Usage:   "target voltage handling capability (V)",
							},
						},
					},
					{
						Name: "inductors", Aliases: []string{"inductance"},
						Action: func(c *cli.Context) error {
							return equivalentValues(c, false, inductanceUnit)
						},
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "desired-value",
								Aliases:  []string{"v"},
								Required: true,
							},
							&cli.Float64Flag{
								Name:    "max-components",
								Aliases: []string{"m"},
								Value:   10.0,
							},
							&cli.IntFlag{
								Name:    "component-series",
								Aliases: []string{"s"},
								Usage:   "E-series (example: 3, 6, 12, 24, 48, 96, 192) 12=10%, 24=5%, 96=1%",
								Value:   24,
							},
							&cli.Float64Flag{
								Name:    "tolerance",
								Aliases: []string{"t"},
								Usage:   "tolerance of the desired resistance (%)",
							},
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
