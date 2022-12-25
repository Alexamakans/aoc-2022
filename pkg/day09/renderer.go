package main

import (
	"fmt"

	"github.com/Alexamakans/aoc-2022/pkg/aoc"
)

type renderer struct {
	minX, maxX int
	minY, maxY int
}

func (r *renderer) render(rope rope) {
	positions := []*aoc.Position{rope.head}
	positions = append(positions, rope.tail...)

	for _, p := range positions {
		if p.X <= r.minX {
			r.minX = p.X - 4
		}
		if p.X >= r.maxX {
			r.maxX = p.X + 4
		}
		if p.Y <= r.minY {
			r.minY = p.Y - 4
		}
		if p.Y >= r.maxY {
			r.maxY = p.Y + 4
		}
	}

	var lines []string
	for y := r.minY; y < r.maxY; y++ {
		line := ""
		for x := r.minX; x < r.maxX; x++ {
			chr := "."
			for idx, p := range positions {
				if p.X == x && p.Y == y {
					if idx == 0 {
						chr = "H"
					} else {
						chr = aoc.IntToStr(idx)
						break
					}
				}
			}
			line += chr
		}
		lines = append(lines, line)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}
