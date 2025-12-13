package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"

	utils "github.com/Blarc/advent-of-code"
)

import (
	_ "embed"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

type Present struct {
	id    int
	shape [][]bool
}

func (p *Present) rotate90() Present {
	newShape := make([][]bool, 3)
	for y := 0; y < 3; y++ {
		newShape[y] = make([]bool, 3)
		for x := 0; x < 3; x++ {
			newShape[y][x] = p.shape[2-x][y]
		}
	}
	return Present{p.id, newShape}
}

func (p *Present) String() string {
	s := ""
	for _, row := range p.shape {
		for _, v := range row {
			if v {
				s += "#"
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	return s
}

type Region struct {
	space [][]bool
	todo  map[int]int
}

func (r *Region) String() string {
	s := ""
	for _, row := range r.space {
		for _, v := range row {
			if v {
				s += "#"
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	return s
}

func (r *Region) fitsAt(y, x int, present *Present) *Region {
	if y+3 > len(r.space) || x+3 > len(r.space[0]) {
		return nil
	}

	newSpace := make([][]bool, len(r.space))
	for i := range r.space {
		newSpace[i] = make([]bool, len(r.space[i]))
		copy(newSpace[i], r.space[i])
	}

	for j := 0; j < 3; j++ {
		for i := 0; i < 3; i++ {
			if present.shape[j][i] && r.space[y+j][x+i] {
				return nil
			}
			if present.shape[j][i] {
				newSpace[y+j][x+i] = present.shape[j][i]
			}
		}
	}

	newTodo := make(map[int]int, len(r.todo))
	for k, v := range r.todo {
		newTodo[k] = v
	}
	newTodo[present.id] -= 1
	if newTodo[present.id] == 0 {
		delete(newTodo, present.id)
	}

	return &Region{
		space: newSpace,
		todo:  newTodo,
	}
}

func (r *Region) fits(present *Present) *Region {
	for y := 0; y < len(r.space); y++ {
		for x := 0; x < len(r.space[0]); x++ {
			newPresent := present.rotate90()
			for rotation := 0; rotation < 4; rotation++ {
				newRegion := r.fitsAt(y, x, &newPresent)
				if newRegion != nil {
					return newRegion
				}
				newPresent = newPresent.rotate90()
			}
		}
	}
	return nil
}

func canFitPresents(region *Region, presents []Present) *Region {
	if len(region.todo) == 0 {
		return region
	}

	for k, _ := range region.todo {
		newRegion := region.fits(&presents[k])
		if newRegion != nil {
			return canFitPresents(newRegion, presents)
		}
	}

	return nil
}

func part1(input string) int {
	result := 0
	const numPresents = 5
	lines := strings.Split(input, "\n")

	presents := make([]Present, numPresents+1)
	for i := 0; i < len(lines); i++ {

		if i <= numPresents*5 {
			id, _ := strconv.Atoi(strings.TrimSuffix(lines[i], ":"))
			//fmt.Println(id)
			i++

			shape := make([][]bool, 3)
			for y := 0; y < 3; y++ {
				shape[y] = make([]bool, 3)
				for x := 0; x < 3; x++ {
					//fmt.Printf("%c ", lines[i][x])
					shape[y][x] = lines[i][x] == '#'
				}
				//fmt.Println()
				i++
			}

			presents[id] = Present{id, shape}
		} else {
			s := strings.Split(lines[i], ":")
			ss := strings.Split(s[0], "x")
			maxX, _ := strconv.Atoi(ss[0])
			maxY, _ := strconv.Atoi(ss[1])

			space := make([][]bool, maxY)
			for y := 0; y < maxY; y++ {
				space[y] = make([]bool, maxX)
				for x := 0; x < maxX; x++ {
					space[y][x] = false
				}
			}

			todo := make(map[int]int)
			for j, v := range utils.ToIntSlice(strings.Fields(s[1])) {
				if v != 0 {
					todo[j] = v
				}
			}

			region := Region{
				space: space,
				todo:  todo,
			}

			numOfPresents := 0
			for _, v := range region.todo {
				numOfPresents += v
			}

			if numOfPresents*9 <= maxX*maxY {
				result++
			}

			//newRegion := region.fits(&presents[4])
			//fmt.Println(newRegion)
			//newRegion = newRegion.fits(&presents[4])
			//fmt.Println(newRegion)

			//resultRegion := canFitPresents(&region, presents)
			//fmt.Println(resultRegion)
			//if resultRegion != nil {
			//	result++
			//}
		}
	}

	// 407 is too low
	return result
}

func part2(input string) int {
	return 0
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
