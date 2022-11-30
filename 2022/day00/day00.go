package day00

import (
	"math"
	"strconv"
	"strings"
)

func part1(input string) (int, error) {
	splits := strings.Split(strings.TrimSpace(input), "\n")

	previous := math.MaxInt
	count := 0
	for _, s := range splits {
		value, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		if value > previous {
			count++
		}
		previous = value
	}

	return count, nil
}

func part2(input string) (int, error) {
	splits := strings.Split(strings.TrimSpace(input), "\n")

	const windowSize = 3
	values := make([]int, len(splits))
	count := 0
	for i, s := range splits {
		value, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		values[i] = value

		indexToRemove := i - windowSize
		if indexToRemove >= 0 && value > values[indexToRemove] {
			count++
		}
	}

	return count, nil
}
