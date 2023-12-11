package day6

import (
	"bufio"
	"math"
	"strconv"
	"strings"
)

func getDistanceByTime(scanner *bufio.Scanner) map[int]int {
	distanceByTime := make(map[int]int)

	scanner.Scan()
	times := strings.Fields(scanner.Text())[1:]

	scanner.Scan()
	distances := strings.Fields(scanner.Text())[1:]

	for i := 0; i < len(times); i++ {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])

		distanceByTime[time] = distance
	}

	return distanceByTime
}

func solveQuadratic(a, b, c int) (float64, float64) {
	sqrt := math.Sqrt(float64(b*b - (4 * a * c)))
	bottom := float64(2 * a)
	return (-float64(b) - sqrt) / bottom, (-float64(b) + sqrt) / bottom
}

func getRanges(first, second float64) int {
	if first >= math.Trunc(first) {
		first += 1
	} else {
		first = math.Ceil(first)
	}

	if second >= math.Trunc(second) {
		second -= 1
	} else {
		second = math.Floor(second)
	}

	return int(second-first) + 1
}
