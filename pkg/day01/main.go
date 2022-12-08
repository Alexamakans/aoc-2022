package main

import (
	"strconv"
	"strings"

	"github.com/Alexamakans/aoc-2022/pkg/aoc"
	"golang.org/x/exp/slices"
)

func main() {
	aoc.Run(24000, 45000, partOne, partTwo)
}

func partOne(input string) int {
	sortedElves := getSortedElves(input)
	sortedElvesCount := len(sortedElves)
	return sortedElves[sortedElvesCount-1]
}

func partTwo(input string) int {
	sortedElves := getSortedElves(input)
	sortedElvesCount := len(sortedElves)
	return sortedElves[sortedElvesCount-1] + sortedElves[sortedElvesCount-2] + sortedElves[sortedElvesCount-3]
}

func getSortedElves(input string) []int {
	var elves []int
	chunks := strings.Split(input, "\r\n\r\n")
	for _, chunk := range chunks {
		sum := 0
		lines := strings.Split(chunk, "\r\n")
		for _, line := range lines {
			calories, err := strconv.ParseInt(line, 10, 32)
			if err != nil {
				panic(err)
			}

			sum += int(calories)
		}

		elves = append(elves, sum)
	}
	slices.Sort(elves)
	return elves
}
