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

// Package luhn contains the logic to determine whether or not a number string
// is valid per the Luhn formula.
package luhn

import (
	"unicode"
)

// Valid will return true if the given string is valid per the Luhn formula.
func Valid(s string) bool {
	values := make([]int, 0)
	for _, c := range s {
		if !unicode.IsSpace(c) && !unicode.IsNumber(c) {
			return false
		}
		if unicode.IsNumber(c) {
			values = append(values, int(c-'0')) // Get int value of rune
		}
	}

	if len(values) <= 1 {
		return false
	}

	dbl := false
	for i := len(values) - 1; i >= 0; i-- {
		if dbl {
			dbl = false
			val := values[i] * 2
			if val > 9 {
				val = val - 9
			}
			values[i] = val
		} else {
			dbl = true
		}
	}

	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum%10 == 0
}
