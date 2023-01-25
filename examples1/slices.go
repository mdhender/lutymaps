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
	"fmt"

	"github.com/fogleman/ln/ln"
)

type Shape struct {
	ln.Mesh
}

func (s *Shape) Paths() ln.Paths {
	var result ln.Paths
	// for i := 0; i < 360; i++ {
	// 	fmt.Println(i)
	// 	a := ln.Radians(float64(i))
	// 	x := math.Cos(a)
	// 	y := math.Sin(a)
	// 	plane := ln.Plane{ln.Vector{}, ln.Vector{x, y, 0}}
	// 	paths := plane.IntersectMesh(&s.Mesh)
	// 	result = append(result, paths...)
	// }
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
		p := float64(i) / 100
		plane := ln.Plane{ln.Vector{0, 0, p*2 - 1}, ln.Vector{0, 0, 1}}
		result = append(result, plane.IntersectMesh(&s.Mesh)...)
		plane = ln.Plane{ln.Vector{p*2 - 1, 0, 0}, ln.Vector{1, 0, 0}}
		result = append(result, plane.IntersectMesh(&s.Mesh)...)
		plane = ln.Plane{ln.Vector{0, p*2 - 1, 0}, ln.Vector{0, 1, 0}}
		result = append(result, plane.IntersectMesh(&s.Mesh)...)
	}
	return result
}

func Slices() {
	scene := ln.Scene{}
	mesh, err := ln.LoadBinarySTL("bowser.stl")
	// mesh, err := ln.LoadOBJ("../pt/examples/bunny.obj")
	if err != nil {
		panic(err)
	}
	mesh.FitInside(ln.Box{ln.Vector{-1, -1, -1}, ln.Vector{1, 1, 1}}, ln.Vector{0.5, 0.5, 0.5})
	scene.Add(&Shape{*mesh})
	// scene.Add(mesh)
	eye := ln.Vector{-2, 2, 1}
	center := ln.Vector{0, 0, 0}
	up := ln.Vector{0, 0, 1}
	width := 1024.0 * 2
	height := 1024.0 * 2
	paths := scene.Render(eye, center, up, width, height, 50, 0.1, 100, 0.01)
	paths.WriteToPNG("slices.png", width, height)
	paths.WriteToSVG("slices.svg", width, height)
	// paths.Print()
}
