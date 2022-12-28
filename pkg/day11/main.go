package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Alexamakans/aoc-2022/pkg/aoc"
	"golang.org/x/exp/slices"
)

func main() {
	aoc.Run(10605, 2713310158, partOne, partTwo)
}

func partOne(input string) int64 {
	monkeys := parseMonkeys(input)

	for round := 0; round < 20; round++ {
		for _, monkey := range monkeys {
			for len(monkey.items) > 0 {
				monkey.timesInspected++
				monkey.items[0] = monkey.op(monkey.items[0])
				monkey.items[0] /= 3
				item := monkey.items[0]
				target := monkey.testFalseTarget
				if item%monkey.testDivisibleBy == 0 {
					target = monkey.testTrueTarget
				}

				monkeys[target].items = append(monkeys[target].items, item)
				monkey.items = monkey.items[1:]
			}
		}
	}

	slices.SortFunc(monkeys, func(a, b *monkey) bool {
		return b.timesInspected < a.timesInspected
	})

	fmt.Printf("%d * %d\n", monkeys[0].timesInspected, monkeys[1].timesInspected)
	return monkeys[0].timesInspected * monkeys[1].timesInspected
}

func partTwo(input string) int64 {
	monkeys := parseMonkeys(input)
	bigMod := getBigMod(monkeys)

	for round := 0; round < 10000; round++ {
		if round == 0 || round == 19 || round == 999 || round == 1999 {
			fmt.Printf("== After round %d ==\n", round+1)
		}
		for idx, monkey := range monkeys {
			for len(monkey.items) > 0 {
				monkey.timesInspected++
				monkey.items[0] = monkey.op(monkey.items[0])
				item := monkey.items[0]
				target := monkey.testFalseTarget
				if item%monkey.testDivisibleBy == 0 {
					target = monkey.testTrueTarget
				}

				item %= bigMod

				monkeys[target].items = append(monkeys[target].items, item)
				monkey.items = monkey.items[1:]
			}
			if round == 0 || round == 19 || round == 999 || round == 1999 {
				fmt.Printf("Monkey %d inspected items %d times.\n", idx, monkey.timesInspected)
			}
		}
		if round == 0 || round == 19 || round == 999 || round == 1999 {
			fmt.Println("")
		}
	}

	slices.SortFunc(monkeys, func(a, b *monkey) bool {
		return b.timesInspected < a.timesInspected
	})

	fmt.Printf("%d * %d\n", monkeys[0].timesInspected, monkeys[1].timesInspected)
	return int64(monkeys[0].timesInspected) * int64(monkeys[1].timesInspected)
}

type monkey struct {
	items           []int64
	op              func(item int64) int64
	testDivisibleBy int64
	testTrueTarget  int64
	testFalseTarget int64
	timesInspected  int64
}

func parseMonkeys(input string) []*monkey {
	var monkeys []*monkey
	lines := aoc.SplitLines(input)
	for idx := 0; idx < len(lines); idx += 7 {
		monkeys = append(monkeys, &monkey{
			items:           parseItems(lines[idx+1]),
			op:              parseOperation(lines[idx+2]),
			testDivisibleBy: parseTestDivisibleBy(lines[idx+3]),
			testTrueTarget:  parseTestTrueTarget(lines[idx+4]),
			testFalseTarget: parseTestFalseTarget(lines[idx+5]),
		})
	}

	return monkeys
}

var startingItemsRegexp = regexp.MustCompile(`Starting items: (.*)$`)

func parseItems(s string) []int64 {
	fmt.Println(s)
	matches := startingItemsRegexp.FindStringSubmatch(s)
	itemsStr := strings.Split(matches[1], ", ")
	var items []int64
	for _, v := range itemsStr {
		items = append(items, aoc.StrToInt64(v))
	}
	return items
}

var operationRegexp = regexp.MustCompile(`Operation: new = (old|\d+) (\+|\*) (old|\d+)`)

func parseOperation(s string) func(item int64) int64 {
	fmt.Println(s)
	matches := operationRegexp.FindStringSubmatch(s)
	leftStr := matches[1]
	operandStr := matches[2]
	rightStr := matches[3]
	return func(item int64) int64 {
		left := item
		if leftStr != "old" {
			left = aoc.StrToInt64(leftStr)
		}

		right := item
		if rightStr != "old" {
			right = aoc.StrToInt64(rightStr)
		}

		switch operandStr {
		case `+`:
			return left + right
		case `*`:
			return left * right
		}

		panic("what the heck")
	}
}

var testRegexp = regexp.MustCompile(`Test: divisible by (\d+)`)

func parseTestDivisibleBy(s string) int64 {
	fmt.Println(s)
	matches := testRegexp.FindStringSubmatch(s)
	numStr := matches[1]
	num := aoc.StrToInt64(numStr)
	return num
}

var targetTrueRegexp = regexp.MustCompile(`If true: throw to monkey (\d+)`)

func parseTestTrueTarget(s string) int64 {
	fmt.Println(s)
	matches := targetTrueRegexp.FindStringSubmatch(s)
	numStr := matches[1]
	num := aoc.StrToInt64(numStr)
	return num
}

var targetFalseRegexp = regexp.MustCompile(`If false: throw to monkey (\d+)`)

func parseTestFalseTarget(s string) int64 {
	fmt.Println(s)
	matches := targetFalseRegexp.FindStringSubmatch(s)
	numStr := matches[1]
	num := aoc.StrToInt64(numStr)
	return num
}

// courtesy of https://github.com/alexchao26/advent-of-code-go/blob/main/2022/day11/main.go
func getBigMod(monkeys []*monkey) int64 {
	var bigMod int64 = 1
	for _, m := range monkeys {
		bigMod *= m.testDivisibleBy
	}
	return bigMod
}
