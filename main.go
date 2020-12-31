package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	resistanceUnit  = "Ω"
	capacitanceUnit = "F"
	inductanceUnit  = "H"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "equivalent",
				Usage: "calculate equivalent components",
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
								Value:   1,
							},
						},
					},
					{
						Name:    "capacitors",
						Aliases: []string{"capacitance"},
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
								Value:   1,
							},
						},
					},
					{
						Name:    "inductors",
						Aliases: []string{"inductance"},
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
								Value:   1,
							},
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
