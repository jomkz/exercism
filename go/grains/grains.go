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

// Package grains contains logic to calculate the number of grains of wheat on a chessboard.
package grains

import (
	"errors"
	"math"
)

// Square will return the number of grains of wheat on a specific square of a
// chessboard, given that the number on each square doubles.
func Square(n int) (uint64, error) {
	if n <= 0 || n > 64 {
		return 0, errors.New("value must be between 1 and 64")
	}

	grains := 1.0
	for i := 2; i <= n; i++ {
		grains = math.Pow(2.0, float64(n-1))
	}

	return uint64(grains), nil
}

// Total will return the total number of grains of wheat on a chessboard, given
// that the number on each square doubles.
func Total() uint64 {
	result := uint64(0)
	for i := 1; i <= 64; i++ {
		val, err := Square(i)
		if err == nil {
			result += val
		}
	}
	return result
}
