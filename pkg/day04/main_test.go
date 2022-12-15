package main

import "testing"

func Test_PairOneIsFullyContained(t *testing.T) {
	testCases := []struct {
		value pair
		want  bool
	}{
		{
			value: pair{section{6, 6}, section{5, 7}},
			want:  true,
		},
		{
			value: pair{section{3, 6}, section{5, 7}},
			want:  false,
		},
	}

	for _, tc := range testCases {
		got := tc.value.isOneFullyContained()
		if got != tc.want {
			t.Fatalf("got %t; want %t (value=%v)", got, tc.want, tc.value)
		}
	}
}

func Test_PairIsOverlapping(t *testing.T) {
	testCases := []struct {
		value pair
		want  bool
	}{
		{
			value: pair{section{2, 4}, section{6, 8}},
			want:  false,
		},
		{
			value: pair{section{2, 3}, section{4, 5}},
			want:  false,
		},
		{
			value: pair{section{5, 7}, section{7, 9}},
			want:  true,
		},
		{
			value: pair{section{2, 8}, section{3, 7}},
			want:  true,
		},
		{
			value: pair{section{6, 6}, section{4, 6}},
			want:  true,
		},
		{
			value: pair{section{2, 6}, section{4, 8}},
			want:  true,
		},
	}

	for _, tc := range testCases {
		got := tc.value.isOverlapping()
		if got != tc.want {
			t.Fatalf("got %t; want %t (value=%v)", got, tc.want, tc.value)
		}
	}
}
