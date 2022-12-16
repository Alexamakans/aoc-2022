package main

import (
	"github.com/Alexamakans/aoc-2022/pkg/aoc"
)

func main() {
	aoc.Run(7, 19, partOne, partTwo)
}

func getFirstMarkerEndIndex(s string, packetLength int) int {
	for windowStart := range s {
		windowEnd := windowStart + packetLength
		for i := 0; i < packetLength; i++ {
			beforeCurrentEnd := windowStart + i
			afterCurrentStart := windowStart + i + 1
			beforeCurrent := []rune(s[windowStart:beforeCurrentEnd])
			afterCurrent := []rune(s[afterCurrentStart:windowEnd])
			r := rune(s[windowStart+i])
			if aoc.Contains(beforeCurrent, r) || aoc.Contains(afterCurrent, r) {
				break
			}

			if i == packetLength-1 {
				return windowStart + i + 1
			}
		}
	}

	// none found, shouldn't happen
	return -1
}

func partOne(input string) int {
	return getFirstMarkerEndIndex(input, 4)
}

func partTwo(input string) int {
	return getFirstMarkerEndIndex(input, 14)
}
