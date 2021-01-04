# eecalc

Simple calculator for electrical engineering things

I am writing this application for my own purposes. I add functionality as I want to experiment with new things related to electrical engineering calculations.
If anyone finds it useful please go ahead and use it, but I cannot guarantee that it works correctly.

The application will give component value suggestions based on a selectable E-series and component tolerance where appropriate.

## Functions

### rc-time-constant - Calculate RC time constant

```text
NAME:
   eecalc.exe rc-time-constant - Calculate RC time constant

USAGE:
   eecalc.exe rc-time-constant [command options] [arguments...]

DESCRIPTION:
   The function will solve for a missing value (t, C or R)

OPTIONS:
   --time value, --time-constant value      time constant
   --capacitance value, -c value            capacitance value
   --resistance value, -r value             resistance value
   --voltage value, -v value, --volt value  charge voltage to calculate energy
   --component-series value, -s value       E-series (example: 3, 6, 12, 24, 48, 96, 192) 12=10%, 24=5%, 96=1% (default: 24)
   --tolerance value, -t value              tolerance of the desired resistance (%) (default: 0)
   --help, -h                               show help (default: false)
```

Examples:

```text
.\eecalc.exe rc-time-constant -c 100u -r 4.7k
τ = time constant = approximately 63% of the charge time
τ = CR = 470 ms
Time to fully charged capacitor = 5τ = 2.35 s
```

```text
.\eecalc.exe rc-time-constant -c 100u --time 1.12s
τ = time constant = approximately 63% of the charge time
R = τ/C = 11.2 kΩ
Suggested component values for R (5%):
  11 kΩ
Time to fully charged capacitor = 5τ = 5.6 s
```

### series-resistor - Calculate series resistor (ex: for LED)

```text
NAME:
   eecalc.exe series-resistor - Calculate series resistor (ex: for LED)

USAGE:
   eecalc.exe series-resistor [command options] [arguments...]

OPTIONS:
   --supply-voltage value, -v value    voltage (ex: 12 V)
   --current value, -i value           desired current (ex 10 mA)
   --voltage-drop value, -d value      voltage drop over the component (ex: 2 V)
   --component-series value, -s value  E-series (example: 3, 6, 12, 24, 48, 96, 192) 12=10%, 24=5%, 96=1% (default: 24)
   --tolerance value, -t value         tolerance of the desired resistance (%) (default: 0)
   --help, -h                          show help (default: false)
```

Example:

```text
.\eecalc.exe series-resistor -v 12 -i 5m -d 2.2
R = 1.96 kΩ
P = 49 mW
Suggested min power rating: 98 mW to 490 mW
Suggested component values (5%):
  2 kΩ
```

### voltage-divider - Calculate voltage devider values

```text
NAME:
   eecalc.exe voltage-divider - Calculate voltage devider values

USAGE:
   eecalc.exe voltage-divider [command options] [arguments...]

OPTIONS:
   --supply-voltage value, --vs value  input voltage (ex: 12 V)
   --output-voltage value, --vo value  output voltage (ex: 6 V)
   --r1 value                          first resistor of the voltage divider
   --r2 value                          second resistor of the voltage divider
   --rl value                          resistance of the load (if applicable)
   --component-series value, -s value  E-series (example: 3, 6, 12, 24, 48, 96, 192) 12=10%, 24=5%, 96=1% (default: 24)
   --tolerance value, -t value         tolerance of the desired resistance (%) (default: 0)
   --help, -h                          show help (default: false)
```

Example:

```text
.\eecalc.exe voltage-divider --vs 12 --vo 3.3 --r1 100k
R2: 37.931034 kΩ
Power supply power delivery: 1.044 mW (87 µA)
R2 resistor % of total power: 27.5%
R2 power: 287.1 µW (87 µA)
Suggested component values for R1 (5%):
  100 kΩ
Suggested component values for R2 (5%):
  39 kΩ
```

### equivalent - Calculate equivalent components

```text
NAME:
   eecalc.exe equivalent - calculate equivalent components

USAGE:
   eecalc.exe equivalent command [command options] [arguments...]

COMMANDS:
   resistors, resistance
   capacitors, capacitance
   inductors, inductance
   help, h                  Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help (default: false)
```

Example:

```text
 .\eecalc.exe equivalent resistors --desired-value 4.7k --max-components 5 --power-handling 3W
2x (2x 4.7 kΩ in parallel @1.5 W) in series = 4.7 kΩ @3 W
2x 2.4 kΩ in series = 4.8 kΩ @3 W
2x 9.1 kΩ in parallel @1.5 W = 4.55 kΩ @3 W
3x 1.6 kΩ in series = 4.8 kΩ @3 W
3x 1.5 kΩ in series = 4.5 kΩ @3 W
4x 1.2 kΩ in series = 4.8 kΩ @3 W
5x 910 Ω in series = 4.55 kΩ @3 W
```

## How to build / install

- Make sure you have downloaded and installed Go (golang) and git
- `git clone https://github.com/PatrLind/eecalc`
- `cd eecalc`
- `go install` this will build and copy the application to your go bin folder (normally `$GOPATH/bin`, `/home/USER/go/bin`, `c:\users\USER\go\bin`).
- If the `$GOPATH/bin` folder is not in you `PATH`, copy the application to a folder that is. For example: `/usr/local/bin` or `c:\windows`
