package day3

import (
	"bufio"
	"regexp"
)

func Part2Solution(scanner *bufio.Scanner) int {
	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	sum := 0
	enabled := true

	for scanner.Scan() {
		text := scanner.Text()

		for _, group := range re.FindAllString(text, -1) {
			if group == "don't()" {
				enabled = false
			} else if group == "do()" {
				enabled = true
			} else if enabled {
				sum += multiply(group)
			}
		}
	}

	return sum
}
