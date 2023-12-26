package main

// Disclaimer: this code is ugly and I have no idea how I made it work, but it works... and this is Advent of Code, so be nice.

import (
	"flag"
	"fmt"
	"math/rand"
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

// Flip-flop (%) - on or off, initially off, ignores high pulse, on off pulse:
// - if off, turn on and send high
// - if on, turn off and send low
// Conjunction (&) - remember the type of the most recent pulse from each
// - default: all low
// - first update
// - if all high, send low
// - otherwise, send high
// broadcaster - receives a pulse, send to all
// button - low to broadcaster

type Module interface {
	process(from string, pulse bool) map[string]map[string]bool
	addInput(input string)
	getInputs() []string
}

type Output struct {
	name   string
	inputs map[string]bool
}

func (o *Output) process(from string, pulse bool) map[string]map[string]bool {
	return make(map[string]map[string]bool)
}

func (o *Output) addInput(input string) {
	if o.inputs == nil {
		o.inputs = make(map[string]bool)
	}
	o.inputs[input] = false
}

func (o *Output) getInputs() []string {
	var tmp []string
	for k, _ := range o.inputs {
		tmp = append(tmp, k)
	}
	return tmp
}

type Broadcast struct {
	name   string
	inputs map[string]bool
	dest   []string
}

func (b *Broadcast) process(from string, pulse bool) map[string]map[string]bool {
	//fmt.Println(b.name, "sending", pulse, "to", b.dest)
	result := make(map[string]map[string]bool)
	result[b.name] = make(map[string]bool)
	for _, d := range b.dest {
		result[b.name][d] = pulse
	}
	return result
}

func (b *Broadcast) addInput(input string) {
	if b.inputs == nil {
		b.inputs = make(map[string]bool)
	}
	b.inputs[input] = false
}

func (b *Broadcast) getInputs() []string {
	var tmp []string
	for k, _ := range b.inputs {
		tmp = append(tmp, k)
	}
	return tmp
}

type FlipFlop struct {
	name     string
	dest     []string
	inputs   map[string]bool
	oldPulse bool
}

func (f *FlipFlop) process(from string, pulse bool) map[string]map[string]bool {
	// If low pulse
	if !pulse {
		f.oldPulse = !f.oldPulse
		//fmt.Println(f.name, "sending", f.oldPulse, "to", f.dest)
		result := make(map[string]map[string]bool)
		result[f.name] = make(map[string]bool)

		for _, d := range f.dest {
			result[f.name][d] = f.oldPulse
		}
		return result
	}
	return make(map[string]map[string]bool)
}

func (f *FlipFlop) addInput(input string) {
	if f.inputs == nil {
		f.inputs = make(map[string]bool)
	}
	f.inputs[input] = false
}

func (f *FlipFlop) getInputs() []string {
	var tmp []string
	for k, _ := range f.inputs {
		tmp = append(tmp, k)
	}
	return tmp
}

type Conjunction struct {
	name   string
	dest   []string
	inputs map[string]bool
}

func (c *Conjunction) process(from string, pulse bool) map[string]map[string]bool {
	c.inputs[from] = pulse

	for _, inputPulse := range c.inputs {
		if !inputPulse {
			//fmt.Println(c.name, "sending", true, "to", c.dest)
			result := make(map[string]map[string]bool)
			result[c.name] = make(map[string]bool)
			for _, d := range c.dest {
				result[c.name][d] = true
			}
			return result
		}
	}
	//fmt.Println(c.name, "sending", false, "to", c.dest)
	result := make(map[string]map[string]bool)
	result[c.name] = make(map[string]bool)
	for _, d := range c.dest {
		result[c.name][d] = false
	}
	return result
}

func (c *Conjunction) addInput(input string) {
	if c.inputs == nil {
		c.inputs = make(map[string]bool)
	}
	c.inputs[input] = false
}

func (b *Conjunction) getInputs() []string {
	var tmp []string
	for k, _ := range b.inputs {
		tmp = append(tmp, k)
	}
	return tmp
}

func part1(input string) int {

	modules := make(map[string]Module)

	for _, l := range strings.Split(input, "\n") {
		fromTo := strings.Split(l, " -> ")
		from := fromTo[0]
		to := strings.Split(fromTo[1], ", ")

		if from == "broadcaster" {
			modules[from] = &Broadcast{name: from, dest: to}
		} else if string(from[0]) == "%" {
			modules[from[1:]] = &FlipFlop{name: from[1:], dest: to}
		} else if string(from[0]) == "&" {
			modules[from[1:]] = &Conjunction{name: from[1:], dest: to}
		}
	}

	for _, l := range strings.Split(input, "\n") {
		fromTo := strings.Split(l, " -> ")
		from := fromTo[0]
		to := strings.Split(fromTo[1], ", ")

		for _, d := range to {
			if conjunction, isConjunction := modules[d].(*Conjunction); isConjunction {
				if from == "broadcaster" {
					conjunction.addInput(from)
				} else {
					conjunction.addInput(from[1:])
				}
			}
		}
	}

	low := 0
	high := 0
	for i := 0; i < 1000; i++ {
		fmt.Println("---", i, "---")
		next := modules["broadcaster"].process("", false)
		low += 1
		for len(next) > 0 {
			fmt.Println("---")
			fmt.Println(next)
			newNext := make(map[string]map[string]bool)
			for fromKey, from := range next {
				fromKey = strings.Split(fromKey, ":")[0]
				for toKey, pulse := range from {
					if pulse {
						high++
					} else {
						low++
					}
					if _, exists := modules[toKey]; exists {
						for km, vm := range modules[toKey].process(fromKey, pulse) {
							//fmt.Println("h:", km, vm)
							//if _, exist := newNext[km]; exist {
							//	for kvm, vvm := range vm {
							//		if _, exist2 := newNext[km][kvm]; exist2 {
							//			panic("WHAT TO DO")
							//		}
							//		newNext[km][kvm] = vvm
							//	}
							//} else {
							//	newNext[key] = vm
							//}
							key := fmt.Sprintf("%s:%d", km, rand.Int())
							newNext[key] = vm
						}
					}
				}
			}
			next = newNext
		}
	}

	// 736377488 too low
	fmt.Println(low, high)
	return low * high
}

func part2(input string) uint64 {
	modules := make(map[string]Module)

	for _, l := range strings.Split(input, "\n") {
		fromTo := strings.Split(l, " -> ")
		from := fromTo[0]
		to := strings.Split(fromTo[1], ", ")

		if from == "broadcaster" {
			modules[from] = &Broadcast{name: from, dest: to}
		} else if string(from[0]) == "%" {
			modules[from[1:]] = &FlipFlop{name: from[1:], dest: to}
		} else if string(from[0]) == "&" {
			modules[from[1:]] = &Conjunction{name: from[1:], dest: to}
		}
	}

	for _, l := range strings.Split(input, "\n") {
		fromTo := strings.Split(l, " -> ")
		from := fromTo[0]
		to := strings.Split(fromTo[1], ", ")

		for _, d := range to {
			module, exists := modules[d]
			if !exists && d != "rx" {
				panic("something is wrong")
			} else if !exists {
				modules["rx"] = &Output{name: "rx"}
				module = modules["rx"]
			}

			if from == "broadcaster" {
				module.addInput(from)
			} else {
				module.addInput(from[1:])
			}
		}
	}

	mainModules := []string{"rx"}
	index := 0
	for len(mainModules) > 0 && index < 2 {
		var newMainModules []string
		for _, mainModule := range mainModules {
			fmt.Print(mainModule, " ")
			for _, moduleInput := range modules[mainModule].getInputs() {
				newMainModules = append(newMainModules, moduleInput)
			}
		}
		fmt.Println()
		mainModules = newMainModules
		index++
	}

	mainModulesMap := make(map[string][]int)
	for _, m := range mainModules {
		//_, isFlipFlop := modules[m].(*FlipFlop)
		mainModulesMap[m] = []int{}
	}

	for i := 1; i < 50000; i++ {
		//for k, v := range modules["lv"].(*Conjunction).inputs {
		//	fmt.Println(k, v)
		//}
		//
		//fmt.Println(modules["lv"].(*Conjunction).inputs)
		//fmt.Println("---", i, "---")
		next := modules["broadcaster"].process("", false)
		for len(next) > 0 {
			//fmt.Println("---")
			//fmt.Println(next)
			newNext := make(map[string]map[string]bool)
			for fromKey, from := range next {
				fromKey = strings.Split(fromKey, ":")[0]
				for toKey, pulse := range from {
					if _, exists := modules[toKey]; exists {
						for km, vm := range modules[toKey].process(fromKey, pulse) {
							for kvm, vvm := range vm {
								if _, exists3 := mainModulesMap[kvm]; exists3 && !vvm {
									if len(mainModulesMap[kvm]) == 0 || mainModulesMap[kvm][len(mainModulesMap[kvm])-1] != i {
										mainModulesMap[kvm] = append(mainModulesMap[kvm], i)
									}
								}
							}
							key := fmt.Sprintf("%s:%d", km, rand.Int())
							newNext[key] = vm
						}
					}
				}
			}
			next = newNext
		}
	}

	maxIndex := 0
	ans := uint64(1)
	for k, m := range mainModulesMap {
		i := abs(m[1] - m[0])
		fmt.Println(k)
		fmt.Println(m[:min(len(m)-1, 60)], i)
		for j := 0; j < min(len(m)-1, 60); j++ {
			fmt.Printf("%d ", m[j+1]-m[j])
		}
		fmt.Println()

		if i == 0 {
			fmt.Println(k)
			panic("zero")
		}
		if m[0] != 0 {
			ans = LCM(ans, uint64(m[0]))
		} else {
			ans = LCM(ans, uint64(m[1]))
		}
		if i > maxIndex {
			maxIndex = i
		}
		//fmt.Println(ans)
	}
	// too high: 44928388009487248
	// incorrect: 7284368133192
	// incorrect: 33858711810
	// too low: 16929357953
	fmt.Println(ans, maxIndex)
	return ans
}

func GCD(a, b uint64) uint64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b uint64, integers ...uint64) uint64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
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

	start := time.Now()
	if part == 1 {
		fmt.Println("Result:", part1(strings.TrimSpace(inputText)))
	} else {
		fmt.Println("Result:", part2(strings.TrimSpace(inputText)))
	}
	fmt.Println("Time:", fmt.Sprintf("%d ms", time.Since(start).Milliseconds()))
}
