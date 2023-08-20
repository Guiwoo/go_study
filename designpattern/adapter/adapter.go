package adapter

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strings"
)

/**
Setting up the scenario
*/

type Line struct {
	X1, Y1, X2, Y2 int
}

type VectorImage struct {
	Line []Line
}

func NewRectangle(width, height int) *VectorImage {
	width -= 1
	height -= 1
	return &VectorImage{[]Line{
		{0, 0, width, 0},
		{0, 0, 0, height},
		{width, 0, width, height},
		{0, height, width, height},
	}}
}

// This is the interface you are given

// Point the interface we have
type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
	maxX, maxY := 0, 0
	points := owner.GetPoints()
	for _, pixel := range points {
		if pixel.X > maxX {
			maxX = pixel.X
		}
		if pixel.Y > maxY {
			maxY = pixel.Y
		}
	}

	maxX += 1
	maxY += 1

	data := make([][]byte, maxY)
	for i := 0; i < maxY; i++ {
		data[i] = make([]byte, maxX)
		for j := 0; j < maxX; j++ {
			data[i][j] = ' '
		}
	}

	for _, points := range points {
		data[points.Y][points.X] = '*'
	}

	b := strings.Builder{}
	for _, line := range data {
		b.Write(line)
		b.WriteRune('\n')
	}

	return b.String()
}

type vectorToRasterAdapter struct {
	points []Point
}

func (v *vectorToRasterAdapter) GetPoints() []Point {
	return v.points
}

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

var pointCache = map[[16]byte][]Point{}

func (v *vectorToRasterAdapter) addLine(line Line) {
	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			v.points = append(v.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			v.points = append(v.points, Point{x, top})
		}
	}
}

func (v *vectorToRasterAdapter) addLineCache(line Line) {
	hash := func(obj interface{}) [16]byte {
		bytes, _ := json.Marshal(obj)
		return md5.Sum(bytes)
	}

	h := hash(line)

	if pp, ok := pointCache[h]; ok {
		for _, pt := range pp {
			v.points = append(v.points, pt)
		}
		return
	}

	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			v.points = append(v.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			v.points = append(v.points, Point{x, top})
		}
	}
	pointCache[h] = v.points
}

func VectorToRaster(vi *VectorImage) RasterImage {
	adapter := vectorToRasterAdapter{}
	for _, line := range vi.Line {
		adapter.addLineCache(line)
	}
	return &adapter
}

func Start() {
	rc := NewRectangle(10, 5)
	a := VectorToRaster(rc)
	fmt.Print(DrawPoints(a))
}
