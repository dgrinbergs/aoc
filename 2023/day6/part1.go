package day6

import (
	"bufio"
)

func Part1Solution(scanner *bufio.Scanner) int {
	total := 1

	for time, distance := range getDistanceByTime(scanner) {
		first, second := solveQuadratic(1, time, distance)
		total *= getRanges(first, second)
	}

	return total
}
