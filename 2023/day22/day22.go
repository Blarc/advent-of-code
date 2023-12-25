package main

import (
	"flag"
	"fmt"
	"slices"
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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func coloredString(text string, color int) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, text)
}

func mustAtoi(s string) int {
	number, _ := strconv.Atoi(s)
	return number
}

type Brick struct {
	id                  string
	fromX, fromY, fromZ int
	toX, toY, toZ       int
	holds               map[string]*Brick
	holdedBy            map[string]*Brick
}

func part1(input string) int {

	maxX, maxY := -1, -1

	var bricks []*Brick
	lines := strings.Split(input, "\n")
	for i, l := range lines {
		l = strings.Replace(l, "~", ",", 1)
		coordinates := strings.Split(l, ",")
		brick := &Brick{
			id:       string(rune('A' + i)),
			fromX:    mustAtoi(coordinates[0]),
			fromY:    mustAtoi(coordinates[1]),
			fromZ:    mustAtoi(coordinates[2]),
			toX:      mustAtoi(coordinates[3]),
			toY:      mustAtoi(coordinates[4]),
			toZ:      mustAtoi(coordinates[5]),
			holds:    make(map[string]*Brick),
			holdedBy: make(map[string]*Brick),
		}

		if brick.fromX > brick.toX || brick.fromY > brick.toY || brick.fromZ > brick.toZ {
			panic("not something I expected")
		}

		if brick.toX > maxX {
			maxX = brick.toX
		}

		if brick.toY > maxY {
			maxY = brick.toY
		}

		bricks = append(bricks, brick)
	}

	// From lowest to highest Z
	slices.SortFunc(bricks, func(a, b *Brick) int {
		return a.fromZ - b.fromZ
	})

	maxX += 1
	maxY += 1

	// Prepare empty grid
	floor := &Brick{
		id:       "X",
		fromX:    0,
		fromY:    0,
		fromZ:    0,
		toX:      maxX,
		toY:      maxY,
		toZ:      0,
		holds:    make(map[string]*Brick),
		holdedBy: make(map[string]*Brick),
	}
	m := make([][]*Brick, maxX)
	for x := 0; x < maxX; x++ {
		m[x] = make([]*Brick, maxY)
		for y := 0; y < maxY; y++ {
			m[x][y] = floor
		}
	}

	for _, brick := range bricks {
		//fmt.Printf("%+v\n", brick)
		var holders []*Brick
		for x := brick.fromX; x <= brick.toX; x++ {
			for y := brick.fromY; y <= brick.toY; y++ {
				current := m[x][y]
				if len(holders) == 0 || holders[0].toZ < current.toZ {
					holders = []*Brick{current}
				} else if holders[0].toZ == current.toZ {
					holders = append(holders, current)
				}
			}
		}

		brick.toZ = holders[0].toZ + (brick.toZ - brick.fromZ) + 1
		brick.fromZ = holders[0].toZ + 1
		for x := brick.fromX; x <= brick.toX; x++ {
			for y := brick.fromY; y <= brick.toY; y++ {
				m[x][y] = brick
			}
		}

		for _, holder := range holders {
			brick.holdedBy[holder.id] = holder
			holder.holds[brick.id] = brick
		}

		//for x := 0; x < maxX; x++ {
		//	for y := 0; y < maxY; y++ {
		//		if m[x][y] != nil {
		//			fmt.Printf("%2d ", m[x][y].toZ)
		//		}
		//	}
		//	fmt.Println()
		//}
		//fmt.Println("#########")
	}

	//for _, brick := range bricks {
	//	fmt.Printf("%s %d-%d holds: ", brick.id, brick.fromZ, brick.toZ)
	//	for _, h := range brick.holds {
	//		fmt.Printf("%s ", h.id)
	//	}
	//	fmt.Print("holded by: ")
	//	for _, h := range brick.holdedBy {
	//		fmt.Printf("%s ", h.id)
	//	}
	//	fmt.Print("can be disintegrated: ", brick.areHoldedAlsoHoldedByOther())
	//	fmt.Println()
	//}

	sum := 0
	for _, brick := range bricks {
		if brick.areHoldedAlsoHoldedByOther() {
			sum += 1
		}
	}

	// 644 too high
	// 643 ??
	return sum
}

func (b *Brick) areHoldedAlsoHoldedByOther() bool {
	for _, brick := range b.holds {
		if len(brick.holdedBy) < 2 {
			return false
		}
	}
	return true
}

