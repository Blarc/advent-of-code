package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

/*
ore
clay
obsidian
geode
*/

type Blueprint struct {
	costs []int
}

func (b Blueprint) canMakeOre(ore int) bool {
	return b.costs[0] <= ore
}

func (b Blueprint) canMakeClay(ore int) bool {
	return b.costs[1] <= ore
}

func (b Blueprint) canMakeObsidian(ore, clay int) bool {
	return b.costs[2] <= ore && b.costs[3] <= clay
}

func (b Blueprint) canMakeGeode(ore, obsidian int) bool {
	return b.costs[4] <= ore && b.costs[5] <= obsidian
}

func (b Blueprint) oreCost() int {
	return b.costs[0]
}

func (b Blueprint) clayCost() int {
	return b.costs[1]
}

func (b Blueprint) obsidianCost() (int, int) {
	return b.costs[2], b.costs[3]
}

func (b Blueprint) geodeCost() (int, int) {
	return b.costs[4], b.costs[5]
}

func (b Blueprint) maxOreProd() int {
	max := -1
	for _, c := range []int{b.costs[0], b.costs[1], b.costs[2], b.costs[4]} {
		if c > max {
			max = c
		}
	}
	return max
}

func (b Blueprint) maxClayProd() int {
	return b.costs[3]
}

func (b Blueprint) maxObsidianProd() int {
	return b.costs[5]
}

func dfs(b Blueprint, oreR, clayR, obsidianR, geodeR, ore, clay, obsidian, geode, time int, memo map[string]int, path [][]int) (int, [][]int) {

	if time <= 0 {
		return geode, path
	}

	memoKey := fmt.Sprintf("%d,%d,%d,%d,%d,%d,%d,%d,%d", oreR, clayR, obsidianR, geodeR, ore, clay, obsidian, geode, time)
	if v, e := memo[memoKey]; e {
		return v, path
	}

	max := geode
	bestPath := path

	result, resultPath := dfs(b, oreR, clayR, obsidianR, geodeR, ore+oreR, clay+clayR, obsidian+obsidianR, geode+geodeR, time-1, memo,
		append(path, []int{oreR, clayR, obsidianR, geodeR, ore + oreR, clay + clayR, obsidian + obsidianR, geode + geodeR, time - 1}))
	if result > max {
		max = result
		bestPath = resultPath
	}

	if b.canMakeOre(ore) {
		result, resultPath := dfs(b, oreR+1, clayR, obsidianR, geodeR, ore+oreR-b.oreCost(), clay+clayR, obsidian+obsidianR, geode+geodeR, time-1, memo,
			append(path, []int{oreR + 1, clayR, obsidianR, geodeR, ore + oreR - b.oreCost(), clay + clayR, obsidian + obsidianR, geode + geodeR, time - 1}))
		if result > max {
			max = result
			bestPath = resultPath
		}
	}

	if b.canMakeClay(ore) {
		result, resultPath := dfs(b, oreR, clayR+1, obsidianR, geodeR, ore+oreR-b.clayCost(), clay+clayR, obsidian+obsidianR, geode+geodeR, time-1, memo,
			append(path, []int{oreR, clayR + 1, obsidianR, geodeR, ore + oreR - b.clayCost(), clay + clayR, obsidian + obsidianR, geode + geodeR, time - 1}))
		if result > max {
			max = result
			bestPath = resultPath
		}
	}

	if b.canMakeObsidian(ore, clay) {
		oreCost, clayCost := b.obsidianCost()
		result, resultPath := dfs(b, oreR, clayR, obsidianR+1, geodeR, ore+oreR-oreCost, clay+clayR-clayCost, obsidian+obsidianR, geode+geodeR, time-1, memo,
			append(path, []int{oreR, clayR, obsidianR + 1, geodeR, ore + oreR - oreCost, clay + clayR - clayCost, obsidian + obsidianR, geode + geodeR, time - 1}))
		if result > max {
			max = result
			bestPath = resultPath
		}
	}

	if b.canMakeGeode(ore, obsidian) {
		oreCost, obsidianCost := b.geodeCost()
		result, resultPath := dfs(b, oreR, clayR, obsidianR, geodeR+1, ore+oreR-oreCost, clay+clayR, obsidian+obsidianR-obsidianCost, geode+geodeR, time-1, memo,
			append(path, []int{oreR, clayR, obsidianR, geodeR + 1, ore + oreR - oreCost, clay + clayR, obsidian + obsidianR - obsidianCost, geode + geodeR, time - 1}))
		if result > max {
			max = result
			bestPath = resultPath
		}
	}

	memo[memoKey] = max

	return max, bestPath

}

