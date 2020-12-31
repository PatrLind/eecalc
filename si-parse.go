package main

import (
	"regexp"
	"strconv"
	"strings"
)

var numRegexp = regexp.MustCompile("([\\d.]+)(.*)")

func parseFloat64(v string) (float64, error) {
	numStr, mult := getMultiple(v)
	f, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return 0, err
	}
	return f * mult, nil
}

func parseInt64(v string) (int64, error) {
	f, err := parseFloat64(v)
	if err != nil {
		return 0, err
	}
	return int64(f), nil
}

func getMultiple(v string) (numStr string, mult float64) {
	subStrings := numRegexp.FindStringSubmatch(v)
	numStr = subStrings[1]
	lc := ""
	if len(subStrings) >= 2 {
		lc = strings.TrimSpace(subStrings[2])
	}
	mult = 1
	switch lc {
	case "Y":
		mult = 1000000000000000000000000
	case "Z":
		mult = 1000000000000000000000
	case "E":
		mult = 1000000000000000000
	case "P":
		mult = 1000000000000000
	case "T":
		mult = 1000000000000
	case "G":
		mult = 1000000000
	case "M":
		mult = 1000000
	case "k":
		mult = 1000
	case "h":
		mult = 100
	case "d":
		mult = 0.1
	case "c":
		mult = 0.01
	case "m":
		mult = 0.001
	case "μ":
		fallthrough
	case "u":
		mult = 0.000001
	case "p":
		mult = 0.000000001
	case "n":
		mult = 0.000000000001
	case "f":
		mult = 0.000000000000001
	case "a":
		mult = 0.000000000000000001
	case "z":
		mult = 0.000000000000000000001
	case "y":
		mult = 0.000000000000000000000001
	}
	return numStr, mult
}
