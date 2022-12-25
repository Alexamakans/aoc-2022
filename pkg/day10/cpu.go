package main

import (
	"strings"

	"github.com/Alexamakans/aoc-2022/pkg/aoc"
)

type instruction interface {
	// returns true if the instruction is finished
	execute(c *cpu) bool
}

type addx struct {
	ticksDone int
}

func (a *addx) execute(c *cpu) bool {
	a.ticksDone++
	if a.ticksDone == 2 {
		c.nextXRegister = c.xRegister + c.argRegister[0]
		a.ticksDone = 0
		return true
	}

	return false
}

type noop struct{}

func (n *noop) execute(c *cpu) bool {
	return true
}

type cpu struct {
	pc            int
	curCycle      int
	code          []string
	argRegister   []int
	xRegister     int
	nextXRegister int
}

func makeCpu(input string) *cpu {
	return &cpu{
		pc:            0,
		curCycle:      0,
		code:          aoc.SplitLines(input),
		argRegister:   make([]int, 1),
		xRegister:     1,
		nextXRegister: 1,
	}
}

var _addx = addx{}
var _noop = noop{}

func (c *cpu) parseInstruction(line string) instruction {
	splt := strings.Split(line, " ")
	ins := splt[0]
	switch ins {
	case "addx":
		c.argRegister[0] = aoc.StrToInt(splt[1])
		return &_addx
	case "noop":
		return &_noop
	}
	return nil
}

func (c *cpu) tick() bool {
	c.xRegister = c.nextXRegister
	ins := c.parseInstruction(c.code[c.pc])
	if ins.execute(c) {
		c.pc++
	}
	c.curCycle++
	return c.pc < len(c.code)
}
