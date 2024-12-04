package day3

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

func Part1Solution(scanner *bufio.Scanner) int {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	sum := 0
	for scanner.Scan() {
		for _, match := range re.FindAllString(scanner.Text(), -1) {
			sum += multiply(match)
		}
	}

	return sum
}

func multiply(mul string) int {
	comma := strings.Index(mul, ",")
	first, _ := strconv.Atoi(mul[4:comma])
	second, _ := strconv.Atoi(mul[comma+1 : len(mul)-1])

	return first * second
}
