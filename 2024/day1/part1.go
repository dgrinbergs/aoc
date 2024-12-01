package day1

import (
	"bufio"
	"slices"
	"strconv"
)

func Part1Solution(scanner *bufio.Scanner) int {
	leftLocations := make([]int, 0)
	rightLocations := make([]int, 0)

	for scanner.Scan() {
		text := scanner.Text()

		left, _ := strconv.Atoi(text[0:5])
		leftLocations = append(leftLocations, left)

		right, _ := strconv.Atoi(text[8:])
		rightLocations = append(rightLocations, right)
	}

	slices.Sort(leftLocations)
	slices.Sort(rightLocations)

	distances := 0
	for i := 0; i < len(leftLocations); i++ {
		gap := leftLocations[i] - rightLocations[i]
		if gap < 0 {
			distances -= gap
			continue
		}
		distances += gap
	}

	return distances
}
