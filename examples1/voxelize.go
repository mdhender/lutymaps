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

import "github.com/fogleman/ln/ln"

func VoxelizeBowser() {
	scene := ln.Scene{}
	mesh, err := ln.LoadBinarySTL("bowser.stl")
	if err != nil {
		panic(err)
	}
	mesh.FitInside(ln.Box{ln.Vector{-1, -1, -1}, ln.Vector{1, 1, 1}}, ln.Vector{0.5, 0.5, 0.5})
	cubes := mesh.Voxelize(1.0 / 64)
	for _, cube := range cubes {
		scene.Add(cube)
	}
	eye := ln.Vector{-1, -2, 0}
	center := ln.Vector{0, 0, 0}
	up := ln.Vector{0, 0, 1}
	width := 1024.0 * 2
	height := 1024.0 * 2
	paths := scene.Render(eye, center, up, width, height, 60, 0.1, 100, 0.01)
	paths.WriteToPNG("voxelize-bowser.png", width, height)
}

func VoxelizeBunny() {
	scene := ln.Scene{}
	mesh, err := ln.LoadBinarySTL("bunny.stl")
	if err != nil {
		panic(err)
	}
	mesh.FitInside(ln.Box{ln.Vector{-1, -1, -1}, ln.Vector{1, 1, 1}}, ln.Vector{0.5, 0.5, 0.5})
	cubes := mesh.Voxelize(1.0 / 64)
	for _, cube := range cubes {
		scene.Add(cube)
	}
	eye := ln.Vector{-1, -2, 0}
	center := ln.Vector{0, 0, 0}
	up := ln.Vector{0, 0, 1}
	width := 1024.0 * 2
	height := 1024.0 * 2
	paths := scene.Render(eye, center, up, width, height, 60, 0.1, 100, 0.01)
	paths.WriteToPNG("voxelize-bunny.png", width, height)
}
