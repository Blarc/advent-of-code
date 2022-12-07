package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

func part1(input string) int {

	structure := make(map[string]int)
	path := []string{}

	for _, l := range strings.Split(input, "\n") {

		s := strings.Split(l, " ")

		if s[0] == "$" {
			if s[1] == "cd" {
				// fmt.Println("Moving to:", s[2])
				if s[2] == "/" {
					// Switches the current directory to the outermost directory
					path = []string{"/"}
				} else if s[2] == ".." {
					// Moves out one level
					path = path[:len(path)-1]
				} else {
					// Moves in one level
					path = append(path, s[2])
				}
				// fmt.Println(path)
			}
		} else if s[0] == "dir" {
			// fmt.Println("Directory:", l)
		} else {
			// fmt.Println("File size:", l)
			for i := 0; i < len(path); i++ {
				fileSize, _ := strconv.Atoi(s[0])
				structure[strings.Join(path[0:i+1], ".")] += fileSize
			}
		}
	}

	result := 0
	for _, dirSize := range structure {
		if dirSize <= 100000 {
			result += dirSize
		}
	}
	return result
}

func part2(input string) int {
	structure := make(map[string]int)
	path := []string{}

	for _, l := range strings.Split(input, "\n") {

		s := strings.Split(l, " ")

		if s[0] == "$" {
			if s[1] == "cd" {
				// fmt.Println("Moving to:", s[2])
				if s[2] == "/" {
					// Switches the current directory to the outermost directory
					path = []string{"/"}
				} else if s[2] == ".." {
					// Moves out one level
					path = path[:len(path)-1]
				} else {
					// Moves in one level
					path = append(path, s[2])
				}
				// fmt.Println(path)
			}
		} else if s[0] == "dir" {
			// fmt.Println("Directory:", l)
		} else {
			// fmt.Println("File size:", l)
			for i := 0; i < len(path); i++ {
				fileSize, _ := strconv.Atoi(s[0])
				structure[strings.Join(path[0:i+1], ".")] += fileSize
			}
		}
	}

	totalDiskSpace := 70000000
	neededSpace := 30000000

	unusedDiskSpace := totalDiskSpace - structure["/"]
	neededSpace -= unusedDiskSpace

	result := 70000000
	for _, dirSize := range structure {
		if dirSize >= neededSpace && dirSize < result {
			result = dirSize
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
