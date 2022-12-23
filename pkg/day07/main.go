package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/Alexamakans/aoc-2022/pkg/aoc"
)

func main() {
	aoc.Run(95437, 24933642, partOne, partTwo)
}

func partOne(input string) int {
	directories := getDirectories(input)

	// Find all of the directories with a total size of at most 100000.
	//
	// What is the sum of the total sizes of those directories?

	sum := 0
	for _, dir := range directories {
		size := dir.getSize()
		if size <= 100000 {
			sum += size
		}
	}
	return sum
}

func partTwo(input string) int {
	directories := getDirectories(input)

	// Total space: 70000000
	// Target unused: >= 30000000
	//
	// What is the total size of the smallest directory that we can delete to
	// achieve 30000000 unused space?

	root := directories[0]
	totalSpace := 70000000
	targetUnused := 30000000
	used := root.getSize()
	unused := totalSpace - used
	reqFreeUp := targetUnused - unused

	lowest := math.MaxInt
	for _, dir := range directories {
		size := dir.getSize()
		if size >= reqFreeUp && size < lowest {
			lowest = size
		}
	}
	return lowest
}

func getDirectories(input string) []entry {
	lines := aoc.SplitLines(input)
	var curDir entry = &directory{
		parent: nil,
	}
	directories := []entry{curDir}

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.HasPrefix(line, "$") {
			split := strings.Split(line, " ")
			if split[1] == "ls" {
				for {
					i++
					if i >= len(lines) {
						break
					}

					curLine := lines[i]
					if strings.HasPrefix(curLine, "$") {
						i--
						break
					} else if strings.HasPrefix(curLine, "dir") {
						splt := strings.Split(curLine, " ")
						name := splt[1]
						fmt.Printf("Created directory '%s'\n", name)
						newDir := &directory{
							name: name,
						}
						curDir.addChild(newDir)
						directories = append(directories, newDir)
					} else {
						splt := strings.Split(curLine, " ")
						size, name := aoc.StrToInt(splt[0]), splt[1]
						fmt.Printf("Created file '%s' (%d)\n", name, size)
						curDir.addChild(&file{
							name: name,
							size: size,
						})
					}
				}
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
	return directories
}
