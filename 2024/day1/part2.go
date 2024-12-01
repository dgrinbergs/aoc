package day1

import (
	"bufio"
	"strconv"
)

func Part2Solution(scanner *bufio.Scanner) int {
	leftLocations := make([]int, 0)
	rightFreq := make(map[int]int)

	for scanner.Scan() {
		text := scanner.Text()

		left, _ := strconv.Atoi(text[0:5])
		leftLocations = append(leftLocations, left)

		right, _ := strconv.Atoi(text[8:])
		rightFreq[right]++
	}

	score := 0
	for _, left := range leftLocations {
		score += left * rightFreq[left]
	}

	return score
}
