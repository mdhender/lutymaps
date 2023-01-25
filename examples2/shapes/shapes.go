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

package shapes

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	. "github.com/fogleman/fauxgl"
	"github.com/nfnt/resize"
)

const (
	scale  = 4    // optional supersampling
	width  = 1600 // output width in pixels
	height = 1600 // output height in pixels
	fovy   = 30   // vertical field of view in degrees
	near   = 1    // near clipping plane
	far    = 100  // far clipping plane
)

var (
	eye        = V(3*4, 3*4, 1.5*4)          // camera position
	center     = V(0, 0, 0)                  // view center position
	up         = V(0, 0, 1)                  // up vector
	light      = V(0.75, 0.5, 1).Normalize() // light direction
	color      = HexColor("#468966")         // object color
	background = HexColor("#FFF8E3")         // background color
)

func randXYZ(radius float64) (x, y, z float64) {
	for {
		x = rand.Float64()*2*radius - radius
		y = rand.Float64()*2*radius - radius
		z = rand.Float64()*2*radius - radius
		if x*x+y*y+z*z < radius {
			return x, y, z
		}
	}
}

func Main() {
	start := time.Now()

	mesh, radius := NewEmptyMesh(), 10.0
	for i := 0; i < 1_500; i++ {
		//var x, y, z float64
		//for {
		//	x = rand.Float64()*2 - 1
		//	y = rand.Float64()*2 - 1
		//	z = rand.Float64()*2 - 1
		//	if x*x+y*y+z*z < 1 {
		//		break
		//	}
		//}
		x, y, z := randXYZ(radius)
		p := Vector{x, y, z}.MulScalar(4)
		s := V(0.2, 0.2, 0.2)
		u := RandomUnitVector()
		a := rand.Float64() * 2 * math.Pi
		c := NewCube()
		c.Transform(Orient(p, s, u, a))
		mesh.Add(c)
	}
	fmt.Println("shapes", time.Since(start))

	// create a rendering context
	context := NewContext(width*scale, height*scale)
	context.ClearColorBufferWith(Black)
	fmt.Println("shapes", time.Since(start))

	// create transformation matrix and light direction
	aspect := float64(width) / float64(height)
	matrix := LookAt(eye, center, up).Perspective(fovy, aspect, near, far)
	fmt.Println("shapes", time.Since(start))

	// render
	shader := NewPhongShader(matrix, light, eye)
	shader.ObjectColor = color
	context.Shader = shader
	context.DrawMesh(mesh)
	fmt.Println("shapes", time.Since(start))

	mesh = NewEmptyMesh()
	for n := -10; n < 0; n = n + 2 {
		sp, err := LoadSTL("sphere.stl")
		if err != nil {
			panic(err)
		}
		sp.SmoothNormals()
		sp.Transform(Scale(V(0.5, 0.5, 0.5)))
		sp.Transform(Translate(V(float64(n), float64(n), -0.0)))
		mesh.Add(sp)
	}
	shader = NewPhongShader(matrix, light, eye)
	shader.ObjectColor = HexColor("FF9DFF").Alpha(0.65)
	shader.SpecularPower = 0
	context.Shader = shader
	context.DrawMesh(mesh)
	context.Wireframe = true
	context.DepthBias = -0.00001
	context.DrawMesh(mesh)
	fmt.Println("shapes", time.Since(start))

	mesh = NewEmptyMesh()
	for n := 0; n < 5; n++ {
		x, y, z := randXYZ(radius)
		sp, err := LoadSTL("sphere.stl")
		if err != nil {
			panic(err)
		}
		sp.SmoothNormals()
		sp.Transform(Scale(V(0.5, 0.5, 0.5)))
		sp.Transform(Translate(V(x, y, z)))
		mesh.Add(sp)
	}
	shader = NewPhongShader(matrix, light, eye)
	shader.ObjectColor = HexColor("FFFF9D").Alpha(0.65)
	shader.SpecularPower = 0
	context.Shader = shader
	context.DrawMesh(mesh)
	context.Wireframe = true
	context.DepthBias = -0.00001
	context.DrawMesh(mesh)
	fmt.Println("shapes", time.Since(start))

	mesh = NewEmptyMesh()
	for n := 0; n < 5; n++ {
		x, y, z := randXYZ(radius)
		sp, err := LoadSTL("sphere.stl")
		if err != nil {
			panic(err)
		}
		sp.SmoothNormals()
		sp.Transform(Scale(V(0.25, 0.25, 0.25)))
		sp.Transform(Translate(V(x, y, z)))
		mesh.Add(sp)
	}
	shader3 := NewPhongShader(matrix, light, eye)
	shader3.ObjectColor = HexColor("9DFFFF").Alpha(0.65)
	shader3.SpecularPower = 0
	context.Shader = shader3
	context.DrawMesh(mesh)
	context.Wireframe = true
	context.DepthBias = -0.00001
	context.DrawMesh(mesh)
	fmt.Println("shapes", time.Since(start))

	// downsample image for antialiasing
	image := context.Image()
	image = resize.Resize(width, height, image, resize.Bilinear)
	fmt.Println("shapes", time.Since(start))

	// save image
	SavePNG("shapes.png", image)
	fmt.Println("shapes", time.Since(start))
}
