package main

import (
	"testing"

	"github.com/Alexamakans/aoc-2022/pkg/aoc"
)

func Test_TakeStack(t *testing.T) {
	testCases := []struct {
		value     stack
		numBoxes  int
		valueWant stack
		want      stack
	}{
		{
			value:     []rune(`ABCDEF`),
			numBoxes:  2,
			valueWant: []rune(`CDEF`),
			want:      []rune(`AB`),
		},
	}

	for _, tc := range testCases {
		orig := make([]rune, len(tc.value))
		copy(orig, tc.value)
		got := tc.value.take(tc.numBoxes)
		if !aoc.EqualSlice(got, tc.want) {
			t.Fatalf("1: got %v; want %v (original=%v, numBoxes=%d)", got, tc.want, orig, tc.numBoxes)
		}

		if !aoc.EqualSlice(tc.value, tc.valueWant) {
			t.Fatalf("2: got %v; want %v (original=%v, numBoxes=%d)", tc.value, tc.valueWant, orig, tc.numBoxes)
		}
	}
}

func Test_PutStack(t *testing.T) {
	testCases := []struct {
		valueTarget stack
		valuePut    stack
		want        stack
	}{
		{
			valueTarget: []rune(`ABC`),
			valuePut:    []rune(`DE`),
			want:        []rune(`EDABC`),
		},
	}

	for _, tc := range testCases {
		orig := make([]rune, len(tc.valueTarget))
		copy(orig, tc.valueTarget)

		tc.valueTarget.putOneAtATime(tc.valuePut)
		if !aoc.EqualSlice(tc.valueTarget, tc.want) {
			t.Fatalf("got %v; want %v (original=%v, put=%v)", tc.valueTarget, tc.want, orig, tc.valuePut)
		}
	}
}

func Test_MoveStack(t *testing.T) {
	testCases := []struct {
		origStacks  stacks
		valueStacks stacks
		valueMove   move
		wantStacks  stacks
	}{
		{
			origStacks:  []stack{[]rune("ABCDEF"), []rune("GHIJKL")},
			valueStacks: []stack{[]rune("ABCDEF"), []rune("GHIJKL")},
			valueMove:   move{from: 0, to: 1, numBoxes: 3},
			wantStacks:  []stack{[]rune("DEF"), []rune("CBAGHIJKL")},
		},
	}

	for _, tc := range testCases {
		tc.valueStacks.moveOneAtATime(tc.valueMove.from, tc.valueMove.to, tc.valueMove.numBoxes)
		for k := range tc.valueStacks {
			if !aoc.EqualSlice(tc.valueStacks[k], tc.wantStacks[k]) {
				t.Fatalf("got %v; want %v (move=%v, original=%v, gotStack=%v, wantStack=%v)", tc.valueStacks[k], tc.wantStacks[k], tc.valueMove, tc.origStacks, tc.valueStacks, tc.wantStacks)
			}
		}
	}
}

func Test_ParseMoves(t *testing.T) {
	input := aoc.GetTestInput()
	inputs := splitInput(input)
	moves := parseMoves(inputs[1])
	i := 0
	if moves[i].numBoxes != 1 && moves[i].from != 2 && moves[i].to != 1 {
		t.Fatalf("%d: got move %d from %d to %d", i, moves[i].numBoxes, moves[i].from, moves[i].to)
	}
	i++
	if moves[i].numBoxes != 3 && moves[i].from != 1 && moves[i].to != 3 {
		t.Fatalf("%d: got move %d from %d to %d", i, moves[i].numBoxes, moves[i].from, moves[i].to)
	}
	i++
	if moves[i].numBoxes != 2 && moves[i].from != 2 && moves[i].to != 1 {
		t.Fatalf("%d: got move %d from %d to %d", i, moves[i].numBoxes, moves[i].from, moves[i].to)
	}
	i++
	if moves[i].numBoxes != 1 && moves[i].from != 1 && moves[i].to != 2 {
		t.Fatalf("%d: got move %d from %d to %d", i, moves[i].numBoxes, moves[i].from, moves[i].to)
	}
}
