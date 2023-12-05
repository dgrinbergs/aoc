package day4

import (
  "bufio"
  "strconv"
  "strings"
)

func Part2Solution(scanner *bufio.Scanner) int {
  winningGameIdsByGameId := make(map[int][]int)
  for _, card := range readAllCards(scanner) {
    winningGameIdsByGameId[card.id] = card.getWinningCardIds()
  }

  return recurse(winningGameIdsByGameId, winningGameIdsByGameId) - 1
}

func recurse(all, winners map[int][]int) int {
  if len(winners) == 0 {
    return 1
  }

  sum := 1
  for _, winnerIds := range winners {
    nested := make(map[int][]int)
    for _, i := range winnerIds {
      nested[i] = all[i]
    }

    sum += recurse(all, nested)
  }

  return sum
}

type scratchcard struct {
  id             int
  winningNumbers map[int]bool
  ownNumbers     map[int]bool
}

func (s *scratchcard) getWinningCardIds() []int {
  matches := 0

  for number := range s.ownNumbers {
    if _, ok := s.winningNumbers[number]; ok {
      matches++
    }
  }

  if matches == 0 {
    return []int{}
  }

  winningCardIds := make([]int, matches)
  for i := 0; i < matches; i++ {
    winningCardIds[i] = s.id + i + 1
  }

  return winningCardIds
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
