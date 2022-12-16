package main

import (
	"github.com/Alexamakans/aoc-2022/pkg/aoc"
)

func main() {
	aoc.Run(7, 19, partOne, partTwo)
}

func getFirstMarkerEndIndex(s string, packetLength int) int {
	for windowStart := range s {
		for i := 0; i < packetLength; i++ {
			slice := append([]rune(s[windowStart:windowStart+i]), []rune(s[windowStart+i+1:windowStart+packetLength])...)
			r := rune(s[windowStart+i])
			if aoc.Contains(slice, r) {
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
