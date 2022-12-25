package main

import (
	"strings"

	"github.com/Alexamakans/aoc-2022/pkg/aoc"
)

func main() {
	aoc.Run(13, 36, partOne, partTwo)
}

func partOne(input string) int {
	lines := aoc.SplitLines(input)

	world := world{
		marks: make(map[string]bool),
	}
	rope := rope{
		head: &aoc.Position{X: 0, Y: 0},
		tail: []*aoc.Position{{X: 0, Y: 0}},
	}

	first := true
	for _, line := range lines {
		splt := strings.Split(line, " ")
		dir, amt := rune(splt[0][0]), aoc.StrToInt(splt[1])
		for i := 0; i < amt; i++ {
			rope.move(dir, !first)
			first = false
			world.mark(rope.tail[0].X, rope.tail[0].Y)
		}
	}
	return world.numMarks
}

func partTwo(input string) int {
	lines := aoc.SplitLines(input)

	world := world{
		marks: make(map[string]bool),
	}
	rope := rope{
		head: &aoc.Position{X: 0, Y: 0},
		tail: []*aoc.Position{
			{X: 0, Y: 0},
			{X: 0, Y: 0},
			{X: 0, Y: 0},
			{X: 0, Y: 0},
			{X: 0, Y: 0},
			{X: 0, Y: 0},
			{X: 0, Y: 0},
			{X: 0, Y: 0},
			{X: 0, Y: 0},
		},
	}

	first := true
	for _, line := range lines {
		splt := strings.Split(line, " ")
		dir, amt := rune(splt[0][0]), aoc.StrToInt(splt[1])
		for i := 0; i < amt; i++ {
			rope.move(dir, !first)
			first = false
			world.mark(rope.tail[8].X, rope.tail[8].Y)
		}
	}
	return world.numMarks
}
