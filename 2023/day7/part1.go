package day7

import (
	"bufio"
	"sort"
	"unicode"
)

func Part1Solution(scanner *bufio.Scanner) int {
	bidsByCards := getBidsByCards(scanner)

	cards := make([]string, len(bidsByCards))
	i := 0
	for card := range bidsByCards {
		cards[i] = card
		i++
	}

	sort.Slice(cards, func(i, j int) bool {
		return sortCards(cards[i], cards[j], groupCardsPart1)
	})

	total := 0
	for cardRank, card := range cards {
		total += bidsByCards[card] * (cardRank + 1)
	}

	return total
}

func sortCards(a, b string, groupFunc func(string) map[rune]int) bool {
	cardsA := groupFunc(a)
	cardsB := groupFunc(b)

	classA := getRank(cardsA)
	classB := getRank(cardsB)

	if classA != classB {
		return classA < classB
	}

	for index, cardA := range a {
		cardB := int32(b[index])

		if unicode.IsDigit(cardA) {
			if unicode.IsDigit(cardB) {
				if cardA == cardB {
					continue
				}

				return cardA < cardB
			}

			return int(cardA-'0') < rankings[cardB]
		}

		if rankings[cardA] == rankings[cardB] {
			continue
		}

		return rankings[cardA] < rankings[cardB]
	}

	return true
}
