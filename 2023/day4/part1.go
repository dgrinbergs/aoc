package day4

import (
	"bufio"
	"strconv"
	"strings"
)

func Part1Solution(scanner *bufio.Scanner) int {
	sum := 0

	for scanner.Scan() {
		card := scanner.Text()
		winningNumbers, ownNumbers := getWinningNumbersAndOwnNumbers(card)
		matches := 0

		for number := range ownNumbers {
			if _, ok := winningNumbers[number]; ok {
				matches++
			}
		}

		sum += calculatePoints(matches)
	}

	return sum
}

func getWinningNumbersAndOwnNumbers(card string) (map[int]bool, map[int]bool) {
	winningNumbers := make(map[int]bool)
	ownNumbers := make(map[int]bool)

	colonIndex := strings.Index(card, ":")
	verticalBarIndex := strings.Index(card, "|")

	for _, number := range strings.Split(card[colonIndex:verticalBarIndex], " ") {
		if n, err := strconv.Atoi(number); err == nil {
			winningNumbers[n] = true
		}
	}

	for _, number := range strings.Split(card[verticalBarIndex:], " ") {
		if n, err := strconv.Atoi(number); err == nil {
			ownNumbers[n] = true
		}
	}

	return winningNumbers, ownNumbers
}

func calculatePoints(matches int) int {
	if matches == 0 {
		return 0
	}

	sum := 1

	for i := 1; i < matches; i++ {
		sum *= 2
	}

	return sum
}
