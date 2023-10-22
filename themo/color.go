package themo

import (
	"strconv"
	"strings"
)

type Color struct {
	R int
	G int
	B int
}

func ParseColor(color string) (*Color, error) {
	color = strings.TrimSpace(color)
	color = strings.TrimPrefix(color, "rgb(")
	color = strings.TrimSuffix(color, ")")

	var r int
	var g int
	var b int

	parts := strings.Split(color, ",")
	if val, err := strconv.Atoi(strings.TrimSpace(parts[0])); err != nil {
		return nil, err
	} else {
		r = val
	}
	if val, err := strconv.Atoi(strings.TrimSpace(parts[1])); err != nil {
		return nil, err
	} else {
		g = val
	}
	if val, err := strconv.Atoi(strings.TrimSpace(parts[2])); err != nil {
		return nil, err
	} else {
		b = val
	}

	return &Color{R: r, G: g, B: b}, nil
}
