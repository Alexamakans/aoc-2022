package main

import (
	"strings"

	"github.com/Alexamakans/aoc-2022/pkg/aoc"
)

type forest [][]int

func parseForest(input string) forest {
	lines := aoc.SplitLines(input)
	// forest[y][x]
	var forest forest
	for _, line := range lines {
		digitsStr := strings.Split(line, "")
		var trees []int
		for _, digitStr := range digitsStr {
			trees = append(trees, aoc.StrToInt(digitStr))
		}
		forest = append(forest, trees)
	}
	return forest
}

func (f forest) isVisible(x, y int) bool {
	tree := f[y][x]
	dirsVisibleFrom := 4
	// Left
	for dx := -1; x+dx >= 0; dx-- {
		if f[y][x+dx] >= tree {
			dirsVisibleFrom--
			break
		}
	}
	// Right
	for dx := 1; x+dx < len(f[y]); dx++ {
		if f[y][x+dx] >= tree {
			dirsVisibleFrom--
			break
		}
	}
	// Up
	for dy := -1; y+dy >= 0; dy-- {
		if f[y+dy][x] >= tree {
			dirsVisibleFrom--
			break
		}
	}
	// Down
	for dy := 1; y+dy < len(f); dy++ {
		if f[y+dy][x] >= tree {
			dirsVisibleFrom--
			break
		}
	}

	return dirsVisibleFrom > 0
}

func (f forest) getScenicScore(x, y int) int {
	if x == 0 || y == 0 || x == len(f[0])-1 || y == len(f)-1 {
		// Edges can see 0 trees in one direction so their score is 0.
		return 0
	}

	tree := f[y][x]
	scenicScore := 1
	// Left
	for dx := -1; x+dx >= 0; dx-- {
		if f[y][x+dx] >= tree || x+dx == 0 {
			scenicScore *= -dx
			break
		}
	}
	// Right
	for dx := 1; x+dx < len(f[y]); dx++ {
		if f[y][x+dx] >= tree || x+dx == len(f[y])-1 {
			scenicScore *= dx
			break
		}
	}
	// Up
	for dy := -1; y+dy >= 0; dy-- {
		if f[y+dy][x] >= tree || y+dy == 0 {
			scenicScore *= -dy
			break
		}
	}
	// Down
	for dy := 1; y+dy < len(f); dy++ {
		if f[y+dy][x] >= tree || y+dy == len(f)-1 {
			scenicScore *= dy
			break
		}
	}

	return scenicScore
}
