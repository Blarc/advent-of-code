package main

import (
	"flag"
	"fmt"
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

func compareMirrorVertical(a [][]string, verticalIndex, dist int) bool {
	//fmt.Println("verticalIndex", verticalIndex, "verticalDist", dist)
	for y := 0; y < len(a); y++ {
		for x := 0; x < dist; x++ {
			if a[y][verticalIndex-x] != a[y][verticalIndex+1+x] {
				return false
			}
		}
	}
	return true
}

func compareMirrorHorizontal(a [][]string, horizontalIndex, dist int) bool {
	//fmt.Println("horizontalIndex", horizontalIndex, "horizontalDist", dist)
	for x := 0; x < len(a[0]); x++ {
		for y := 0; y < dist; y++ {
			if a[horizontalIndex-y][x] != a[horizontalIndex+1+y][x] {
				return false
			}
		}
	}
	return true
}

func compareMirrorVerticalSmudge(a [][]string, verticalIndex, dist int) bool {
	//fmt.Println("verticalIndex", verticalIndex, "verticalDist", dist)
	smudge := false
	for y := 0; y < len(a); y++ {
		for x := 0; x < dist; x++ {
			if a[y][verticalIndex-x] != a[y][verticalIndex+1+x] {
				if !smudge {
					smudge = true
				} else {
					return false
				}
			}
		}
	}
	return smudge
}

func compareMirrorHorizontalSmudge(a [][]string, horizontalIndex, dist int) bool {
	//fmt.Println("horizontalIndex", horizontalIndex, "horizontalDist", dist)
	smudge := false
	for x := 0; x < len(a[0]); x++ {
		for y := 0; y < dist; y++ {
			if a[horizontalIndex-y][x] != a[horizontalIndex+1+y][x] {
				if !smudge {
					smudge = true
				} else {
					return false
				}
			}
		}
	}
	return smudge
}

func compareMirror(mirror [][]string) int {
	verticalDist := len(mirror[0]) / 2
	verticalIndex := len(mirror[0])/2 - 1

	maxVerticalIndex := 0
	for i := 0; i < verticalDist; i++ {
		vertical := compareMirrorVertical(mirror, verticalIndex-i, verticalDist-i)
		if vertical {
			if verticalIndex+1-i > maxVerticalIndex {
				maxVerticalIndex = verticalIndex + 1 - i
			}
			break
		}

		vertical = compareMirrorVertical(mirror, verticalIndex+1+i, verticalDist-i)
		if vertical {
			if verticalIndex+2+i > maxVerticalIndex {
				maxVerticalIndex = verticalIndex + 2 + i
			}
			break
		}
	}

	//fmt.Println("Max vertical:", maxVerticalIndex)

	horizontalDist := len(mirror) / 2
	horizontalIndex := len(mirror)/2 - 1

	maxHorizontalIndex := 0
	for i := 0; i < horizontalDist; i++ {
		horizontal := compareMirrorHorizontal(mirror, horizontalIndex-i, horizontalDist-i)
		if horizontal {
			if horizontalIndex+1-i > maxHorizontalIndex {
				maxHorizontalIndex = horizontalIndex + 1 - i
			}
			break
		}

		horizontal = compareMirrorHorizontal(mirror, horizontalIndex+1+i, horizontalDist-i)
		if horizontal {
			if horizontalIndex+2+i > maxHorizontalIndex {
				maxHorizontalIndex = horizontalIndex + 2 + i
			}
			break
		}
	}

	//fmt.Println("Max horizontal:", maxHorizontalIndex)

	maxIndex := 0
	if maxHorizontalIndex > maxVerticalIndex {
		maxIndex = maxHorizontalIndex * 100
	} else {
		maxIndex = maxVerticalIndex
	}
	if maxIndex == 0 {
		panic("hehe")
	}
	return maxIndex

}

func compareMirrorSmudge(mirror [][]string) int {
	verticalDist := len(mirror[0]) / 2
	verticalIndex := len(mirror[0])/2 - 1

	maxVerticalIndex := 0
	for i := 0; i < verticalDist; i++ {
		vertical := compareMirrorVerticalSmudge(mirror, verticalIndex-i, verticalDist-i)
		if vertical {
			if verticalIndex+1-i > maxVerticalIndex {
				maxVerticalIndex = verticalIndex + 1 - i
			}
			break
		}

		vertical = compareMirrorVerticalSmudge(mirror, verticalIndex+1+i, verticalDist-i)
		if vertical {
			if verticalIndex+2+i > maxVerticalIndex {
				maxVerticalIndex = verticalIndex + 2 + i
			}
			break
		}
	}

	//fmt.Println("Max vertical:", maxVerticalIndex)

	horizontalDist := len(mirror) / 2
	horizontalIndex := len(mirror)/2 - 1

	maxHorizontalIndex := 0
	for i := 0; i < horizontalDist; i++ {
		horizontal := compareMirrorHorizontalSmudge(mirror, horizontalIndex-i, horizontalDist-i)
		if horizontal {
			if horizontalIndex+1-i > maxHorizontalIndex {
				maxHorizontalIndex = horizontalIndex + 1 - i
			}
			break
		}

		horizontal = compareMirrorHorizontalSmudge(mirror, horizontalIndex+1+i, horizontalDist-i)
		if horizontal {
			if horizontalIndex+2+i > maxHorizontalIndex {
				maxHorizontalIndex = horizontalIndex + 2 + i
			}
			break
		}
	}

	//fmt.Println("Max horizontal:", maxHorizontalIndex)

	maxIndex := 0
	if maxHorizontalIndex > maxVerticalIndex {
		maxIndex = maxHorizontalIndex * 100
	} else {
		maxIndex = maxVerticalIndex
	}
	if maxIndex == 0 {
		panic("hehe")
	}
	return maxIndex

}

func part1(input string) int {
	lines := strings.Split(input, "\n")

	mirror := make([][]string, 0)
	sum := 0
	for _, l := range lines {
		if l == "" {
			//for k := 0; k < len(mirror); k++ {
			//	fmt.Println(mirror[k])
			//}

			sum += compareMirror(mirror)

			mirror = make([][]string, 0)
		} else {
			mirror = append(mirror, strings.Split(l, ""))
		}

	}

	// 1708
	// 30155 too high
	return sum
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	mirror := make([][]string, 0)
	sum := 0
	for _, l := range lines {
		if l == "" {
			//for k := 0; k < len(mirror); k++ {
			//	fmt.Println(mirror[k])
			//}

			sum += compareMirrorSmudge(mirror)

			mirror = make([][]string, 0)
		} else {
			mirror = append(mirror, strings.Split(l, ""))
		}

	}

	// 1708
	// 30155 too high
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
		fmt.Println("Result:", part1(inputText))
	} else {
		fmt.Println("Result:", part2(inputText))
	}
	fmt.Println("Time:", fmt.Sprintf("%d ms", time.Since(start).Milliseconds()))
}
