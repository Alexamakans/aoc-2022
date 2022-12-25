package main

import (
	"github.com/Alexamakans/aoc-2022/pkg/aoc"
)

func main() {
	// Second part prints a string containing capital letters rendered using # and .s
	aoc.Run(13140, -1, partOne, partTwo)
}

func partOne(input string) int {
	cpu := makeCpu(input)

	signal := 0
	for cpu.tick() {
		if cpu.curCycle == 20 || (cpu.curCycle-20)%40 == 0 {
			signal += cpu.curCycle * cpu.xRegister
		}
	}

	return signal
}

func partTwo(input string) int {
	cpu := makeCpu(input)
	crt := makeCrt()
	for cpu.tick() {
		crt.draw(cpu.xRegister, cpu.curCycle)
	}
	crt.print()
	return -1
}
