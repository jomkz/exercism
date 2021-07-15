// MIT License

// Copyright (c) 2021 John McKenzie

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package space contains the logic to calculate ages in earth years on various planets.
package space

import "strings"

type Planet string

const (
	EarthSeconds = 31557600
	RatioMercury = 0.2408467
	RatioVenus   = 0.61519726
	RatioEarth   = 1.0
	RatioMars    = 1.8808158
	RatioJupiter = 11.862615
	RatioSaturn  = 29.447498
	RatioUranus  = 84.016846
	RatioNeptune = 164.79132
)

// Age returns the age in Earth years for the given Planet.
func Age(s float64, p Planet) float64 {
	ratio := 0.0

	switch strings.ToLower(string(p)) {
	case "mercury":
		ratio = RatioMercury
	case "venus":
		ratio = RatioVenus
	case "earth":
		ratio = RatioEarth
	case "mars":
		ratio = RatioMars
	case "jupiter":
		ratio = RatioJupiter
	case "saturn":
		ratio = RatioSaturn
	case "uranus":
		ratio = RatioUranus
	case "neptune":
		ratio = RatioNeptune
	}

	return earthYears(s) / ratio
}

// earthYears returns the number of earth years as a floating point value for the given number of seconds.
func earthYears(s float64) float64 {
	return float64(s / EarthSeconds)
}
