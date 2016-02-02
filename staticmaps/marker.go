// Copyright 2016 Florian Pigorsch. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package staticmaps

import (
	"fmt"
	"image/color"
	"regexp"
	"strconv"
	"strings"

	"github.com/golang/geo/s2"
)

type Marker struct {
	Position s2.LatLng
	Color    color.RGBA
	Size     float64
}

func ParseColorString(s string) (color.RGBA, error) {
	re := regexp.MustCompile(`^\s*0x([A-Fa-f0-9]{2})([A-Fa-f0-9]{2})([A-Fa-f0-9]{2})\s*$`)
	matches := re.FindStringSubmatch(s)
	if matches != nil {
		r, errr := strconv.ParseInt(matches[1], 16, 8)
		g, errg := strconv.ParseInt(matches[2], 16, 8)
		b, errb := strconv.ParseInt(matches[3], 16, 8)
		if errr != nil || errg != nil || errb != nil {
			return color.RGBA{}, fmt.Errorf("Cannot parse color string: %s", s)
		}
		return color.RGBA{uint8(r), uint8(g), uint8(b), 0xff}, nil
	} else if s == "black" {
		return color.RGBA{0x00, 0x00, 0x00, 0xff}, nil
	} else if s == "blue" {
		return color.RGBA{0x00, 0x00, 0xff, 0xff}, nil
	} else if s == "brown" {
		return color.RGBA{0x96, 0x4b, 0x00, 0xff}, nil
	} else if s == "green" {
		return color.RGBA{0x00, 0xff, 0x00, 0xff}, nil
	} else if s == "orange" {
		return color.RGBA{0xff, 0x7f, 0x00, 0xff}, nil
	} else if s == "purple" {
		return color.RGBA{0x7f, 0x00, 0x7f, 0xff}, nil
	} else if s == "red" {
		return color.RGBA{0xff, 0x00, 0, 0xff}, nil
	} else if s == "yellow" {
		return color.RGBA{0xff, 0xff, 0x00, 0xff}, nil
	} else if s == "white" {
		return color.RGBA{0xff, 0xff, 0xff, 0xff}, nil
	} else {
		return color.RGBA{}, fmt.Errorf("Cannot parse color string: %s", s)
	}
}

func ParseSizeString(s string) (float64, error) {
	if s == "mid" {
		return 16.0, nil
	} else if s == "small" {
		return 12.0, nil
	} else if s == "tiny" {
		return 8.0, nil
	}

	return 0.0, fmt.Errorf("Cannot parse size string: %s", s)
}

func ParseMarkerString(s string) ([]Marker, error) {
	markers := make([]Marker, 0, 0)

	color := color.RGBA{0xff, 0, 0, 0xff}
	size := 16.0

	for _, ss := range strings.Split(s, "|") {
		if strings.HasPrefix(ss, "color:") {
			color_, err := ParseColorString(strings.TrimPrefix(ss, "color:"))
			if err != nil {
				return nil, err
			}
			color = color_
		} else if strings.HasPrefix(ss, "label:") {
			// TODO
		} else if strings.HasPrefix(ss, "size:") {
			size_, err := ParseSizeString(strings.TrimPrefix(ss, "size:"))
			if err != nil {
				return nil, err
			}
			size = size_
		} else {
			ll, err := ParseLatLngFromString(ss)
			if err != nil {
				return nil, err
			}
			marker := Marker{ll, color, size}
			// append marker
			n := len(markers)
			if n == cap(markers) {
				newMarkers := make([]Marker, n, 2*n+1)
				copy(newMarkers, markers)
				markers = newMarkers
			}
			markers = markers[0 : n+1]
			markers[n] = marker
		}

	}
	return markers, nil
}