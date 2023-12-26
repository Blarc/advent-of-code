package shared

import "strconv"

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func MustAtoi(s string) int {
	number, _ := strconv.Atoi(s)
	return number
}

type Point struct {
	X, Y int
}

func (p Point) Left() Point {
	if Abs(p.X) > 1 || Abs(p.Y) > 1 {
		panic("Point is not a direction.")
	}
	return Point{p.Y, -p.X}
}

func (p Point) Right() Point {
	if Abs(p.X) > 1 || Abs(p.Y) > 1 {
		panic("Point is not a direction.")
	}
	return Point{-p.Y, p.X}
}

func (p Point) Straight() Point {
	if Abs(p.X) > 1 || Abs(p.Y) > 1 {
		panic("Point is not a direction.")
	}
	return Point{p.X, p.Y}
}

func (p Point) InsideArrayGrid(grid [][]int) bool {
	return 0 <= p.X && p.X < len(grid[0]) && 0 <= p.Y && p.Y < len(grid)
}

func (p Point) InsideMapGrid(grid map[int]map[int]any) bool {
	return 0 <= p.Y && p.Y <= len(grid) && 0 <= p.X && p.X < len(grid[p.Y])
}
