package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"
	"time"
)

import (
	_ "embed"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type Point struct {
	X, Y int
}

func rectangleSurface(p1, p2 Point) int {
	return (AbsInt(p1.X-p2.X) + 1) * (AbsInt(p1.Y-p2.Y) + 1)
}

func part1(input string) int {
	lines := strings.Split(input, "\n")

	points := make([]Point, len(lines))
	for i, l := range lines {
		s := strings.Split(l, ",")
		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])
		points[i] = Point{x, y}
	}

	maxRectangle := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			surface := rectangleSurface(points[i], points[j])
			if surface > maxRectangle {
				maxRectangle = surface
			}
		}
	}

	return maxRectangle
}

func isPointInOrOnPolygon(point Point, polygon []Point) bool {
	n := len(polygon)
	inside := false

	j := n - 1
	for i := 0; i < n; i++ {
		pi, pj := polygon[i], polygon[j]

		// Check if the point is on the edge between pi and pj
		if isPointOnLine(point, pi, pj) {
			return true
		}

		// Ray casting algorithm
		if ((pi.Y > point.Y) != (pj.Y > point.Y)) && (point.X < (pj.X-pi.X)*(point.Y-pi.Y)/(pj.Y-pi.Y)+pi.X) {
			inside = !inside
		}

		j = i
	}

	return inside
}

func isPointOnLine(point, p1, p2 Point) bool {
	minX, maxX := p1.X, p2.X
	if minX > maxX {
		minX, maxX = maxX, minX
	}
	minY, maxY := p1.Y, p2.Y
	if minY > maxY {
		minY, maxY = maxY, minY
	}

	if point.X < minX || point.X > maxX || point.Y < minY || point.Y > maxY {
		return false
	}
	return true
}

func drawPointsToPNG(points []Point, rect []Point, filename string) error {
	// Find bounds
	minX, maxX := points[0].X, points[0].X
	minY, maxY := points[0].Y, points[0].Y
	for _, p := range points {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	// Scale to reasonable image size (e.g., max 2000x2000)
	maxImageSize := 2000
	dataWidth := maxX - minX
	dataHeight := maxY - minY

	scale := 1.0
	if dataWidth > maxImageSize || dataHeight > maxImageSize {
		scaleX := float64(maxImageSize) / float64(dataWidth)
		scaleY := float64(maxImageSize) / float64(dataHeight)
		if scaleX < scaleY {
			scale = scaleX
		} else {
			scale = scaleY
		}
	}

	padding := 20
	width := int(float64(dataWidth)*scale) + 2*padding
	height := int(float64(dataHeight)*scale) + 2*padding

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// White background
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.White)
		}
	}

	// Draw rectangle
	if len(rect) > 0 {
		green := color.RGBA{0, 255, 0, 255}
		for i := 0; i < len(rect); i++ {
			x1 := int(float64(rect[i].X-minX)*scale) + padding
			y1 := int(float64(rect[i].Y-minY)*scale) + padding
			x2 := int(float64(rect[(i+1)%len(rect)].X-minX)*scale) + padding
			y2 := int(float64(rect[(i+1)%len(rect)].Y-minY)*scale) + padding
			drawLine(img, x1, y1, x2, y2, 2, green)
		}

		// Draw rectangle points larger so they stand out
		for _, p := range rect {
			x := int(float64(p.X-minX)*scale) + padding
			y := int(float64(p.Y-minY)*scale) + padding
			for dy := -3; dy <= 3; dy++ {
				for dx := -3; dx <= 3; dx++ {
					if x+dx >= 0 && x+dx < width && y+dy >= 0 && y+dy < height {
						img.Set(x+dx, y+dy, green)
					}
				}
			}
		}
	}

	// Draw lines first (so they're behind the points)
	blue := color.RGBA{0, 0, 255, 255}
	for i := 0; i < len(points)-1; i++ {
		x1 := int(float64(points[i].X-minX)*scale) + padding
		y1 := int(float64(points[i].Y-minY)*scale) + padding
		x2 := int(float64(points[i+1].X-minX)*scale) + padding
		y2 := int(float64(points[i+1].Y-minY)*scale) + padding
		drawLine(img, x1, y1, x2, y2, 1, blue)
	}

	// Draw points
	red := color.RGBA{255, 0, 0, 255}
	pointSize := 2
	if scale < 0.5 {
		pointSize = 1 // Smaller points for zoomed out view
	}
	for _, p := range points {
		x := int(float64(p.X-minX)*scale) + padding
		y := int(float64(p.Y-minY)*scale) + padding
		for dy := -pointSize; dy <= pointSize; dy++ {
			for dx := -pointSize; dx <= pointSize; dx++ {
				if x+dx >= 0 && x+dx < width && y+dy >= 0 && y+dy < height {
					img.Set(x+dx, y+dy, red)
				}
			}
		}
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Printf("Created PNG: %dx%d (scale: %.3f)\n", width, height, scale)
	return png.Encode(f, img)
}

