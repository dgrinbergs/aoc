package day5

import (
	"bufio"
	"strconv"
	"strings"
)

func Part1Solution(scanner *bufio.Scanner) int {
	rules := make(map[int][]int)
	numbers := make([][]int, 0)

	isRules := true
	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			isRules = false
			continue
		}

		if isRules {
			parts := strings.Split(text, "|")
			a, _ := strconv.Atoi(parts[0])
			b, _ := strconv.Atoi(parts[1])

			rules[a] = append(rules[a], b)
		} else {
			group := make([]int, 0)
			for _, part := range strings.Split(text, ",") {
				number, _ := strconv.Atoi(part)
				group = append(group, number)
			}
			numbers = append(numbers, group)
		}
	}

	result := 0

	for _, validIndex := range getValidIndexes(numbers, rules) {
		mp := len(numbers[validIndex]) / 2
		result += numbers[validIndex][mp]
	}

	return result
}

func getValidIndexes(numbers [][]int, rules map[int][]int) []int {
	validIndexes := make([]int, 0)

groupLoop:
	for i, group := range numbers {
		for j := 1; j < len(group); j++ {
			curr := group[j]

			for k := range rules {
				if curr != k {
					continue
				}

				for _, constraint := range rules[curr] {
					for _, prev := range group[:j] {
						if prev == constraint {
							continue groupLoop
						}
					}
				}
			}

		}
		validIndexes = append(validIndexes, i)
	}

	return validIndexes
}
