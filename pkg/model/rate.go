package model

import (
	"fmt"
	"regexp"
	"strconv"
)

// Duration wraps time.Duration. It is used to parse the custom duration format
// from YAML.
// This type should not propagate beyond the scope of input/output processing.
type Rate int64

// Set implements pflag/flag.Value
func (d *Rate) Set(s string) error {
	var err error
	*d, err = ParseRate(s)
	return err
}

// Type implements pflag.Value
func (d *Rate) Type() string {
	return "rate"
}

var rateRE = regexp.MustCompile("^([0-9]+)([mMKkB])$")

// ParseRate parses a string into a int64.
func ParseRate(rateStr string) (Rate, error) {
	matches := rateRE.FindStringSubmatch(rateStr)
	if len(matches) != 3 {
		return 0, fmt.Errorf("not a valid duration string: %q", rateStr)
	}
	var (
		n, _ = strconv.Atoi(matches[1])
	)
	switch unit := matches[2]; {
	case unit == "g" || unit == "G":
		n *= 1024 * 1024 * 1024
	case unit == "m" || unit == "M":
		n *= 1024 * 1024
	case unit == "k" || unit == "K":
		n *= 1024
	case unit == "B":
		// Value already correct
	default:
		return 0, fmt.Errorf("invalid unit in rate string: %q", unit)
	}
	return Rate(n), nil
}

func (d Rate) String() string {
	var (
		n    = int64(d)
		unit = "B"
	)
	if n == 0 {
		return "0B"
	}
	factors := map[string]int64{
		"G": 1024 * 1024 * 1024,
		"M": 1024 * 1024,
		"K": 1024,
		"B": 1,
	}

	switch int64(0) {
	case n % factors["G"]:
		unit = "G"
	case n % factors["M"]:
		unit = "M"
	case n % factors["K"]:
		unit = "K"
	case n % factors["B"]:
		unit = "B"
	}
	return fmt.Sprintf("%v%v", n/factors[unit], unit)
}

// MarshalYAML implements the yaml.Marshaler interface.
func (d Rate) MarshalYAML() (interface{}, error) {
	return d.String(), nil
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (d *Rate) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	rate, err := ParseRate(s)
	if err != nil {
		return err
	}
	*d = rate
	return nil
}