func dfsPart1(b Blueprint, oreR, clayR, obsidianR, geodeR, ore, clay, obsidian, geode, time int, memo map[string]int, currentMax int) int {

	if time <= 0 {
		return geode
	}

	memoKey := fmt.Sprintf("%d,%d,%d,%d,%d,%d,%d,%d,%d", oreR, clayR, obsidianR, geodeR, ore, clay, obsidian, geode, time)
	if v, e := memo[memoKey]; e {
		return v
	}

	potentialMax := geode + geodeR*time + time*(time+1)/2
	if currentMax >= potentialMax {
		return currentMax
	}

	if oreR > b.maxOreProd() || clayR > b.maxClayProd() || obsidianR > b.maxObsidianProd() {
		return 0
	}

	max := geode

	if b.canMakeGeode(ore, obsidian) {
		oreCost, obsidianCost := b.geodeCost()
		result := dfs2(b, oreR, clayR, obsidianR, geodeR+1, ore+oreR-oreCost, clay+clayR, obsidian+obsidianR-obsidianCost, geode+geodeR, time-1, memo, max)
		if result > max {
			max = result
		}
	} else if b.canMakeObsidian(ore, clay) {
		oreCost, clayCost := b.obsidianCost()
		result := dfs2(b, oreR, clayR, obsidianR+1, geodeR, ore+oreR-oreCost, clay+clayR-clayCost, obsidian+obsidianR, geode+geodeR, time-1, memo, max)
		if result > max {
			max = result
		}
	} else {
		result := dfs2(b, oreR, clayR, obsidianR, geodeR, ore+oreR, clay+clayR, obsidian+obsidianR, geode+geodeR, time-1, memo, max)
		if result > max {
			max = result
		}

		if b.canMakeOre(ore) {
			result := dfs2(b, oreR+1, clayR, obsidianR, geodeR, ore+oreR-b.oreCost(), clay+clayR, obsidian+obsidianR, geode+geodeR, time-1, memo, max)
			if result > max {
				max = result
			}
		}

		if b.canMakeClay(ore) {
			result := dfs2(b, oreR, clayR+1, obsidianR, geodeR, ore+oreR-b.clayCost(), clay+clayR, obsidian+obsidianR, geode+geodeR, time-1, memo, max)
			if result > max {
				max = result
			}
		}
	}

	memo[memoKey] = max

	return max

}

func dfs2(b Blueprint, oreR, clayR, obsidianR, geodeR, ore, clay, obsidian, geode, time int, memo map[string]int, currentMax int) int {

	if time <= 0 {
		return geode
	}

	memoKey := fmt.Sprintf("%d,%d,%d,%d,%d,%d,%d,%d,%d", oreR, clayR, obsidianR, geodeR, ore, clay, obsidian, geode, time)
	if v, e := memo[memoKey]; e {
		return v
	}

	potentialMax := geode + geodeR*time + time*(time+1)/2
	if currentMax >= potentialMax {
		memo[memoKey] = 0
		return 0
	}

	if oreR > b.maxOreProd() || clayR >= b.maxClayProd() || obsidianR >= b.maxObsidianProd() {
		memo[memoKey] = 0
		return 0
	}

	max := geode

	if b.canMakeGeode(ore, obsidian) {
		oreCost, obsidianCost := b.geodeCost()
		result := dfs2(b, oreR, clayR, obsidianR, geodeR+1, ore+oreR-oreCost, clay+clayR, obsidian+obsidianR-obsidianCost, geode+geodeR, time-1, memo, max)
		if result > max {
			max = result
		}
	} else {
		if b.canMakeObsidian(ore, clay) {
			oreCost, clayCost := b.obsidianCost()
			result := dfs2(b, oreR, clayR, obsidianR+1, geodeR, ore+oreR-oreCost, clay+clayR-clayCost, obsidian+obsidianR, geode+geodeR, time-1, memo, max)
			if result > max {
				max = result
			}
		}

		if b.canMakeClay(ore) {
			result := dfs2(b, oreR, clayR+1, obsidianR, geodeR, ore+oreR-b.clayCost(), clay+clayR, obsidian+obsidianR, geode+geodeR, time-1, memo, max)
			if result > max {
				max = result
			}
		}

		if b.canMakeOre(ore) {
			result := dfs2(b, oreR+1, clayR, obsidianR, geodeR, ore+oreR-b.oreCost(), clay+clayR, obsidian+obsidianR, geode+geodeR, time-1, memo, max)
			if result > max {
				max = result
			}
		}

		result := dfs2(b, oreR, clayR, obsidianR, geodeR, ore+oreR, clay+clayR, obsidian+obsidianR, geode+geodeR, time-1, memo, max)
		if result > max {
			max = result
		}
	}

	memo[memoKey] = max

	return max

}

func part1(input string) int {
	result := 0
	for index, l := range strings.Split(input, "\n") {

		re := regexp.MustCompile("[0-9]+")
		strings := re.FindAllString(l, -1)[1:]

		numbers := make([]int, len(strings))
		for i, x := range strings {
			numbers[i], _ = strconv.Atoi(x)
		}

		blueprint := Blueprint{numbers}
		start := time.Now()
		dfsResult := dfs2(blueprint, 1, 0, 0, 0, 0, 0, 0, 0, 24, make(map[string]int), 0)
		fmt.Println(time.Since(start).Seconds())
		/*
			for _, p := range dfsPath {
				fmt.Println(p)
			}
		*/

		fmt.Println("dfs:", dfsResult)
		result += (index + 1) * dfsResult

	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func part2(input string) int {
	result := 1
	lines := strings.Split(input, "\n")
	for _, l := range lines[:min(3, len(lines))] {

		re := regexp.MustCompile("[0-9]+")
		strings := re.FindAllString(l, -1)[1:]

		numbers := make([]int, len(strings))
		for i, x := range strings {
			numbers[i], _ = strconv.Atoi(x)
		}

		blueprint := Blueprint{numbers}
		start := time.Now()
		dfsResult := dfs2(blueprint, 1, 0, 0, 0, 0, 0, 0, 0, 32, make(map[string]int), 0)
		fmt.Println(time.Since(start).Seconds())

		fmt.Println("dfs:", dfsResult)
		result *= dfsResult

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
