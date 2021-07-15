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

// Package bob contains the logic for responding as Bob.
package bob

import (
	"strings"
	"unicode"
)

const (
	ResponseDefault         = "Whatever."
	ResponseEmpty           = "Fine. Be that way!"
	ResponseGeneralQuestion = "Sure."
	ResponseGeneralYelling  = "Whoa, chill out!"
	ResponseYellingQuestion = "Calm down, I know what I'm doing!"
)

// Hey will return the appropriate response based on the given remark.
func Hey(remark string) string {
	sanitized := strings.TrimSpace(remark)

	if len(sanitized) <= 0 {
		return ResponseEmpty
	}

	if isQuestion(sanitized) {
		if isAllCaps(sanitized) {
			return ResponseYellingQuestion
		} else {
			return ResponseGeneralQuestion
		}
	} else if isAllCaps(sanitized) {
		return ResponseGeneralYelling
	}

	return ResponseDefault
}

// isAllCaps will return true when the given value contains letters that are all uppercase.
func isAllCaps(s string) bool {
	letterFound := false
	lowerFound := false

	for _, c := range s {
		if unicode.IsLetter(c) {
			letterFound = true
		}
		if unicode.IsLower(c) {
			lowerFound = true
			break
		}
	}

	return letterFound && !lowerFound
}

// isQuestion will return true if the given value ends with a question mark.
func isQuestion(s string) bool {
	return strings.HasSuffix(s, "?")
}
