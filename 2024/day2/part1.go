package day2

import (
	"bufio"
	"math"
	"strconv"
	"strings"
)

func Part1Solution(scanner *bufio.Scanner) int {
	safeCount := 0
	for scanner.Scan() {
		levels := make([]int, 0)
		for _, levelStr := range strings.Split(scanner.Text(), " ") {
			level, _ := strconv.Atoi(levelStr)
			levels = append(levels, level)
		}

		if isValidPart1(levels) {
			safeCount++
		}
	}

	return safeCount
}

func isValidPart1(levels []int) bool {
	sign := math.Signbit(float64(levels[1] - levels[0]))

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		absDiff := abs(diff)

		if absDiff > 3 || absDiff < 1 {
			return false
		}

		if math.Signbit(float64(diff)) != sign {
			return false
		}
	}

	return true
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
