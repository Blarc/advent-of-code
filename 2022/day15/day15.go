package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattan(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func part1(input string) int {

	// This works but it's wrong.

	scanned := make(map[int][]int)
	beacons := make(map[int]int)

	minX := math.MaxInt
	minY := math.MaxInt
	maxX := math.MinInt
	maxY := math.MinInt

	for _, l := range strings.Split(input, "\n") {

		// fmt.Println(l)

		re := regexp.MustCompile("-?[0-9]+")
		matches := re.FindAllString(l, -1)

		sensorX, _ := strconv.Atoi(matches[0])
		sensorY, _ := strconv.Atoi(matches[1])

		beaconX, _ := strconv.Atoi(matches[2])
		beaconY, _ := strconv.Atoi(matches[3])

		beacons[beaconY] += 1

		radius := manhattan(sensorX, sensorY, beaconX, beaconY)

		for y := -radius; y <= radius; y++ {

			cy := sensorY + y

			startX := sensorX - radius + abs(y)
			endX := sensorX + radius - abs(y)

			if _, exists := scanned[cy]; exists {
				if startX < scanned[cy][0] {
					scanned[cy][0] = startX
				}
				if scanned[cy][1] < endX {
					scanned[cy][1] = endX
				}
			} else {
				scanned[cy] = []int{startX, endX}
			}

			if startX < minX {
				minX = startX
			} else if endX > maxX {
				maxX = endX
			}

			if cy < minY {
				minY = cy
			} else if cy > maxY {
				maxY = cy
			}
		}
	}

	// fmt.Println(scanned)
	y := 2000000

	// fmt.Println(scanned[y][0])
	// fmt.Println(scanned[y][1])
	// fmt.Println(beacons[y])

	return scanned[y][1] - scanned[y][0]
}

func reduce(x int) int {
	// 4000000
	if x < 0 {
		return 0
	} else if x > 4000000 {
		return 4000000
	}
	return x
}

func mergeIntervals(start1, end1, start2, end2 int) []int {
	if start1 <= start2 && end2 <= end1 {
		return []int{start1, end1}
	} else if start2 <= start1 && end1 <= end2 {
		return []int{start2, end2}
	} else if start2 <= start1 && end2 <= end1 && start1 <= end2+1 {
		return []int{start2, end1}
	} else if start1 <= start2 && end1 <= end2 && end1+1 >= start2 {
		return []int{start1, end2}
	}

	return []int{}
}

func part2(input string) int {
	scanned := make(map[int][][]int)
	beacons := make(map[int]int)

	/*
		minX := 0
		minY := 0
		maxX := 20
		maxY := 20
	*/

	for _, l := range strings.Split(input, "\n") {

		// fmt.Println(l)

		re := regexp.MustCompile("-?[0-9]+")
		matches := re.FindAllString(l, -1)

		sensorX, _ := strconv.Atoi(matches[0])
		sensorY, _ := strconv.Atoi(matches[1])

		beaconX, _ := strconv.Atoi(matches[2])
		beaconY, _ := strconv.Atoi(matches[3])

		beacons[beaconY] += 1

		radius := manhattan(sensorX, sensorY, beaconX, beaconY)

		startY := reduce(sensorY - radius)
		endY := reduce(sensorY + radius)

		for y := startY; y <= endY; y++ {

			startX := reduce(sensorX - radius + abs(y-sensorY))
			endX := reduce(sensorX + radius - abs(y-sensorY))

			// fmt.Println("x:", startX, endX)

			if _, exists := scanned[y]; exists {

				updatedIntervalIndex := -1

				// addNewInterval := true
				numberOfIntervals := len(scanned[y])
				for i := 0; i < numberOfIntervals; i++ {

					mergedInterval := mergeIntervals(startX, endX, scanned[y][i][0], scanned[y][i][1])
					if len(mergedInterval) > 0 {
						if updatedIntervalIndex != -1 {
							updatedInterval := scanned[y][updatedIntervalIndex]
							mergedIntervalTmp := mergeIntervals(mergedInterval[0], mergedInterval[1], updatedInterval[0], updatedInterval[1])
							if len(mergedIntervalTmp) > 0 {
								mergedInterval = mergedIntervalTmp
								scanned[y] = append(scanned[y][:updatedIntervalIndex], scanned[y][updatedIntervalIndex+1:]...)
								i -= 1
								numberOfIntervals -= 1
							}
						}
						scanned[y][i][0] = mergedInterval[0]
						scanned[y][i][1] = mergedInterval[1]

						updatedIntervalIndex = i
					}

				}
				if updatedIntervalIndex == -1 {
					scanned[y] = append(scanned[y], []int{startX, endX})
				}

			} else {
				scanned[y] = append(scanned[y], []int{startX, endX})
			}
		}

		// fmt.Println(scanned[11])

	}

	result := 0
	for i := 0; i < len(scanned); i++ {
		if len(scanned[i]) > 1 {
			fmt.Println(i, scanned[i])
			// x * 4000000 + y
			if scanned[i][0][1] > scanned[i][1][1] {
				fmt.Println(scanned[i][1][1]+1, "*", 4000000, "+", i)
				result = (scanned[i][1][1]+1)*4000000 + i
			}
		}
	}

	return result
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

	if part == 1 {
		fmt.Println("Result:", part1(inputText))
	} else {
		fmt.Println("Result:", part2(inputText))
	}
}
