package main

import (
	"regexp"
	"strings"

	"github.com/Alexamakans/aoc-2022/pkg/aoc"
)

func main() {
	aoc.Run("CMZ", "MCD", partOne, partTwo)
}

type stack []rune

// take takes from the stack, modifying the stack in place and returns the
// number of specified boxes from the beginning.
func (s *stack) take(numBoxes int) stack {
	taken := (*s)[:numBoxes]
	v := make([]rune, len(*s)-numBoxes)
	copy(v, (*s)[numBoxes:])
	*s = v
	return taken
}

func (s *stack) putOneAtATime(st stack) {
	st = aoc.ReverseSlice(st)
	// prepend
	*s = append(st, (*s)...)
}

func (s *stack) putMultipleAtATime(st stack) {
	// prepend
	*s = append(st, (*s)...)
}

type stacks []stack

func (s stacks) moveOneAtATime(from, to, numBoxes int) {
	taken := s[from].take(numBoxes)
	s[to].putOneAtATime(taken)
}

func (s stacks) moveMultipleAtATime(from, to, numBoxes int) {
	taken := s[from].take(numBoxes)
	s[to].putMultipleAtATime(taken)
}

func splitInput(input string) []string {
	return strings.Split(input, "\r\n\r\n")
}

func parseStacks(stacksInput string) stacks {
	lines := aoc.SplitLines(stacksInput)
	s := make([]stack, (len(lines[0])+1)/4)
	for lineIdx, line := range lines {
		if lineIdx == len(lines)-1 {
			break
		}
		for i := 1; i < len(line); i += 4 {
			stackIdx := (i - 1) / 4
			if line[i:i+1] == " " {
				continue
			}
			// prepend
			s[stackIdx] = append(s[stackIdx], rune(line[i]))
		}
	}
	return s
}

type move struct {
	from     int
	to       int
	numBoxes int
}

func parseMoves(movesInput string) []move {
	lines := aoc.SplitLines(movesInput)
	var moves []move
	for _, line := range lines {
		// move 1 from 2 to 1
		re := regexp.MustCompile(`\d+`)
		results := re.FindAllString(line, -1)
		moves = append(moves, move{
			numBoxes: aoc.StrToInt(results[0]),
			from:     aoc.StrToInt(results[1]) - 1,
			to:       aoc.StrToInt(results[2]) - 1,
		})
	}
	return moves
}

func getTopsJoined(s stacks) string {
	res := ""
	for _, st := range s {
		res += string(st[0])
	}
	return res
}

func partOne(input string) string {
	inputs := splitInput(input)
	s := parseStacks(inputs[0])
	moves := parseMoves(inputs[1])
	for _, move := range moves {
		s.moveOneAtATime(move.from, move.to, move.numBoxes)
	}
	return getTopsJoined(s)
}

func partTwo(input string) string {
	inputs := splitInput(input)
	s := parseStacks(inputs[0])
	moves := parseMoves(inputs[1])
	for _, move := range moves {
		s.moveMultipleAtATime(move.from, move.to, move.numBoxes)
	}
	return getTopsJoined(s)
}
