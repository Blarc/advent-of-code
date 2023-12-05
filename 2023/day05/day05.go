package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

type MapRange struct {
	SourceFrom int
	// Exclusive
	SourceTo        int
	DestinationFrom int
	// Exclusive
	DestinationTo int
}

type SeedRange struct {
	From int
	// Exclusive
	To int
}

func extractNumbersToSlice(line string) []int {
	numberRe := regexp.MustCompile("[0-9]+")
	matches := numberRe.FindAllString(line, -1)

	seeds := make([]int, 0)
	for _, m := range matches {
		intM, _ := strconv.Atoi(m)
		seeds = append(seeds, intM)
	}
	return seeds
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	seeds := extractNumbersToSlice(lines[0])

	fmt.Printf("%v\n", seeds)

	m := make(map[int]map[int]MapRange)
	order := 0

	for index := 2; index < len(lines); index++ {
		line := lines[index]
		if strings.HasSuffix(line, "map:") {
			name := strings.Split(line, " ")[0]
			m[order] = make(map[int]MapRange)
			index++
			println(name, order)
			mapLine := lines[index]
			for mapLine != "" {
				numbers := extractNumbersToSlice(mapLine)
				m[order][numbers[0]] = MapRange{
					SourceFrom:      numbers[1],
					SourceTo:        numbers[1] + numbers[2],
					DestinationFrom: numbers[0],
					DestinationTo:   numbers[0] + numbers[2],
				}
				fmt.Printf("%v\n", m[order][numbers[0]])

				index++
				mapLine = lines[index]
			}
		}
		order++
	}

	for orderIndex := 0; orderIndex < order; orderIndex++ {
		for i := range seeds {
			seed := seeds[i]
			for _, value := range m[orderIndex] {
				if value.SourceFrom <= seed && seed < value.SourceTo {
					seeds[i] = value.DestinationFrom + (seed - value.SourceFrom)
					break
				}
			}
			//if i == len(seeds)-1 {
			//	fmt.Printf("%v\n", seeds)
			//}
		}
	}

	return slices.Min(seeds)
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	seedsRange := extractNumbersToSlice(lines[0])

	var seeds []SeedRange
	for i := 0; i < len(seedsRange); i += 2 {
		seeds = append(seeds, SeedRange{seedsRange[i], seedsRange[i] + seedsRange[i+1]})
	}

	fmt.Printf("%v\n", seeds)

	m := make(map[int]map[int]MapRange)
	order := 0

	for index := 2; index < len(lines); index++ {
		line := lines[index]
		if strings.HasSuffix(line, "map:") {
			//name := strings.Split(line, " ")[0]
			m[order] = make(map[int]MapRange)
			index++
			//println(name, order)
			mapLine := lines[index]
			for mapLine != "" {
				numbers := extractNumbersToSlice(mapLine)
				m[order][numbers[0]] = MapRange{
					SourceFrom:      numbers[1],
					SourceTo:        numbers[1] + numbers[2],
					DestinationFrom: numbers[0],
					DestinationTo:   numbers[0] + numbers[2],
				}
				//fmt.Printf("%v\n", m[order][numbers[0]])

				index++
				mapLine = lines[index]
			}
		}
		order++
	}

	for orderIndex := 0; orderIndex < order; orderIndex++ {
		var newSeeds []SeedRange
		//fmt.Printf("Order: %d, Seeds: %v\n", orderIndex, seeds)

		for i := range seeds {
			seed := seeds[i]
			newSeeds = append(newSeeds, dorec(&seed, orderIndex, m)...)
		}
		seeds = newSeeds
	}

	return slices.MinFunc(seeds, func(a, b SeedRange) int {
		if a.From < b.From {
			return -1
		}
		if a.From > b.From {
			return 1
		}
		return 0
	}).From

}

func dorec(seed *SeedRange, orderIndex int, m map[int]map[int]MapRange) []SeedRange {
	for _, value := range m[orderIndex] {
		//fmt.Printf("# Seed: %v Source: %v\n", seed, value)
		if value.SourceFrom <= seed.From && seed.To <= value.SourceTo {
			newSeed := SeedRange{
				From: value.DestinationFrom + (seed.From - value.SourceFrom),
				To:   value.DestinationFrom + (seed.To - value.SourceFrom),
			}
			//fmt.Printf("## New seed: %v\n", newSeed)
			return []SeedRange{newSeed}
		}
		if value.SourceFrom <= seed.From && seed.From < value.SourceTo {
			var newSeeds []SeedRange

			firstSeed := SeedRange{
				From: value.DestinationFrom + (seed.From - value.SourceFrom),
				To:   value.DestinationTo,
			}
			//fmt.Printf("## First seed: %v\n", firstSeed)
			newSeeds = append(newSeeds, firstSeed)

			// Recursive call for the rest of the seed
			restSeed := SeedRange{
				From: value.SourceTo,
				To:   seed.To,
			}
			//fmt.Printf("## Second seed %v\n", restSeed)
			recSeeds := dorec(&restSeed, orderIndex, m)
			newSeeds = append(newSeeds, recSeeds...)
			return newSeeds
		}
	}
	//fmt.Println("### No match")
	return []SeedRange{*seed}
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
