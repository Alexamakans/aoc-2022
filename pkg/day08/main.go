package main

import (
	"github.com/Alexamakans/aoc-2022/pkg/aoc"
)

func main() {
	aoc.Run(21, 8, partOne, partTwo)
}

func partOne(input string) int {
	forest := parseForest(input)
	numVisibleTrees := 0
	for y := 0; y < len(forest); y++ {
		for x := 0; x < len(forest[y]); x++ {
			if forest.isVisible(x, y) {
				numVisibleTrees++
			}
		}
	}
	return numVisibleTrees
}

func partTwo(input string) int {
	forest := parseForest(input)
	bestScenicScore := 0
	for y := 0; y < len(forest); y++ {
		for x := 0; x < len(forest[y]); x++ {
			curScenicScore := forest.getScenicScore(x, y)
			if curScenicScore > bestScenicScore {
				bestScenicScore = curScenicScore
			}
		}
	}
	return bestScenicScore
}
