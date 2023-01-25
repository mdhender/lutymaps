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

package examples1

import (
	"math/rand"

	"github.com/fogleman/ln/ln"
)

func cube(x, y, z float64) ln.Shape {
	size := 0.5
	v := ln.Vector{x, y, z}
	return ln.NewCube(v.SubScalar(size), v.AddScalar(size))
}

func Example1() {
	scene := ln.Scene{}
	for x := -2; x <= 2; x++ {
		for y := -2; y <= 2; y++ {
			z := rand.Float64()
			scene.Add(cube(float64(x), float64(y), z))
		}
	}
	eye := ln.Vector{6, 5, 3}
	center := ln.Vector{0, 0, 0}
	up := ln.Vector{0, 0, 1}
	width := 1920.0
	height := 1200.0
	fovy := 30.0
	paths := scene.Render(eye, center, up, width, height, fovy, 0.1, 100, 0.01)
	paths.WriteToPNG("example1.png", width, height)
	paths.WriteToSVG("example1.svg", width, height)
}
