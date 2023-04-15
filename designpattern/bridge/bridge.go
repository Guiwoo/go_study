package bridge

import "fmt"

/**
Cartesian product
Common type ThreadScheduler

A mechanism that decouples an interface from an implementation(hierarchy).
*/

// Circle, square
// Raster, vector

// RasterCircle, VectorCircle, RasterSquare, VectorSquare

type Renderer interface {
	RenderCircle(radius float32)
}

type VectorRenderer struct {
	// any kind of utility functions
}

func (v VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of radius", radius)
}

type RasterRenderer struct {
	// pixel ?
	Dpi int
}

func (r RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing pixels for a circle of radius", radius)
}

type Circle struct {
	renderer Renderer
	radius   float32
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer, radius}
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}
func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}

var _ Renderer = (*RasterRenderer)(nil)
var _ Renderer = (*VectorRenderer)(nil)

func Start() {
	raster := RasterRenderer{}
	vector := VectorRenderer{}

	circle := NewCircle(raster, 5)
	circle.Draw()
	circle.Resize(2)

	circle.renderer = vector
	circle.Draw()
}
