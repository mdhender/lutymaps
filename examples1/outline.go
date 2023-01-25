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

func Outline() {
	eye := ln.Vector{8, 8, 8}
	center := ln.Vector{0, 0, 0}
	up := ln.Vector{0, 0, 1}
	scene := ln.Scene{}
	n := 10
	for x := -n; x <= n; x++ {
		for y := -n; y <= n; y++ {
			z := rand.Float64() * 3
			v := ln.Vector{float64(x), float64(y), z}
			sphere := ln.NewOutlineSphere(eye, up, v, 0.45)
			scene.Add(sphere)
		}
	}
	width := 1920.0
	height := 1200.0
	fovy := 50.0
	paths := scene.Render(eye, center, up, width, height, fovy, 0.1, 100, 0.01)
	paths.WriteToPNG("outline.png", width, height)
}
