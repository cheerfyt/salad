package goms

import (
	"errors"
	"regexp"
	"strconv"
)

const (
	s = 1000
	m = s * 60
	h = m * 60
	d = h * 24
	y = d * 365
)

const re = `^((?:\d+)?\.?\d+) *(microseconds?|microsecond|μs|milliseconds?|msecs?|ms|seconds?|secs?|s|minutes?|mins?|m|hours?|hrs?|h|days?|d|years?|yrs?|y)?$`

//Ms export ms func
func Ms(val string) (float32, error) {
	if len(val) > 15 {
		return 0, errors.New("string too long")
	}
	reg, _ := regexp.Compile(re)
	match := reg.FindStringSubmatch(val)
	if len(match) < 3 {
		return 0, errors.New("Match Error")
	}
	real, err := strconv.ParseFloat(match[1], 32)
	if err != nil {
		return 0.0, errors.New("parse to float error")
	}
	switch match[2] {
	case "years", "year", "yrs", "yr", "y":
		return float32(real * y), nil
	case "days", "day", "d":
		return float32(real * d), nil
	case "hours", "hour", "hrs", "hr", "h":
		return float32(real * h), nil
	case "minutes", "minute", "min", "m":
		return float32(real * m), nil
	case "seconds", "second", "secs", "sec", "s":
		return float32(real * s), nil
	case "milliseconds", "millisecond", "msecs", "msec", "ms":
		return float32(real), nil
	default:
		return 0, errors.New("unit not support")
	}
}