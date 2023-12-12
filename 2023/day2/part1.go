package day2

import (
	"bufio"
	"strconv"
	"strings"
)

func Part1Solution(scanner *bufio.Scanner) int {
	limit := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	sum := 0

games:
	for scanner.Scan() {
		game := scanner.Text()

		id := getGameId(game)
		turns := GetTurns(game)

		for k, v := range turns {
			if limit[k] < v {
				continue games
			}
		}

		sum += id
	}

	return sum
}

func getGameId(game string) int {
	colon := strings.Index(game, ":")
	id, _ := strconv.Atoi(game[5:colon])
	return id
}

func GetTurns(game string) map[string]int {
	colon := strings.Index(game, ":")
	game = game[colon+1:]

	seen := make(map[string]int)

	for _, turn := range strings.Split(game, ";") {
		for _, pair := range strings.Split(turn, ",") {
			pair = strings.TrimSpace(pair)
			space := strings.Index(pair, " ")

			colour := pair[space+1:]
			count, _ := strconv.Atoi(pair[0:space])

			if _, ok := seen[colour]; !ok {
				seen[colour] = count
				continue
			}

			if seen[colour] > count {
				continue
			}

			seen[colour] = count
		}
	}

	return seen
}
