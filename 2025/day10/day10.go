package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/aclements/go-z3/z3"
)

import (
	_ "embed"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func ToIntSlice(s []string) []int {
	r := make([]int, len(s))
	for i, v := range s {
		r[i], _ = strconv.Atoi(v)
	}
	return r
}

type LightsState struct {
	state   []bool
	presses int
}

func (m LightsState) PressButton(b []int) LightsState {
	newState := make([]bool, len(m.state))
	for i, v := range m.state {
		newState[i] = v
	}

	for _, v := range b {
		newState[v] = !newState[v]
	}

	return LightsState{
		state:   newState,
		presses: m.presses + 1,
	}
}

func (m LightsState) IsGoal(g []bool) bool {
	for i, v := range m.state {
		if v != g[i] {
			return false
		}
	}
	return true
}

type JoltageState struct {
	state   []int
	presses int
}

func (j JoltageState) PressButton(b []int) JoltageState {
	return j.PressButtonN(b, 1)
}

func (j JoltageState) PressButtonN(b []int, n int) JoltageState {
	newState := make([]int, len(j.state))
	for i, v := range j.state {
		newState[i] = v
	}

	for _, v := range b {
		newState[v] += n
	}

	return JoltageState{
		state:   newState,
		presses: j.presses + n,
	}
}

func (j JoltageState) MaxPresses(b, g []int) int {
	maxPresses := math.MaxInt32
	for _, v := range b {
		tmp := g[v] - j.state[v]
		if tmp < maxPresses {
			maxPresses = tmp
		}
	}

	if maxPresses == math.MaxInt32 {
		return 0
	}

	if maxPresses < 0 {
		panic("Negative max presses")
	}

	return maxPresses
}

func (j JoltageState) IsGoal(g []int) bool {
	for i, v := range j.state {
		if v != g[i] {
			return false
		}
	}
	return true
}

func (j JoltageState) IsInvalid(g []int) bool {
	for i, v := range j.state {
		if v > g[i] {
			return true
		}
	}
	return false
}

func (j JoltageState) String() string {
	return fmt.Sprintf("%v:%d", j.state, j.presses)
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	result := 0

	for _, l := range lines {
		fields := strings.Fields(l)

		goal := make([]bool, len(fields[0])-2)
		for j, c := range fields[0][1 : len(fields[0])-1] {
			if c == '#' {
				goal[j] = true
			}
		}

		buttons := make([][]int, len(fields)-2)
		for j, b := range fields[1 : len(fields)-1] {
			buttons[j] = ToIntSlice(strings.Split(b[1:len(b)-1], ","))
		}

		queue := []LightsState{{
			state:   make([]bool, len(goal)),
			presses: 0,
		}}

		for len(queue) > 0 {
			m := queue[0]
			queue = queue[1:]

			if m.IsGoal(goal) {
				result += m.presses
				// fmt.Printf("Machine %d: %d\n", i, m.presses)
				break
			}

			for _, b := range buttons {
				newLightState := m.PressButton(b)
				queue = append(queue, newLightState)
			}
		}

	}

	return result
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	result := 0

	for _, l := range lines {
		fields := strings.Fields(l)

		// Parse the goal from the last field by removing the brackets and parsing integers
		goal := ToIntSlice(strings.Split(fields[len(fields)-1][1:len(fields[len(fields)-1])-1], ","))

		ctx := z3.NewContext(nil)
		solver := z3.NewSolver(ctx)

		intSort := ctx.IntSort()
		zero := ctx.FromInt(0, intSort).(z3.Int)

		buttons := make([]z3.Int, len(fields)-2)
		targetsEquations := make([][]z3.Int, len(goal))

		for j, b := range fields[1 : len(fields)-1] {
			buttons[j] = ctx.IntConst(fmt.Sprintf("b%d", j))

			// Number of presses must be non-negative
			solver.Assert(buttons[j].GE(zero))

			// Parse which joltage levels are affected by this button
			affectedJoltageLevels := ToIntSlice(strings.Split(b[1:len(b)-1], ","))
			for _, a := range affectedJoltageLevels {
				targetsEquations[a] = append(targetsEquations[a], buttons[j])
			}
		}

		for j, g := range goal {
			target := ctx.FromInt(int64(g), intSort).(z3.Int)
			sum := targetsEquations[j][0]
			// Create the equation for the target
			for _, s := range targetsEquations[j][1:] {
				sum = sum.Add(s)
			}
			solver.Assert(sum.Eq(target))
		}

		// Add up all the button presses
		total := buttons[0]
		for _, b := range buttons[1:] {
			total = total.Add(b)
		}

		minResult := -1
		for {
			sat, err := solver.Check()
			if !sat || err != nil {
				break
			}

			model := solver.Model()
			res := model.Eval(total, true)
			val, _, _ := res.(z3.Int).AsInt64()

			minResult = int(val)

			cur := ctx.FromInt(val, intSort).(z3.Int)
			solver.Assert(total.LT(cur))
		}

		result += minResult
	}

	// 4754 too low
	return result
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
