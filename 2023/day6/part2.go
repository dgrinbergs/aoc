package day6

import (
	"bufio"
	"strconv"
	"strings"
)

func Part2Solution(scanner *bufio.Scanner) int {
	time, distance := getDistanceAndTime(scanner)
	return getRanges(solveQuadratic(1, time, distance))
}

func getDistanceAndTime(scanner *bufio.Scanner) (int, int) {
	scanner.Scan()
	timeStr := strings.TrimSpace(scanner.Text()[5:])

	scanner.Scan()
	distanceStr := strings.TrimSpace(scanner.Text())[10:]

	time, _ := strconv.Atoi(strings.ReplaceAll(timeStr, " ", ""))
	distance, _ := strconv.Atoi(strings.ReplaceAll(distanceStr, " ", ""))

	return time, distance
}
