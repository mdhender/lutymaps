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

package examples

import (
	"math"
	"math/rand"
	"time"

	"github.com/fogleman/ln/ln"
)

func Beads() {
	rand.Seed(1211)
	rand.Seed(time.Now().UnixNano())

	eye := ln.Vector{8, 8, 8}
	center := ln.Vector{0, 0, 0}
	up := ln.Vector{0, 0, 1}
	scene := ln.Scene{}

	n := 200
	for a := 0; a < 50; a++ {
		xs := LowPassNoise(n, 0.3, 4)
		ys := LowPassNoise(n, 0.3, 4)
		zs := LowPassNoise(n, 0.3, 4)
		ss := LowPassNoise(n, 0.3, 4)
		position := ln.Vector{}
		for i := 0; i < n; i++ {
			sphere := ln.NewOutlineSphere(eye, up, position, 0.1)
			scene.Add(sphere)
			s := (ss[i]+1)/2*0.1 + 0.01
			v := ln.Vector{xs[i], ys[i], zs[i]}.Normalize().MulScalar(s)
			position = position.Add(v)
		}
	}

	width := 380.0 * 5
	height := 315.0 * 5
	fovy := 50.0
	paths := scene.Render(eye, center, up, width, height, fovy, 0.1, 100, 0.01)
	paths.WriteToPNG("beads.png", width, height)
	//paths.Print()
}

func Normalize(values []float64, a, b float64) []float64 {
	result := make([]float64, len(values))
	lo := values[0]
	hi := values[0]
	for _, x := range values {
		lo = math.Min(lo, x)
		hi = math.Max(hi, x)
	}
	for i, x := range values {
		p := (x - lo) / (hi - lo)
		result[i] = a + p*(b-a)
	}
	return result
}

func LowPass(values []float64, alpha float64) []float64 {
	result := make([]float64, len(values))
	var y float64
	for i, x := range values {
		y -= alpha * (y - x)
		result[i] = y
	}
	return result
}

func LowPassNoise(n int, alpha float64, iterations int) []float64 {
	result := make([]float64, n)
	for i := range result {
		result[i] = rand.Float64()
	}
	for i := 0; i < iterations; i++ {
		result = LowPass(result, alpha)
	}
	result = Normalize(result, -1, 1)
	return result
}
