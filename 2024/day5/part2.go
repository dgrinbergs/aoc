package day5

import (
	"bufio"
	"slices"
	"strconv"
	"strings"
)

func Part2Solution(scanner *bufio.Scanner) int {
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

	validIndexes := make(map[int]bool)
	for _, i := range getValidIndexes(numbers, rules) {
		validIndexes[i] = true
	}

	result := 0
	for i, group := range numbers {
		if ok := validIndexes[i]; ok {
			continue
		}

		slices.SortFunc(group, func(a, b int) int {
			for _, r := range rules[a] {
				if r == b {
					return -1
				}
			}

			return 0
		})

		result += group[len(group)/2]
	}

	return result
}
