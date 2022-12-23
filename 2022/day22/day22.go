package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

type Side struct {
	y int
	x int

	up    string
	down  string
	left  string
	right string

	moveUp    func(y, x, a int) (int, int, string)
	moveDown  func(y, x, a int) (int, int, string)
	moveLeft  func(y, x, a int) (int, int, string)
	moveRight func(y, x, a int) (int, int, string)

	walls map[string]bool
}

func createKey(y, x int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func part1(input string) int {

	surface := 0
	for _, c := range input {
		if c == '.' || c == '#' {
			surface += 1
		}
	}
	a := int(math.Sqrt(float64(surface) / 6))

	sides := make(map[string]Side)

	maxY := make(map[int]int)
	maxX := make(map[int]int)

	minY := make(map[int]int)
	minX := make(map[int]int)

	lines := strings.Split(input, "\n")
	for y := 0; y < len(lines)-2; y++ {
		x := 0
		line := lines[y]
		for line[x] == ' ' {
			x += a
		}

		for x < len(line) {
			key := createKey(y/a, x/a)

			side, sideExists := sides[key]
			if !sideExists {
				sY := y / a
				sX := x / a

				side = Side{}
				side.y = sY
				side.x = sX
				side.walls = make(map[string]bool)
				sides[key] = side

				if v, e := maxX[sY]; !e {
					maxX[sY] = sX
				} else if sX > v {
					maxX[sY] = sX
				}

				if v, e := minX[sY]; !e {
					minX[sY] = sX
				} else if sX < v {
					minX[sY] = sX
				}

				if v, e := maxY[sX]; !e {
					maxY[sX] = sY
				} else if sY > v {
					maxY[sX] = sY
				}

				if v, e := minY[sX]; !e {
					minY[sX] = sY
				} else if sY < v {
					minY[sX] = sY
				}
			}

			if line[x] == '#' {
				side.walls[createKey(y%a, x%a)] = true
			}
			x++
		}
	}

	for me, side := range sides {

		var up, down, right, left string
		if side.y-1 < minY[side.x] {
			up = createKey(maxY[side.x], side.x)
		} else {
			up = createKey(side.y-1, side.x)
		}
		_, upExists := sides[up]
		if upExists {
			side.up = up
		} else {
			side.up = me
		}

		if side.y+1 > maxY[side.x] {
			down = createKey(minY[side.x], side.x)
		} else {
			down = createKey(side.y+1, side.x)
		}
		_, downExists := sides[down]
		if downExists {
			side.down = down
		} else {
			side.down = me
		}

		if side.x-1 < minX[side.y] {
			left = createKey(side.y, maxX[side.y])
		} else {
			left = createKey(side.y, side.x-1)
		}
		_, leftExists := sides[left]
		if leftExists {
			side.left = left
		} else {
			side.left = me
		}

		if side.x+1 > maxX[side.y] {
			right = createKey(side.y, minX[side.y])
		} else {
			right = createKey(side.y, side.x+1)
		}
		_, rightExists := sides[right]
		if rightExists {
			side.right = right
		} else {
			side.right = me
		}

		sides[me] = side
	}

	for _, side := range sides {
		fmt.Println(side.y, side.x, sides[side.up].y, sides[side.up].x, sides[side.down].y, sides[side.down].x, sides[side.left].y, sides[side.left].x, sides[side.right].y, sides[side.right].x)

	}

	// Starting position
	y := 0
	x := 0

	// Starting side
	side := sides[createKey(y, minX[x])]

	// Starting direction
	dir := "right"

	re := regexp.MustCompile("[0-9]+|R|L")
	path := re.FindAllString(lines[len(lines)-1], -1)
	fmt.Println(path)
	for i := 0; i < len(path); i++ {
		n, err := strconv.Atoi(path[i])
		if err == nil {
			for j := 0; j < n; j++ {
				fmt.Println(y+side.y*a, x+side.x*a)
				if dir == "up" {
					newY := y - 1

					if newY < 0 {
						newY = a - 1

						if !sides[side.up].walls[createKey(newY, x)] {
							y = newY
							side = sides[side.up]
						} else {
							break
						}

					} else if !side.walls[createKey(newY, x)] {
						y = newY
					} else {
						break
					}

				} else if dir == "down" {
					newY := y + 1

					if newY >= a {
						newY = 0

						if !sides[side.down].walls[createKey(newY, x)] {
							y = newY
							side = sides[side.down]
						} else {
							break
						}

					} else if !side.walls[createKey(newY, x)] {
						y = newY
					} else {
						break
					}

				} else if dir == "left" {
					newX := x - 1
					if newX < 0 {
						newX = a - 1

						if !sides[side.left].walls[createKey(y, newX)] {
							x = newX
							side = sides[side.left]
						} else {
							break
						}
					} else if !side.walls[createKey(y, newX)] {
						x = newX
					} else {
						break
					}

				} else if dir == "right" {
					newX := x + 1
					if newX >= a {
						newX = 0

						if !sides[side.right].walls[createKey(y, newX)] {
							x = newX
							side = sides[side.right]
						} else {
							break
						}
					} else if !side.walls[createKey(y, newX)] {
						x = newX
					} else {
						break
					}

				} else {
					panic("Wrong direction!")
				}
			}
		} else {
			if path[i] == "R" {
				if dir == "up" {
					dir = "right"
				} else if dir == "right" {
					dir = "down"
				} else if dir == "down" {
					dir = "left"
				} else if dir == "left" {
					dir = "up"
				} else {
					panic("Wrong direction set!")
				}
			} else if path[i] == "L" {
				if dir == "up" {
					dir = "left"
				} else if dir == "left" {
					dir = "down"
				} else if dir == "down" {
					dir = "right"
				} else if dir == "right" {
					dir = "up"
				} else {
					panic("Wrong direction set!")
				}
			} else {
				panic("Wrong rotation!")
			}
		}
	}

	dirValue := 0
	if dir == "down" {
		dirValue = 1
	} else if dir == "left" {
		dirValue = 2
	} else if dir == "up" {
		dirValue = 3
	}

	fmt.Println(y+side.y*a+1, x+side.x*a+1, dirValue)

	return 1000*(y+side.y*a+1) + 4*(x+side.x*a+1) + dirValue
}

/*
func mapSidesSample(sides map[string]Side) {
	key := createKey(0, 2)
	A := sides[key]
	A.up = createKey(1, 0)
	A.down = createKey(1, 2)
	A.left = createKey(1, 1)
	A.right = createKey(2, 3)

	A.moveDown = 0
	A.leftRot = 2
	A.rightRot = 0
	sides[key] = A

	key = createKey(1, 0)
	B := sides[key]
	B.up = createKey(0, 2)
	B.down = createKey(2, 2)
	B.left = createKey(2, 3)
	B.right = createKey(1, 1)
	B.moveUp = 0
	B.moveDown = 1
	B.leftRot = 0
	B.rightRot = 2
	sides[key] = B

	key = createKey(1, 1)
	C := sides[key]
	C.up = createKey(0, 2)
	C.down = createKey(2, 2)
	C.left = createKey(1, 0)
	C.right = createKey(1, 2)
	C.moveUp = 0
	C.moveDown = 0
	C.leftRot = 3
	C.rightRot = 3
	sides[key] = C

	key = createKey(1, 2)
	D := sides[key]
	D.up = createKey(0, 2)
	D.down = createKey(2, 2)
	D.left = createKey(1, 1)
	D.right = createKey(2, 3)
	C.moveUp = 0
	C.moveDown = 0
	C.leftRot = 3
	C.rightRot = 3
	sides[key] = D

	key = createKey(2, 1)
	E := sides[key]
	E.up = createKey(1, 2)
	E.down = createKey(1, 0)
	E.left = createKey(1, 1)
	E.right = createKey(2, 3)
	E.moveUp = 0
	E.moveDown = 1
	E.leftRot = 0
	E.rightRot = 2
	sides[key] = E

	key = createKey(2, 3)
	F := sides[key]
	F.up = createKey(1, 2)
	F.down = createKey(1, 0)
	F.left = createKey(2, 2)
	F.right = createKey(0, 2)
	F.moveUp = 0
	F.moveDown = 0
	F.leftRot = 3
	F.rightRot = 3
	sides[key] = F
}

*/

func mapSides(sides map[string]Side) {
	key := createKey(0, 1)
	A := sides[key]
	A.up = createKey(3, 0)
	A.down = createKey(1, 1)
	A.left = createKey(2, 0)
	A.right = createKey(0, 2)

	A.moveUp = func(y, x, a int) (int, int, string) {
		return x, 0, "right"
	}

	A.moveDown = func(y, x, a int) (int, int, string) {
		return 0, x, "down"
	}

	A.moveLeft = func(y, x, a int) (int, int, string) {
		return a - y - 1, 0, "right"
	}

	A.moveRight = func(y, x, a int) (int, int, string) {
		return y, 0, "right"
	}

	sides[key] = A

	key = createKey(0, 2)
	B := sides[key]
	B.up = createKey(3, 0)
	B.down = createKey(1, 1)
	B.left = createKey(0, 1)
	B.right = createKey(2, 1)

	B.moveUp = func(y, x, a int) (int, int, string) {
		return a - 1, x, "up"
	}

	B.moveDown = func(y, x, a int) (int, int, string) {
		return x, a - 1, "left"
	}

	B.moveLeft = func(y, x, a int) (int, int, string) {
		return y, a - 1, "left"
	}

	B.moveRight = func(y, x, a int) (int, int, string) {
		return a - y - 1, a - 1, "left"
	}

	sides[key] = B

	key = createKey(1, 1)
	C := sides[key]
	C.up = createKey(0, 1)
	C.down = createKey(2, 1)
	C.left = createKey(2, 0)
	C.right = createKey(0, 2)

	C.moveUp = func(y, x, a int) (int, int, string) {
		return a - 1, x, "up"
	}
	C.moveDown = func(y, x, a int) (int, int, string) {
		return 0, x, "down"
	}
	C.moveLeft = func(y, x, a int) (int, int, string) {
		return 0, y, "down"
	}
	C.moveRight = func(y, x, a int) (int, int, string) {
		return a - 1, y, "up"
	}
	sides[key] = C

	key = createKey(2, 0)
	D := sides[key]
	D.up = createKey(1, 1)
	D.down = createKey(3, 0)
	D.left = createKey(0, 1)
	D.right = createKey(2, 1)

	D.moveUp = func(y, x, a int) (int, int, string) {
		return x, 0, "right"
	}
	D.moveDown = func(y, x, a int) (int, int, string) {
		return 0, x, "down"
	}
	D.moveLeft = func(y, x, a int) (int, int, string) {
		return a - y - 1, 0, "right"
	}
	D.moveRight = func(y, x, a int) (int, int, string) {
		return y, 0, "right"
	}
	sides[key] = D

	key = createKey(2, 1)
	E := sides[key]
	E.up = createKey(1, 1)
	E.down = createKey(3, 0)
	E.left = createKey(2, 0)
	E.right = createKey(0, 2)
	E.moveUp = func(y, x, a int) (int, int, string) {
		return a - 1, x, "up"
	}
	E.moveDown = func(y, x, a int) (int, int, string) {
		return x, a - 1, "left"
	}
	E.moveLeft = func(y, x, a int) (int, int, string) {
		return y, a - 1, "left"
	}
	E.moveRight = func(y, x, a int) (int, int, string) {
		return a - y - 1, a - 1, "left"
	}
	sides[key] = E

	key = createKey(3, 0)
	F := sides[key]
	F.up = createKey(2, 0)
	F.down = createKey(0, 2)
	F.left = createKey(0, 1)
	F.right = createKey(2, 1)
	F.moveUp = func(y, x, a int) (int, int, string) {
		return a - 1, x, "up"
	}
	F.moveDown = func(y, x, a int) (int, int, string) {
		return 0, x, "down"
	}
	F.moveLeft = func(y, x, a int) (int, int, string) {
		return 0, y, "down"
	}
	F.moveRight = func(y, x, a int) (int, int, string) {
		return a - 1, y, "up"
	}
	sides[key] = F
}

func part2(input string) int {
	surface := 0
	for _, c := range input {
		if c == '.' || c == '#' {
			surface += 1
		}
	}
	a := int(math.Sqrt(float64(surface) / 6))

	sides := make(map[string]Side)

	maxY := make(map[int]int)
	maxX := make(map[int]int)

	minY := make(map[int]int)
	minX := make(map[int]int)

	lines := strings.Split(input, "\n")
	for y := 0; y < len(lines)-2; y++ {
		x := 0
		line := lines[y]
		for line[x] == ' ' {
			x += a
		}

		for x < len(line) {
			key := createKey(y/a, x/a)

			side, sideExists := sides[key]
			if !sideExists {
				sY := y / a
				sX := x / a

				side = Side{}
				side.y = sY
				side.x = sX
				side.walls = make(map[string]bool)
				sides[key] = side

				if v, e := maxX[sY]; !e {
					maxX[sY] = sX
				} else if sX > v {
					maxX[sY] = sX
				}

				if v, e := minX[sY]; !e {
					minX[sY] = sX
				} else if sX < v {
					minX[sY] = sX
				}

				if v, e := maxY[sX]; !e {
					maxY[sX] = sY
				} else if sY > v {
					maxY[sX] = sY
				}

				if v, e := minY[sX]; !e {
					minY[sX] = sY
				} else if sY < v {
					minY[sX] = sY
				}
			}

			if line[x] == '#' {
				side.walls[createKey(y%a, x%a)] = true
			}
			x++
		}
	}

	mapSides(sides)
	// mapSidesSample(sides)

	// for _, side := range sides {
	// 	fmt.Println(
	// 		side.y, side.x,
	// 		"up", sides[side.up].y, sides[side.up].x,
	// 		"down", sides[side.down].y, sides[side.down].x,
	// 		"left", sides[side.left].y, sides[side.left].x,
	// 		"right", sides[side.right].y, sides[side.right].x,
	// 		"moveUp", side.moveUp,
	// 		"moveDown", side.moveDown)

	// }

	// Starting position
	y := 0
	x := 0

	// Starting side
	side := sides[createKey(y, minX[x])]
	// side := sides[createKey(3, 0)]

	fmt.Println("start", y+side.y*a+1, x+side.x*a+1)

	// Starting direction
	dir := "right"

	re := regexp.MustCompile("[0-9]+|R|L")
	path := re.FindAllString(lines[len(lines)-1], -1)
	// path := re.FindAllString("200R200", -1)
	// fmt.Println(path)
	for i := 0; i < len(path); i++ {
		n, err := strconv.Atoi(path[i])
		if err == nil {
			for j := 0; j < n; j++ {
				// fmt.Println(y+side.y*a, x+side.x*a)
				if dir == "up" {
					newY := y - 1

					if newY < 0 {
						newY, newX, newDir := side.moveUp(y, x, a)
						if !sides[side.up].walls[createKey(newY, newX)] {
							y = newY
							x = newX
							dir = newDir
							side = sides[side.up]
						} else {
							break
						}

					} else if !side.walls[createKey(newY, x)] {
						y = newY
					} else {
						break
					}

				} else if dir == "down" {
					newY := y + 1

					if newY >= a {
						newY, newX, newDir := side.moveDown(y, x, a)

						if !sides[side.down].walls[createKey(newY, newX)] {
							y = newY
							x = newX
							dir = newDir
							side = sides[side.down]
						} else {
							break
						}

					} else if !side.walls[createKey(newY, x)] {
						y = newY
					} else {
						break
					}

				} else if dir == "left" {
					newX := x - 1
					if newX < 0 {
						newY, newX, newDir := side.moveLeft(y, x, a)

						if !sides[side.left].walls[createKey(newY, newX)] {
							y = newY
							x = newX
							dir = newDir
							side = sides[side.left]
						} else {
							break
						}
					} else if !side.walls[createKey(y, newX)] {
						x = newX
					} else {
						break
					}

				} else if dir == "right" {
					newX := x + 1
					if newX >= a {
						newY, newX, newDir := side.moveRight(y, x, a)
						if !sides[side.right].walls[createKey(newY, newX)] {
							y = newY
							x = newX
							dir = newDir
							side = sides[side.right]
						} else {
							break
						}
					} else if !side.walls[createKey(y, newX)] {
						x = newX
					} else {
						break
					}

				} else {
					panic("Wrong direction!")
				}
			}
		} else {
			if path[i] == "R" {
				if dir == "up" {
					dir = "right"
				} else if dir == "right" {
					dir = "down"
				} else if dir == "down" {
					dir = "left"
				} else if dir == "left" {
					dir = "up"
				} else {
					panic("Wrong direction set!")
				}
			} else if path[i] == "L" {
				if dir == "up" {
					dir = "left"
				} else if dir == "left" {
					dir = "down"
				} else if dir == "down" {
					dir = "right"
				} else if dir == "right" {
					dir = "up"
				} else {
					panic("Wrong direction set!")
				}
			} else {
				panic("Wrong rotation!")
			}
		}
	}

	dirValue := 0
	if dir == "down" {
		dirValue = 1
	} else if dir == "left" {
		dirValue = 2
	} else if dir == "up" {
		dirValue = 3
	}

	fmt.Println("end", y+side.y*a+1, x+side.x*a+1, dirValue)

	// 40456 TOO LOW
	// 67282 TOO LOW
	return 1000*(y+side.y*a+1) + 4*(x+side.x*a+1) + dirValue
}

func main() {

	inputPtr := flag.Bool("input", false, "sample or input")

	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")

	flag.Parse()

	var inputText string
	if *inputPtr {
		inputText = input
		fmt.Println("Running part", part, "on input.txt.")
	} else {
		inputText = sample
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
