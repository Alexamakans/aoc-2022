package main

import "fmt"

type world struct {
	marks    map[string]bool
	numMarks int
}

func (w *world) mark(x, y int) {
	idx := fmt.Sprintf("%d:%d", x, y)
	if !w.marks[idx] {
		w.numMarks++
		w.marks[idx] = true
	}
}
