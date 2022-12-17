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

type Valve2 struct {
	name         string
	tunnels      []string
	flow         int
	superTunnels map[string]int
}

type Result struct {
	flow int
	time int
	path string
}

type Result2 struct {
	flow   int
	time   int
	mePath string
	elPath string
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

func createKey2(me, el Valve2, meTime, elTime int, a map[string]bool) string {
	return fmt.Sprintf("%s,%s,%d,%d,%s", me.name, el.name, meTime, elTime, fmt.Sprint(a))
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

func dfs3(me Valve2, meTime int, valves map[string]Valve2, opened map[string]bool, path string) Result {
	if meTime <= 0 {
		return Result{0, meTime, path}
	}

	maxFlow := 0
	bestTime := 0
	bestPath := path

	for meKey, meValue := range me.superTunnels {
		// Open only me
		if !opened[meKey] && meTime-meValue-1 > 0 {
			opened[meKey] = true
			result := dfs3(valves[meKey], meTime-meValue-1, valves, opened, path+","+meKey+"(open)")
			flow := result.flow + valves[meKey].flow*(meTime-meValue-1)
			if flow > maxFlow {
				maxFlow = flow
				bestTime = result.time
				bestPath = result.path
			}
			opened[meKey] = false
		}
	}

	return Result{maxFlow, bestTime, bestPath}
}

func part1(input string) int {

	valvesRe := regexp.MustCompile("[A-Z][A-Z]")
	flowRe := regexp.MustCompile("[0-9]+")

	valves := make(map[string]Valve2)
	closedValves := make(map[string]Valve2)

	for _, l := range strings.Split(input, "\n") {

		v := valvesRe.FindAllString(l, -1)
		flow, _ := strconv.Atoi(flowRe.FindString(l))

		newValve := Valve2{v[0], v[1:], flow, make(map[string]int)}

		valves[v[0]] = newValve
		if flow > 0 {
			closedValves[v[0]] = newValve
		}
	}

	for _, a := range closedValves {
		for _, b := range closedValves {
			if a.name != b.name {
				a.superTunnels[b.name] = shortestPath(a, b, valves)
			}
		}
	}

	start := valves["AA"]
	for _, b := range closedValves {
		start.superTunnels[b.name] = shortestPath(start, b, valves)
	}

	// fmt.Println(notJammed)
	// fmt.Println(valves)
	//fmt.Println(createKey(10, map[string]bool{"haha": true, "fas": false}))
	// fmt.Println(dfs(valves["AA"], 31, valves, map[string]int{"AA": 1}, map[string]bool{}, 1, "AA"))
	fmt.Println(dfs3(valves["AA"], 30, valves, map[string]bool{}, "AA"))
	return 1
}

func dfs2(me, el Valve2, meTime, elTime int, valves map[string]Valve2, opened map[string]bool, memo map[string]Result2, mePath, elPath string) Result2 {
	if meTime <= 0 || elTime <= 0 {
		return Result2{0, 0, mePath, elPath}
	}

	saved, exists := memo[createKey2(me, el, meTime, elTime, opened)]
	if exists {
		return saved
	}

	maxFlow := 0
	bestTime := 0
	bestMePath := mePath
	bestElPath := elPath

	for meKey, meValue := range me.superTunnels {
		for elKey, elValue := range el.superTunnels {

			// Open both
			if meKey != elKey && !opened[meKey] && !opened[elKey] && meTime-meValue-1 > 0 && elTime-elValue-1 > 0 {
				opened[meKey] = true
				opened[elKey] = true

				result := dfs2(valves[meKey], valves[elKey], meTime-meValue-1, elTime-elValue-1, valves, opened, memo, mePath+","+meKey+"(open)", elPath+","+elKey+"(open)")
				flow := result.flow + valves[meKey].flow*(meTime-meValue-1) + valves[elKey].flow*(elTime-elValue-1)

				if flow > maxFlow {
					maxFlow = flow
					bestTime = result.time
					bestMePath = result.mePath
					bestElPath = result.elPath
				}

				opened[meKey] = false
				opened[elKey] = false
			} /*else if !opened[meKey] && meTime-meValue-1 > 0 {
				// Open only me, if meKey is not yet opened and me has enough time

				opened[meKey] = true
				result := dfs2(valves[meKey], valves[elKey], meTime-meValue-1, elTime-elValue, valves, opened, memo, mePath+","+meKey+"(open)", elPath)
				flow := result.flow + valves[meKey].flow*(meTime-meValue-1)
				if flow > maxFlow {
					maxFlow = flow
					bestTime = result.time
					bestMePath = result.mePath
					bestElPath = result.elPath
				}
				opened[meKey] = false
			} else if !opened[elKey] && elTime-elValue-1 > 0 {
				// Open only el, if elKey is not yet opened and el has enough time

				opened[elKey] = true
				result := dfs2(valves[meKey], valves[elKey], meTime-meValue, elTime-elValue-1, valves, opened, memo, mePath, elPath+","+elKey+"(open)")
				flow := result.flow + valves[elKey].flow*(elTime-elValue-1)

				if flow > maxFlow {
					maxFlow = flow
					bestTime = result.time
					bestMePath = result.mePath
					bestElPath = result.elPath
				}
				opened[elKey] = false
			}*/

		}
	}

	memo[createKey2(me, el, meTime, elTime, opened)] = Result2{maxFlow, bestTime, bestMePath, bestElPath}
	memo[createKey2(el, me, meTime, elTime, opened)] = Result2{maxFlow, bestTime, bestMePath, bestElPath}
	memo[createKey2(me, el, elTime, meTime, opened)] = Result2{maxFlow, bestTime, bestMePath, bestElPath}
	memo[createKey2(el, me, elTime, meTime, opened)] = Result2{maxFlow, bestTime, bestMePath, bestElPath}

	return Result2{maxFlow, bestTime, bestMePath, bestElPath}
}

type QueueNode struct {
	v Valve2
	d int
}

func shortestPath(a Valve2, b Valve2, valves map[string]Valve2) int {
	queue := []QueueNode{{a, 0}}
	visited := make(map[string]bool)

	for {
		// Pop
		v := queue[0]

		// Discard top element
		queue = queue[1:]

		if v.v.name == b.name {
			return v.d
		}

		for _, t := range v.v.tunnels {
			if !visited[t] {
				visited[t] = true
				queue = append(queue, QueueNode{valves[t], v.d + 1})
			}
		}
	}
}

func part2(input string) int {
	valvesRe := regexp.MustCompile("[A-Z][A-Z]")
	flowRe := regexp.MustCompile("[0-9]+")

	valves := make(map[string]Valve2)
	closedValves := make(map[string]Valve2)

	for _, l := range strings.Split(input, "\n") {

		v := valvesRe.FindAllString(l, -1)
		flow, _ := strconv.Atoi(flowRe.FindString(l))

		newValve := Valve2{v[0], v[1:], flow, make(map[string]int)}

		valves[v[0]] = newValve
		if flow > 0 {
			closedValves[v[0]] = newValve
		}
	}

	for _, a := range closedValves {
		for _, b := range closedValves {
			if a.name != b.name {
				a.superTunnels[b.name] = shortestPath(a, b, valves)
			}
		}
	}

	start := valves["AA"]
	for _, b := range closedValves {
		start.superTunnels[b.name] = shortestPath(start, b, valves)
	}

	/*
		fmt.Println("Computed paths")
		for _, cv := range closedValves {
			fmt.Println(cv)
		}
		fmt.Println(start)
	*/
	result := dfs2(start, start, 26, 26, closedValves, map[string]bool{}, map[string]Result2{}, "AA", "AA")
	fmt.Println("Me:", result.mePath)
	fmt.Println("El:", result.elPath)

	return result.flow
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
