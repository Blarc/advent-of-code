package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

func mapValues(m map[string]int) []int {
	v := make([]int, 0, len(m))
	for _, value := range m {
		v = append(v, value)
	}
	return v
}

type Hand struct {
	cards    string
	bid      int
	strength int
}

func part1(input string) int {

	// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, 2
	var strength = map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
	}

	lines := strings.Split(input, "\n")

	hands := make([]Hand, 0, len(lines))
	for _, line := range lines {
		s := strings.Split(line, " ")
		cards := s[0]
		bid, _ := strconv.Atoi(s[1])

		c := make(map[string]int)
		for _, card := range cards {
			c[string(card)]++
		}

		v := mapValues(c)
		var handStrength int
		if len(c) == 1 {
			// Five of a kind
			// 1 1 1 1 1
			handStrength = strength["A"] + 6
		} else if len(c) == 2 {
			// Four of a kind or Full house
			// 1 1 1 1 2 or 1 1 2 2 2
			if v[0] == 4 || v[1] == 4 {
				handStrength = strength["A"] + 5
			} else {
				handStrength = strength["A"] + 4
			}
		} else if len(c) == 3 {
			// Three of a kind or Two pair
			// 1 1 1 2 3 or 1 1 2 2 3
			if v[0] == 3 || v[1] == 3 || v[2] == 3 {
				handStrength = strength["A"] + 3
			} else {
				handStrength = strength["A"] + 2
			}
		} else if len(c) == 4 {
			// One pair
			handStrength = strength["A"] + 1
		} else {
			// High card
			//split := strings.Split(cards, "")
			//maxCard := slices.MaxFunc(split, func(a, b string) int {
			//	return strength[a] - strength[b]
			//})
			handStrength = strength[string(cards[0])]
		}

		println(cards, bid, handStrength)
		fmt.Printf("%+v\n", c)
		fmt.Printf("%+v\n", v)
		hands = append(hands, Hand{cards, bid, handStrength})
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		if a.strength == b.strength {
			for i := 0; i < len(a.cards); i++ {
				if a.cards[i] != b.cards[i] {
					return strength[string(a.cards[i])] - strength[string(b.cards[i])]
				}
			}
		}
		return a.strength - b.strength
	})

	sum := 0
	for i, hand := range hands {
		t := (i + 1) * hand.bid
		//println(hand.cards, hand.strength, hand.bid, i+1, t)
		fmt.Printf("%s, %d\n", hand.cards, hand.bid)
		sum += t
	}

	return sum
}

func part2(input string) int {

	// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, 2
	var strength = map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": 1,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
	}

	lines := strings.Split(input, "\n")

	hands := make([]Hand, 0, len(lines))
	for _, line := range lines {
		s := strings.Split(line, " ")
		cards := s[0]
		bid, _ := strconv.Atoi(s[1])

		maxC := ""
		maxCV := 0
		c := make(map[string]int)
		for _, card := range cards {
			c[string(card)]++
			if c[string(card)] > maxCV && string(card) != "J" {
				maxC = string(card)
				maxCV = c[string(card)]
			}
		}
		if _, ok := c["J"]; ok {
			c[maxC] += c["J"]
			delete(c, "J")
		}

		v := mapValues(c)
		var handStrength int
		if len(c) == 1 {
			// Five of a kind
			// 1 1 1 1 1
			handStrength = strength["A"] + 6
		} else if len(c) == 2 {
			// Four of a kind or Full house
			// 1 1 1 1 2 or 1 1 2 2 2
			if v[0] == 4 || v[1] == 4 {
				handStrength = strength["A"] + 5
			} else {
				handStrength = strength["A"] + 4
			}
		} else if len(c) == 3 {
			// Three of a kind or Two pair
			// 1 1 1 2 3 or 1 1 2 2 3
			if v[0] == 3 || v[1] == 3 || v[2] == 3 {
				handStrength = strength["A"] + 3
			} else {
				handStrength = strength["A"] + 2
			}
		} else if len(c) == 4 {
			// One pair
			handStrength = strength["A"] + 1
		} else {
			// High card
			//split := strings.Split(cards, "")
			//maxCard := slices.MaxFunc(split, func(a, b string) int {
			//	return strength[a] - strength[b]
			//})
			handStrength = strength[string(cards[0])]
		}

		println(cards, bid, handStrength)
		fmt.Printf("%+v\n", c)
		fmt.Printf("%+v\n", v)
		hands = append(hands, Hand{cards, bid, handStrength})
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		if a.strength == b.strength {
			for i := 0; i < len(a.cards); i++ {
				if a.cards[i] != b.cards[i] {
					return strength[string(a.cards[i])] - strength[string(b.cards[i])]
				}
			}
		}
		return a.strength - b.strength
	})

	sum := 0
	for i, hand := range hands {
		t := (i + 1) * hand.bid
		//println(hand.cards, hand.strength, hand.bid, i+1, t)
		fmt.Printf("%s, %d\n", hand.cards, hand.bid)
		sum += t
	}

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