func drawLine(img *image.RGBA, x1, y1, x2, y2, lineWidth int, col color.Color) {
	// Simple Bresenham's line algorithm
	dx := AbsInt(x2 - x1)
	dy := AbsInt(y2 - y1)
	sx, sy := 1, 1
	if x1 > x2 {
		sx = -1
	}
	if y1 > y2 {
		sy = -1
	}
	err := dx - dy

	for {
		// Draw a thicker line by drawing around the center point
		for dy := -lineWidth / 2; dy <= lineWidth/2; dy++ {
			for dx := -lineWidth / 2; dx <= lineWidth/2; dx++ {
				img.Set(x1+dx, y1+dy, col)
			}
		}

		if x1 == x2 && y1 == y2 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x1 += sx
		}
		if e2 < dx {
			err += dx
			y1 += sy
		}
	}
}

func isRectangleInPolygon(corner1, corner2 Point, polygon []Point) bool {
	// Calculate the other two corners
	corner3 := Point{corner1.X, corner2.Y}
	corner4 := Point{corner2.X, corner1.Y}

	corners := []Point{corner1, corner3, corner2, corner4}

	// Check all corners are in or on the polygon
	for _, corner := range corners {
		if !isPointInOrOnPolygon(corner, polygon) {
			return false
		}
	}

	// Check if any rectangle edge intersects with any polygon edge
	// (excluding touches at shared vertices)
	rectEdges := [][2]Point{
		{corner1, corner3},
		{corner3, corner2},
		{corner2, corner4},
		{corner4, corner1},
	}

	n := len(polygon)
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		polyEdge := [2]Point{polygon[i], polygon[j]}

		for _, rectEdge := range rectEdges {
			if lineIntersectsLine(rectEdge[0], rectEdge[1], polyEdge[0], polyEdge[1]) {
				return false
			}
		}
	}

	return true
}

func lineIntersectsLine(p1, p2, p3, p4 Point) bool {
	// Check if segments p1-p2 and p3-p4 properly intersect (cross each other)
	// Not just touch at endpoints

	d1 := direction(p3, p4, p1)
	d2 := direction(p3, p4, p2)
	d3 := direction(p1, p2, p3)
	d4 := direction(p1, p2, p4)

	// Proper intersection: points are on opposite sides of each line
	if ((d1 > 0 && d2 < 0) || (d1 < 0 && d2 > 0)) &&
		((d3 > 0 && d4 < 0) || (d3 < 0 && d4 > 0)) {
		return true
	}

	return false
}

func direction(p1, p2, p3 Point) int {
	// Cross product to determine which side p3 is on relative to line p1-p2
	val := (p2.Y-p1.Y)*(p3.X-p2.X) - (p2.X-p1.X)*(p3.Y-p2.Y)
	if val == 0 {
		return 0 // Collinear
	}
	if val > 0 {
		return 1 // Clockwise/right
	}
	return -1 // Counterclockwise/left
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	p := make([]Point, len(lines))
	for i, l := range lines {
		s := strings.Split(l, ",")
		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])
		p[i] = Point{x, y}
	}

	var bestRect []Point
	maxRectangle := 0
	for i := 0; i < len(p); i++ {
		for j := i + 1; j < len(p); j++ {

			if !isRectangleInPolygon(p[i], p[j], p) {
				continue
			}

			surface := rectangleSurface(p[i], p[j])
			if surface > maxRectangle {
				maxRectangle = surface
				bestRect = []Point{p[i], {p[j].X, p[i].Y}, p[j], {p[i].X, p[j].Y}}
			}
		}
	}

	drawPointsToPNG(p, bestRect, "out.png")

	// 181508672 too low
	// 1572047142
	// 158707536 too low
	// 110687382 too low
	return maxRectangle
}

func main() {

	inputPtr := flag.Bool("input", false, "sample or input")

	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")

	flag.Parse()

	var inputText string
	if *inputPtr {
		inputText = strings.TrimSpace(input)
		fmt.Println("Running part", part, "on input.txt.")
	} else {
		inputText = strings.TrimSpace(sample)
		fmt.Println("Running part", part, "on sample.txt.")
	}

	start := time.Now()
	if part == 1 {
		fmt.Println("Result:", part1(inputText))
	} else {
		fmt.Println("Result:", part2(inputText))
	}
	fmt.Println("Time:", fmt.Sprintf("%d ms", time.Since(start).Milliseconds()))
}
