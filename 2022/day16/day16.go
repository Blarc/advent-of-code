package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

type Valve struct {
	name    string
	tunnels []string
	flow    int
	saved   map[string]Result
}

type Result struct {
	flow int
	time int
	path string
}

func createKey(time int, a map[string]bool) string {
	r := fmt.Sprint(time)
	for k, v := range a {
		if v {
			r += k
		}
	}
	return r
}

func dfs(valve Valve, time int, valves map[string]Valve, visited map[string]int, opened map[string]bool, depth int, path string) Result {
	// fmt.Printf("%*s %s\n", depth*2, valve.name, valve.tunnels)
	// || opened >= notJammed || (notJammed-opened)*2 >= time
	// test := "AA,DD(open),CC,BB(open),AA,II,JJ(open),II,AA,DD,EE,FF,GG,HH(open),GG,FF,EE(open),DD,CC"
	if time <= 0 {
		// fmt.Println(1, path)
		/*
			if strings.HasPrefix(path, test) {
				fmt.Println(1, path)
			}
		*/
		return Result{0, 0, path}
	}

	if visited[valve.name] > len(valve.tunnels) {
		// fmt.Println(2, path)
		/*
			if strings.HasPrefix(path, test) {
				fmt.Println(2, path)
			}
		*/
		return Result{0, time, path}
	}

	savedFlow, exists := valve.saved[createKey(time, opened)]
	if exists {
		return savedFlow
	}

	maxFlow := -1
	bestTime := -1
	bestPath := ""
	v, e := opened[valve.name]
	if valve.flow > 0 && (!e || !v) {
		for _, tunnel := range valve.tunnels {

			visited[valve.name] += 1
			opened[valve.name] = true
			result := dfs(valves[tunnel], time-2, valves, visited, opened, depth+1, path+"(open),"+tunnel)
			flow := result.flow + valve.flow*(time-2)
			opened[valve.name] = false
			visited[valve.name] -= 1

			if flow > maxFlow {
				maxFlow = flow
				bestTime = result.time
				bestPath = result.path
			}
		}
	}

	for _, tunnel := range valve.tunnels {
		visited[valve.name] += 1
		result := dfs(valves[tunnel], time-1, valves, visited, opened, depth+1, path+","+tunnel)
		visited[valve.name] -= 1

		if result.flow > maxFlow {
			maxFlow = result.flow
			bestTime = result.time
			bestPath = result.path

		}
	}

	valve.saved[createKey(time, opened)] = Result{maxFlow, bestTime, bestPath}

	return Result{maxFlow, bestTime, bestPath}
}

func part1(input string) int {

	valvesRe := regexp.MustCompile("[A-Z][A-Z]")
	flowRe := regexp.MustCompile("[0-9]+")

	valves := make(map[string]Valve)
	notJammed := 0

	for _, l := range strings.Split(input, "\n") {

		v := valvesRe.FindAllString(l, -1)
		flow, _ := strconv.Atoi(flowRe.FindString(l))
		valves[v[0]] = Valve{v[0], v[1:], flow, map[string]Result{}}
		if flow > 0 {
			notJammed++
		}
	}

	// fmt.Println(notJammed)
	// fmt.Println(valves)
	//fmt.Println(createKey(10, map[string]bool{"haha": true, "fas": false}))
	fmt.Println(dfs(valves["AA"], 31, valves, map[string]int{"AA": 1}, map[string]bool{}, 1, "AA"))
	return 1
}

func part2(input string) int {
	return 2
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
