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

// Package scan implements a sector scan.
package scan

import (
	"fmt"
	gl "github.com/fogleman/fauxgl"
	"github.com/mdhender/lutymaps/pkg/stores/mem"
	"github.com/nfnt/resize"
	"time"
)

const (
	scale  = 4    // optional supersampling
	width  = 3200 // output width in pixels
	height = 3200 // output height in pixels
	fovy   = 60   // vertical field of view in degrees
	near   = 1    // near clipping plane
	far    = 100  // far clipping plane
)

var (
	eye        = gl.V(50, 50, 0)                // camera position
	center     = gl.V(0, 0, 0)                  // view center position
	up         = gl.V(0, 0, 1)                  // up vector
	light      = gl.V(0.75, 0.5, 1).Normalize() // light direction
	color      = gl.HexColor("#468966")         // object color
	background = gl.HexColor("#FFF8E3")         // background color
)

func New(store *mem.Store, filter func(*mem.System) bool, path string) error {
	start := time.Now()
	defer func(s time.Time) {
		fmt.Printf("scan: %v\n", time.Since(s))
	}(start)

	systems := store.Filter(filter)
	fmt.Printf("scan: systems %d\n", len(systems))

	// create a rendering context
	context := gl.NewContext(width*scale, height*scale)
	context.ClearColorBufferWith(gl.Black)

	// create transformation matrix and light direction
	aspect := float64(width) / float64(height)
	matrix := gl.LookAt(eye, center, up).Perspective(fovy, aspect, near, far)

	// create a mesh for the grid points
	gridMesh := gl.NewEmptyMesh()
	for x := -30; x <= 30; x = x + 10 {
		for y := -30; y <= 30; y = y + 10 {
			for z := -30; z <= 30; z = z + 10 {
				p := gl.Vector{float64(x), float64(y), float64(z)}.MulScalar(4)
				s := gl.V(0.2, 0.2, 0.2)
				u := gl.Vector{1, 1, 1}.Normalize()
				a := 0.0 // rand.Float64() * 2 * math.Pi
				c := gl.NewCube()
				c.Transform(gl.Orient(p, s, u, a))
				gridMesh.Add(c)
			}
		}
	}
	// render the grid mesh
	gridShader := gl.NewPhongShader(matrix, light, eye)
	gridShader.ObjectColor = color
	context.Shader = gridShader
	context.DrawMesh(gridMesh)
	fmt.Printf("scan: grid lines: %v\n", time.Since(start))

	// create a mesh for the stars
	starMesh := gl.NewEmptyMesh()
	for _, sys := range systems {
		if sys == nil {
			break
		}
		sp, err := gl.LoadSTL("sphere.stl")
		if err != nil {
			panic(err)
		}
		sp.SmoothNormals()
		sp.Transform(gl.Scale(gl.V(0.4, 0.4, 0.4)))
		x, y, z := sys.Points()
		sp.Transform(gl.Translate(gl.V(x, y, z)))
		starMesh.Add(sp)
	}
	// render the starmesh
	starShader := gl.NewPhongShader(matrix, light, eye)
	starShader.ObjectColor = gl.HexColor("9DFFFF").Alpha(0.75)
	starShader.SpecularPower = 0
	context.Shader = starShader
	context.DrawMesh(starMesh)
	context.Wireframe = true
	context.DepthBias = -0.00001
	context.DrawMesh(starMesh)
	fmt.Printf("scan: stars: %5d: %v\n", len(systems), time.Since(start))

	// down-sample image for antialiasing
	img := resize.Resize(width, height, context.Image(), resize.Bilinear)
	err := gl.SavePNG(path, img)
	if err != nil {
		return fmt.Errorf("scan: %w", err)
	}

	fmt.Printf("scan: created %q\n", path)
	return nil
}
