package aoc

import (
	"fmt"
	"os"
)

func Run(want1, want2 int, f1, f2 func(string) int) {
	{
		got1 := f1(getTestInput())
		if got1 != want1 {
			panic(fmt.Errorf("got %d; want %d", got1, want1))
		}

		fmt.Printf("Passed the first test, got %d\n", got1)

		got1 = f1(getInput())
		fmt.Printf("Part 1: %d\n", got1)
	}

	{
		got2 := f2(getTestInput())
		if got2 != want2 {
			panic(fmt.Errorf("got %d; want %d", got2, want2))
		}

		fmt.Printf("Passed the second test, got %d\n", got2)

		got2 = f2(getInput())
		fmt.Printf("Part 2: %d", got2)
	}
}

func getInput() string {
	s, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return string(s)
}

func getTestInput() string {
	s, err := os.ReadFile("input_test.txt")
	if err != nil {
		panic(err)
	}

	return string(s)
}
