package aoc

type Position struct {
	X int
	Y int
}

func (p *Position) Move(dx, dy int) {
	p.X += dx
	p.Y += dy
}
