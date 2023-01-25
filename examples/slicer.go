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
	"fmt"

	"github.com/fogleman/ln/ln"
)

const Slicers = 32
const Size = 1024

func SliceBowser() {
	mesh, err := ln.LoadBinarySTL("bowser.stl")
	if err != nil {
		panic(err)
	}
	mesh.FitInside(ln.Box{ln.Vector{-1, -1, -1}, ln.Vector{1, 1, 1}}, ln.Vector{0.5, 0.5, 0.5})
	for i := 0; i < Slicers; i++ {
		fmt.Printf("bowser/slice%04d\n", i)
		p := (float64(i)/(Slicers-1))*2 - 1
		point := ln.Vector{0, 0, p}
		plane := ln.Plane{point, ln.Vector{0, 0, 1}}
		paths := plane.IntersectMesh(mesh)
		paths = paths.Transform(ln.Scale(ln.Vector{Size / 2, Size / 2, 1}).Translate(ln.Vector{Size / 2, Size / 2, 0}))
		paths.WriteToPNG(fmt.Sprintf("bowser/slice%04d.png", i), Size, Size)
	}
}

func SliceSuzanne() {
	mesh, err := ln.LoadOBJ("suzanne.obj")
	if err != nil {
		panic(err)
	}
	mesh.FitInside(ln.Box{ln.Vector{-1, -1, -1}, ln.Vector{1, 1, 1}}, ln.Vector{0.5, 0.5, 0.5})
	for i := 0; i < Slicers; i++ {
		fmt.Printf("suzanne/slice%04d\n", i)
		p := (float64(i)/(Slicers-1))*2 - 1
		point := ln.Vector{0, 0, p}
		plane := ln.Plane{point, ln.Vector{0, 0, 1}}
		paths := plane.IntersectMesh(mesh)
		paths = paths.Transform(ln.Scale(ln.Vector{Size / 2, Size / 2, 1}).Translate(ln.Vector{Size / 2, Size / 2, 0}))
		paths.WriteToPNG(fmt.Sprintf("suzanne/slice%04d.png", i), Size, Size)
	}
}
