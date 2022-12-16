package main

import (
	"strings"

	"github.com/Alexamakans/aoc-2022/pkg/aoc"
)

func main() {
	aoc.Run(7, 19, partOne, partTwo)
}

func partOne(input string) int {
	lines := aoc.SplitLines(input)
	var curDir entry = &directory{
		parent: nil,
	}
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.HasPrefix(line, "$") {
			split := strings.Split(line, " ")
			if split[1] == "ls" {
				// TODO: Parse lines until one starts with $ and create the
				//       files/directories read there.
			} else if split[1] == "cd" {
				nextDir := split[2]
				if nextDir == "/" {
					for p := curDir.getParent(); p != nil; p = curDir.getParent() {
						curDir = curDir.getParent()
					}
				} else if nextDir == ".." {
					curDir = curDir.getParent()
				} else {
					for _, child := range curDir.getEntries() {
						if child.getName() == nextDir {
							curDir = child
						}
					}
				}
			}
		}
	}
}

func partTwo(input string) int {
	return -1
}
