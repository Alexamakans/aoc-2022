package aoc

import (
	"fmt"
	"os"
)

func GetInput() string {
	s, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return string(s)
}

func GetTestInput() string {
	s, err := os.ReadFile("input_test.txt")
	if err != nil {
		panic(err)
	}

	return string(s)
}

func Check(input string, want int, f func(string) int) error {
	got := f(input)
	if got != want {
		return fmt.Errorf("got %d; want %d\n", got, want)
	}

	return nil
}
