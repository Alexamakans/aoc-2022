package main

import (
	"testing"

	"github.com/Alexamakans/aoc-2022/pkg/aoc"
)

func Test_AllInputIsEvenLength(t *testing.T) {
	assertAllLinesAreEvenLength(t, aoc.GetTestInput())
	assertAllLinesAreEvenLength(t, aoc.GetInput())
}

func assertAllLinesAreEvenLength(t *testing.T, s string) {
	lines := aoc.SplitLines(s)
	for _, line := range lines {
		if len(line)%2 != 0 {
			t.Fatalf("line '%s' was not of even length", line)
		}
	}
}
