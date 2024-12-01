package day7

import (
	"bufio"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

var rankings = map[rune]int{
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func getBidsByCards(scanner *bufio.Scanner) map[string]int {
	bidsByCards := make(map[string]int)

	for scanner.Scan() {

		text := scanner.Text()
		split := strings.Split(text, " ")

		bid, _ := strconv.Atoi(split[1])
		bidsByCards[split[0]] = bid
	}

	return bidsByCards
}

func groupCardsPart1(cards string) map[rune]int {
	groupedCards := make(map[rune]int)

	for _, r := range []rune(cards) {
		groupedCards[r]++
	}

	return groupedCards
}

func rankCards(groupedCards map[rune]int) []rune {
	r := make([]rune, len(groupedCards))

	i := 0
	for k := range groupedCards {
		r[i] = k
		i++
	}

	sort.Slice(r, func(i, j int) bool {
		if unicode.IsDigit(r[i]) {
			if unicode.IsDigit(r[j]) {
				return int(r[i]-'0') < int(r[j]-'0')
			}

			return true
		}

		return rankings[r[i]] < rankings[r[j]]
	})

	return r
}

func getRank(groupedCards map[rune]int) int {
	if len(groupedCards) == 1 {
		return 6 // five of a kind
	}

	if len(groupedCards) == 5 {
		return 0 // high card
	}

	freq := make(map[int]int)
	for _, v := range groupedCards {
		freq[v]++
	}

	if _, containsQuad := freq[4]; containsQuad {
		return 5 // four of a kind
	}

	if _, containsTriple := freq[3]; containsTriple {
		if _, containsDouble := freq[2]; containsDouble {
			return 4 // three of a kind
		}
		return 3 // full house
	}

	if q, containsDouble := freq[2]; containsDouble {
		if q == 2 {
			return 2 // two pair
		}
		return 1 // one pair
	}

	return 0
}
