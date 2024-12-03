package day2

import (
	"bufio"
	"math"
	"strconv"
	"strings"
)

func Part2Solution(scanner *bufio.Scanner) int {
	safeCount := 0
	for scanner.Scan() {
		levels := make([]int, 0)
		for _, levelStr := range strings.Split(scanner.Text(), " ") {
			level, _ := strconv.Atoi(levelStr)
			levels = append(levels, level)
		}

		if isValidPart2(levels) {
			safeCount++
		} else {
			for i := range len(levels) {
				dampened := make([]int, 0)
				dampened = append(dampened, levels[:i]...)
				dampened = append(dampened, levels[i+1:]...)
				if isValidPart2(dampened) {
					safeCount++
					break
				}
			}
		}
	}

	return safeCount
}

func isValidPart2(levels []int) bool {
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
