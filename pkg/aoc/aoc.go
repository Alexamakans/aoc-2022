package aoc

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

func Run[T constraints.Ordered, T2 constraints.Ordered](want1 T, want2 T2, f1 func(input string) T, f2 func(input string) T2) {
	{
		got1 := f1(GetTestInput())
		if got1 != want1 {
			panic(fmt.Errorf("got %v; want %v", got1, want1))
		}

		fmt.Printf("Passed the first test, got %v\n", got1)

		got1 = f1(GetInput())
		fmt.Printf("Part 1: %v\n", got1)
	}

	{
		got2 := f2(GetTestInput2())
		if got2 != want2 {
			panic(fmt.Errorf("got %v; want %v", got2, want2))
		}

		fmt.Printf("Passed the second test, got %v\n", got2)

		got2 = f2(GetInput())
		fmt.Printf("Part 2: %v", got2)
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

func GetTestInput2() string {
	s, err := os.ReadFile("input_test2.txt")
	if err != nil {
		return GetTestInput()
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

func ChunkSlice[T any](slice []T, numChunks int) [][]T {
	chunkSize := len(slice) / numChunks
	var chunks [][]T
	for i := 0; i < len(slice); i += chunkSize {
		chunks = append(chunks, slice[i:i+chunkSize])
	}
	return chunks
}

func EqualSlice[T constraints.Ordered](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
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

func StrToInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func StrToInt64(s string) int64 {
	v, _ := strconv.ParseInt(s, 10, 64)
	return v
}

func IntToStr(v int) string {
	return strconv.Itoa(v)
}

func ReverseSlice[T any](s []T) []T {
	a := make([]T, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

func Clamp(v, min, max int) int {
	return int(math.Min(float64(max), math.Max(float64(v), float64(min))))
}
