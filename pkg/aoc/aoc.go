package aoc

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/constraints"
)

func Run(want1, want2 int, f1, f2 func(string) int) {
	{
		got1 := f1(GetTestInput())
		if got1 != want1 {
			panic(fmt.Errorf("got %d; want %d", got1, want1))
		}

		fmt.Printf("Passed the first test, got %d\n", got1)

		got1 = f1(GetInput())
		fmt.Printf("Part 1: %d\n", got1)
	}

	{
		got2 := f2(GetTestInput())
		if got2 != want2 {
			panic(fmt.Errorf("got %d; want %d", got2, want2))
		}

		fmt.Printf("Passed the second test, got %d\n", got2)

		got2 = f2(GetInput())
		fmt.Printf("Part 2: %d", got2)
	}
}

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

func SplitLines(s string) []string {
	return strings.Split(s, "\r\n")
}

func ChunkString(s string, numChunks int) []string {
	chunkSize := len(s) / numChunks
	var chunks []string
	for i := 0; i < len(s); i += chunkSize {
		start := i
		end := i + chunkSize
		chunks = append(chunks, s[start:end])
	}
	return chunks
}

func ChunkStringSlice(sSlice []string, numChunks int) [][]string {
	chunkSize := len(sSlice) / numChunks
	var groups [][]string
	for i := 0; i < len(sSlice); i += chunkSize {
		groups = append(groups, sSlice[i:i+chunkSize])
	}
	return groups
}

func SliceEqual[T constraints.Ordered](a, b []T) bool {
	for k := range a {
		if a[k] != b[k] {
			return false
		}
	}
	return true
}

func Contains[T constraints.Ordered](arr []T, a T) bool {
	for _, v := range arr {
		if v == a {
			return true
		}
	}
	return false
}

func ExcludeUniques[T constraints.Ordered](a, b []T) []T {
	var duplicates []T
	for _, v := range a {
		if !Contains(duplicates, v) {
			if Contains(b, v) {
				duplicates = append(duplicates, v)
			}
		}
	}

	return duplicates
}
