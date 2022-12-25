package main

import "github.com/Alexamakans/aoc-2022/pkg/aoc"

type rope struct {
	head *aoc.Position
	tail []*aoc.Position
}

func (r *rope) move(dir rune, follow bool) {
	switch dir {
	case 'U':
		r.head.Y -= 1
	case 'R':
		r.head.X += 1
	case 'D':
		r.head.Y += 1
	case 'L':
		r.head.X -= 1
	}
	if follow {
		r.follow()
	}
}

func (r *rope) follow() {
	r.followPart(r.head, r.tail[0])
	for i := 0; i < len(r.tail)-1; i++ {
		followee := r.tail[i]
		follower := r.tail[i+1]
		r.followPart(followee, follower)
	}
}

func (r *rope) followPart(followee, follower *aoc.Position) {
	dx := followee.X - follower.X
	dy := followee.Y - follower.Y
	clampedDX := aoc.Clamp(dx, -1, 1)
	clampedDY := aoc.Clamp(dy, -1, 1)
	if dx == clampedDX && dy == clampedDY {
		return
	}
	follower.Move(clampedDX, clampedDY)
}
