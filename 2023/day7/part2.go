package day7

import (
	"bufio"
	"sort"
)

func Part2Solution(scanner *bufio.Scanner) int {
	bidsByCards := getBidsByCards(scanner)

	cards := make([]string, len(bidsByCards))
	i := 0
	for card := range bidsByCards {
		cards[i] = card
		i++
	}

	sort.Slice(cards, func(i, j int) bool {
		return sortCards(cards[i], cards[j], groupCardsPart2)
	})

	total := 0
	for cardRank, card := range cards {
		total += bidsByCards[card] * (cardRank + 1)
	}

	return total
}

func groupCardsPart2(cards string) map[rune]int {
	groupedCards := make(map[rune]int)

	for _, r := range []rune(cards) {
		groupedCards[r]++
	}

	return groupedCards
}
