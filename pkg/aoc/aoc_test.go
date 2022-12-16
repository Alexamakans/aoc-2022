package aoc_test

import (
	"testing"

	"github.com/Alexamakans/aoc-2022/pkg/aoc"
)

func Test_ChunkString(t *testing.T) {
	testCases := []struct {
		value     string
		numChunks int
		want      []string
	}{
		{
			value:     "ab",
			numChunks: 2,
			want:      []string{"a", "b"},
		},
		{
			value:     "abcdef",
			numChunks: 2,
			want:      []string{"abc", "def"},
		},
		{
			value:     "abcdef",
			numChunks: 3,
			want:      []string{"ab", "cd", "ef"},
		},
	}

	for _, tc := range testCases {
		got := aoc.ChunkString(tc.value, tc.numChunks)
		if !aoc.EqualSlice(got, tc.want) {
			t.Fatalf("want %v; got %v (value=%s, numChunks=%d)", tc.want, got, tc.value, tc.numChunks)
		}
	}
}

func Test_ChunkSlice(t *testing.T) {
	testCases := []struct {
		value     []string
		numChunks int
		want      [][]string
	}{
		{
			value:     []string{"abc", "def", "ghi", "jkl"},
			numChunks: 2,
			want: [][]string{
				{"abc", "def"},
				{"ghi", "jkl"},
			},
		},
		{
			value:     []string{"abc", "def", "ghi", "jkl", "mno", "pqr"},
			numChunks: 3,
			want: [][]string{
				{"abc", "def"},
				{"ghi", "jkl"},
				{"mno", "pqr"},
			},
		},
	}

	for _, tc := range testCases {
		got := aoc.ChunkSlice(tc.value, tc.numChunks)
		for k := range got {
			if !aoc.EqualSlice(got[k], tc.want[k]) {
				t.Fatalf("want %v; got %v (value=%s, numChunks=%d)", tc.want[k], got[k], tc.value, tc.numChunks)
			}
		}
	}
}
