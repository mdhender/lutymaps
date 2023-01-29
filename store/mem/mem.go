/*
 * lutymaps - a mapping engine for luty
 *
 * Copyright (c) 2023 2023 Michael D Henderson
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

// Package mem implements an in-memory data store.
package mem

// Store implements an in-memory data store.
type Store struct {
	Systems Systems
}

type Systems []*System

// System implements the data for a system.
type System struct {
	X, Y, Z int
	Kind    SystemKind
}

func (s *System) Points() (float64, float64, float64) {
	if s == nil {
		return 0, 0, 0
	}
	return float64(s.X), float64(s.Y), float64(s.Z)
}

// SystemKind is an enum for the type of system
type SystemKind int

const (
	SKEmpty SystemKind = iota
	SKBlueSuperGiant
	SKDenseDustCloud
	SKMediumDustCloud
	SKYellowMainSequence
)

// String implements the Stringer interface.
func (sk SystemKind) String() string {
	switch sk {
	case SKEmpty:
		return "Empty"
	case SKBlueSuperGiant:
		return "Blue Super Giant"
	case SKDenseDustCloud:
		return "Dense Dust Cloud"
	case SKMediumDustCloud:
		return "Medium Dust Cloud"
	case SKYellowMainSequence:
		return "Yellow Main Sequence"
	}
	return ""
}

func (s *Store) Filter(fn func(*System) bool) Systems {
	var systems Systems
	for _, system := range s.Systems {
		if !fn(system) {
			continue
		}
		systems = append(systems, system)
	}
	return systems
}