func (b *Brick) howManyWouldFall(deleted map[string]bool) int {
	sum := 0
	deleted[b.id] = true
	for _, v := range b.holds {
		if v.areHoldersDeleted(deleted) {
			sum += v.howManyWouldFall(deleted) + 1
		}
	}
	return sum
}

func (b *Brick) areHoldersDeleted(deleted map[string]bool) bool {
	for _, v := range b.holdedBy {
		if !deleted[v.id] {
			return false
		}
	}
	return true
}

func part2(input string) int {
	maxX, maxY := -1, -1

	var bricks []*Brick
	lines := strings.Split(input, "\n")
	for i, l := range lines {
		l = strings.Replace(l, "~", ",", 1)
		coordinates := strings.Split(l, ",")
		brick := &Brick{
			id:       string(rune('A' + i)),
			fromX:    mustAtoi(coordinates[0]),
			fromY:    mustAtoi(coordinates[1]),
			fromZ:    mustAtoi(coordinates[2]),
			toX:      mustAtoi(coordinates[3]),
			toY:      mustAtoi(coordinates[4]),
			toZ:      mustAtoi(coordinates[5]),
			holds:    make(map[string]*Brick),
			holdedBy: make(map[string]*Brick),
		}

		if brick.fromX > brick.toX || brick.fromY > brick.toY || brick.fromZ > brick.toZ {
			panic("not something I expected")
		}

		if brick.toX > maxX {
			maxX = brick.toX
		}

		if brick.toY > maxY {
			maxY = brick.toY
		}

		bricks = append(bricks, brick)
	}

	// From lowest to highest Z
	slices.SortFunc(bricks, func(a, b *Brick) int {
		return a.fromZ - b.fromZ
	})

	maxX += 1
	maxY += 1

	// Prepare empty grid
	floor := &Brick{
		id:       "X",
		fromX:    0,
		fromY:    0,
		fromZ:    0,
		toX:      maxX,
		toY:      maxY,
		toZ:      0,
		holds:    make(map[string]*Brick),
		holdedBy: make(map[string]*Brick),
	}
	m := make([][]*Brick, maxX)
	for x := 0; x < maxX; x++ {
		m[x] = make([]*Brick, maxY)
		for y := 0; y < maxY; y++ {
			m[x][y] = floor
		}
	}

	for _, brick := range bricks {
		//fmt.Printf("%+v\n", brick)
		var holders []*Brick
		for x := brick.fromX; x <= brick.toX; x++ {
			for y := brick.fromY; y <= brick.toY; y++ {
				current := m[x][y]
				if len(holders) == 0 || holders[0].toZ < current.toZ {
					holders = []*Brick{current}
				} else if holders[0].toZ == current.toZ {
					holders = append(holders, current)
				}
			}
		}

		brick.toZ = holders[0].toZ + (brick.toZ - brick.fromZ) + 1
		brick.fromZ = holders[0].toZ + 1
		for x := brick.fromX; x <= brick.toX; x++ {
			for y := brick.fromY; y <= brick.toY; y++ {
				m[x][y] = brick
			}
		}

		for _, holder := range holders {
			brick.holdedBy[holder.id] = holder
			holder.holds[brick.id] = brick
		}

		//for x := 0; x < maxX; x++ {
		//	for y := 0; y < maxY; y++ {
		//		if m[x][y] != nil {
		//			fmt.Printf("%2d ", m[x][y].toZ)
		//		}
		//	}
		//	fmt.Println()
		//}
		//fmt.Println("#########")
	}

	//for _, brick := range bricks {
	//	fmt.Printf("%s %d-%d holds: ", brick.id, brick.fromZ, brick.toZ)
	//	for _, h := range brick.holds {
	//		fmt.Printf("%s ", h.id)
	//	}
	//	fmt.Print("holded by: ")
	//	for _, h := range brick.holdedBy {
	//		fmt.Printf("%s ", h.id)
	//	}
	//	fmt.Print("can be disintegrated: ", brick.areHoldedAlsoHoldedByOther())
	//	fmt.Println()
	//}

	sum := 0
	for _, brick := range bricks {
		if !brick.areHoldedAlsoHoldedByOther() {
			sum += brick.howManyWouldFall(make(map[string]bool))
		}
	}

	// 644 too high
	// 643 ??
	return sum
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
		fmt.Println("Result:", part1(strings.TrimSpace(inputText)))
	} else {
		fmt.Println("Result:", part2(strings.TrimSpace(inputText)))
	}
	fmt.Println("Time:", fmt.Sprintf("%d ms", time.Since(start).Milliseconds()))
}
