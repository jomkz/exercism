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

// Package letter contains the logic to counts the frequency of each rune in a given text.
package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// ConcurrentFrequency counts the frequency of each rune in the given slice of
// text concurrently and returns this data as a FreqMap.
func ConcurrentFrequency(ss []string) FreqMap {
	c := make(chan FreqMap)
	m := FreqMap{}
	for _, s := range ss {
		go func(s string, c chan FreqMap) {
			c <- Frequency(s)
		}(s, c)

		cm := <-c
		for k, v := range cm {
			m[k] += v
		}
	}
	return m
}

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}
