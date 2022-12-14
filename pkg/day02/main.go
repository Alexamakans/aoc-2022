package main

import (
	"strings"

	"github.com/Alexamakans/aoc-2022/pkg/aoc"
)

func main() {
	aoc.Run(15, 12, partOne, partTwo)
}

type result int

const (
	resultWin result = iota
	resultLoss
	resultDraw
)

var resultScoreMappings = map[result]int{
	resultLoss: 0,
	resultDraw: 3,
	resultWin:  6,
}

var pickScoreMappings = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func partOne(input string) int {
	var resultMappings = map[string]result{
		"AX": resultDraw,
		"AY": resultWin,
		"AZ": resultLoss,

		"BX": resultLoss,
		"BY": resultDraw,
		"BZ": resultWin,

		"CX": resultWin,
		"CY": resultLoss,
		"CZ": resultDraw,
	}

	score := 0
	lines := aoc.SplitLines(input)
	for _, line := range lines {
		picks := strings.Split(line, " ")
		opponent, me := picks[0], picks[1]
		res := resultMappings[opponent+me]
		gain := resultScoreMappings[res] + pickScoreMappings[me]
		score += gain
	}
	return score
}

func partTwo(input string) int {
	var resultMappings = map[string]result{
		"X": resultLoss,
		"Y": resultDraw,
		"Z": resultWin,
	}

	pick := func(opp string, res result) string {
		if opp == "A" {
			if res == resultWin {
				return "Y"
			}
			if res == resultDraw {
				return "X"
			}
			if res == resultLoss {
				return "Z"
			}
		}
		if opp == "B" {
			if res == resultWin {
				return "Z"
			}
			if res == resultDraw {
				return "Y"
			}
			if res == resultLoss {
				return "X"
			}
		}
		if opp == "C" {
			if res == resultWin {
				return "X"
			}
			if res == resultDraw {
				return "Z"
			}
			if res == resultLoss {
				return "Y"
			}
		}
		return "wont happen with our inputs"
	}

	score := 0
	lines := aoc.SplitLines(input)
	for _, line := range lines {
		picks := strings.Split(line, " ")
		opponent, me := picks[0], picks[1]
		res := resultMappings[me]
		gain := resultScoreMappings[res] + pickScoreMappings[pick(opponent, res)]
		score += gain
	}
	return score
}
