package main

import (
	"strings"

	"github.com/Alexamakans/aoc-2022/pkg/aoc"
)

func main() {
	aoc.Run(2, 4, partOne, partTwo)
}

type section struct {
	start int
	end   int
}

type pair struct {
	a section
	b section
}

func (p pair) isOneFullyContained() bool {
	if p.a.start >= p.b.start && p.a.end <= p.b.end {
		return true
	}

	if p.b.start >= p.a.start && p.b.end <= p.a.end {
		return true
	}

	return false
}

func (p pair) isOverlapping() bool {
	if p.a.end < p.b.start || p.a.start > p.b.end {
		return false
	}

	return true
}

func lineToPair(line string) pair {
	s := strings.Split(line, ",")
	a, b := strings.Split(s[0], "-"), strings.Split(s[1], "-")
	p := pair{}
	p.a.start = aoc.StrToInt(a[0])
	p.a.end = aoc.StrToInt(a[1])
	p.b.start = aoc.StrToInt(b[0])
	p.b.end = aoc.StrToInt(b[1])

	return p
}

func partOne(input string) int {
	lines := aoc.SplitLines(input)
	count := 0
	for _, line := range lines {
		p := lineToPair(line)
		if p.isOneFullyContained() {
			count++
		}
	}
	return count
}

func partTwo(input string) int {
	lines := aoc.SplitLines(input)
	count := 0
	for _, line := range lines {
		p := lineToPair(line)
		if p.isOverlapping() {
			count++
		}
	}
	return count
}
