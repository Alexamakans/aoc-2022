package main

import (
	"log"

	"github.com/Alexamakans/aoc-2022/pkg/aoc"
)

func main() {
	aoc.Run(157, 70, partOne, partTwo)
}

func partOne(input string) int {
	lines := aoc.SplitLines(input)
	const numChunks = 2
	score := 0
	for _, line := range lines {
		chunks := aoc.ChunkString(line, numChunks)
		duplicates := aoc.ExcludeUniques([]rune(chunks[0]), []rune(chunks[1]))
		for _, v := range duplicates {
			score += alphaScore(v)
		}
	}
	return score
}

func partTwo(input string) int {
	lines := aoc.SplitLines(input)
	const chunkSize = 3
	numChunks := len(lines) / chunkSize
	groups := aoc.ChunkSlice(lines, numChunks)
	score := 0
	for _, group := range groups {
		badge := getBadge([]rune(group[0]), []rune(group[1]), []rune(group[2]))
		score += alphaScore(badge)
	}
	return score
}

// Shared functions start

func alphaScore(r rune) int {
	i := int(r)

	if i >= 97 && i <= 122 { // a - z, 1 - 26
		return i - 96
	}

	if i >= 65 && i <= 90 { // A - Z, 27 - 52
		return i - 38
	}

	log.Fatalf("invalid value for alphaScore: %d", i)
	return -1
}

// Shared functions end

// Part two functions start

func getBadge(a, b, c []rune) rune {
	var uniques []rune
	for _, v := range a {
		if !aoc.Contains(uniques, v) {
			if aoc.Contains(b, v) && aoc.Contains(c, v) {
				return v
			}
		}
	}

	log.Fatalf("failed finding unique in\n\t%v\n\t%v\n\t%v", a, b, c)
	return -1
}

// Part two functions end
