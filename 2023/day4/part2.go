package day4

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Part2Solution(scanner *bufio.Scanner) int {
	allCards := readAllCards(scanner)
	sum := 0

	for _, card := range allCards {
		sum += card.getNumberOfNestedMatches(allCards)
	}

	return sum
}

type scratchcard struct {
	id             int
	winningNumbers map[int]bool
	ownNumbers     map[int]bool
}

func (s *scratchcard) getNumberOfNestedMatches(all []*scratchcard) int {
	matches := 0

	for number := range s.ownNumbers {
		if _, ok := s.winningNumbers[number]; ok {
			matches++
		}
	}

	for i := 0; i < matches; i++ {
		fmt.Printf("%d+%d=%d [%d]\n", s.id, i, s.id+i, matches)
		matches += all[s.id+i].getNumberOfNestedMatches(all)
	}

	return matches
}

func readAllCards(scanner *bufio.Scanner) []*scratchcard {
	scratchCards := make([]*scratchcard, 0)

	for scanner.Scan() {
		card := scanner.Text()

		colonIndex := strings.Index(card, ":")
		verticalBarIndex := strings.Index(card, "|")

		winning := make(map[int]bool)
		for _, number := range strings.Split(card[colonIndex:verticalBarIndex], " ") {
			if n, err := strconv.Atoi(number); err == nil {
				winning[n] = true
			}
		}

		own := make(map[int]bool)
		for _, number := range strings.Split(card[verticalBarIndex:], " ") {
			if n, err := strconv.Atoi(number); err == nil {
				own[n] = true
			}
		}

		id, _ := strconv.Atoi(strings.TrimSpace(card[5:colonIndex]))
		scratchCards = append(scratchCards, &scratchcard{id, winning, own})
	}

	return scratchCards
}
