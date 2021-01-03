package main

import (
	"errors"
	"regexp"

	"github.com/dustin/go-humanize"
)

var errInvalid = errors.New("invalid input")

// Same as humanize but includes u as an alias for µ
var riParseRegex = regexp.MustCompile(`^([\-0-9.]+)\s?([yzafpnµumkMGTPEZY]?)(.*)`)

func parseSI(v string) (float64, string, error) {
	if v == "" {
		return 0, "", nil
	}

	found := riParseRegex.FindStringSubmatch(v)
	if len(found) != 4 {
		return 0, "", errInvalid
	}
	if found[2] == "u" {
		v = found[1] + "µ" + found[3]
	}
	return humanize.ParseSI(v)
}
