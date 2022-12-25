package main

import (
	"fmt"
	"math"
)

type crt struct {
	width, height int
	pixels        [][]byte
}

func makeCrt() *crt {
	w := 40
	h := 6
	pixels := make([][]byte, 0)
	for y := 0; y < h; y++ {
		pixels = append(pixels, make([]byte, 0))
		for x := 0; x < w; x++ {
			pixels[y] = append(pixels[y], '.')
		}
	}
	return &crt{
		width:  w,
		height: h,
		pixels: pixels,
	}
}

func (c *crt) draw(x, cycle int) {
	xPos := (cycle - 1) % c.width
	yPos := (cycle - 1) / c.width
	if int(math.Abs(float64(x-xPos))) <= 1 {
		c.pixels[yPos][xPos] = '#'
	}
}

func (c *crt) print() {
	for _, line := range c.pixels {
		fmt.Println(string(line))
	}
}
