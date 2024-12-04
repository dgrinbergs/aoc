package day4

import (
	"bufio"
)

func Part1Solution(scanner *bufio.Scanner) int {
	wordSearch := make([][]rune, 0)

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, len(line))
		for i, ch := range line {
			row[i] = ch
		}
		wordSearch = append(wordSearch, row)
	}

	found := 0
	word := []rune{'X', 'M', 'A', 'S'}

	for i := 0; i < len(wordSearch); i++ {
		for j := 0; j < len(wordSearch[i]); j++ {
			if hasNextLetter(wordSearch, i, j, directionNorth, word, 0) {
				found++
			}
			if hasNextLetter(wordSearch, i, j, directionNorthEast, word, 0) {
				found++
			}
			if hasNextLetter(wordSearch, i, j, directionEast, word, 0) {
				found++
			}
			if hasNextLetter(wordSearch, i, j, directionSouthEast, word, 0) {
				found++
			}
			if hasNextLetter(wordSearch, i, j, directionSouth, word, 0) {
				found++
			}
			if hasNextLetter(wordSearch, i, j, directionSouthWest, word, 0) {
				found++
			}
			if hasNextLetter(wordSearch, i, j, directionWest, word, 0) {
				found++
			}
			if hasNextLetter(wordSearch, i, j, directionNorthWest, word, 0) {
				found++
			}
		}
	}

	return found
}

type direction string

const (
	directionNorth     = direction("N")
	directionNorthEast = direction("NE")
	directionEast      = direction("E")
	directionSouthEast = direction("SE")
	directionSouth     = direction("S")
	directionSouthWest = direction("SW")
	directionWest      = direction("W")
	directionNorthWest = direction("NW")
)

func hasNextLetter(wordSearch [][]rune, i, j int, direction direction, word []rune, index int) bool {
	if wordSearch[i][j] != word[index] {
		return false
	}

	north := i > 0
	east := j < len(wordSearch[i])-1
	south := i < len(wordSearch[i])-1
	west := j > 0

	letter := word[index]

	if index == len(word)-1 {
		if wordSearch[i][j] == letter {
			return true
		}
		return false
	}

	switch direction {
	case directionNorth:
		if north {
			return hasNextLetter(wordSearch, i-1, j, direction, word, index+1)
		}
	case directionNorthEast:
		if north && east {
			return hasNextLetter(wordSearch, i-1, j+1, direction, word, index+1)
		}
	case directionEast:
		if east {
			return hasNextLetter(wordSearch, i, j+1, direction, word, index+1)
		}
	case directionSouthEast:
		if south && east {
			return hasNextLetter(wordSearch, i+1, j+1, direction, word, index+1)
		}
	case directionSouth:
		if south {
			return hasNextLetter(wordSearch, i+1, j, direction, word, index+1)
		}
	case directionSouthWest:
		if south && west {
			return hasNextLetter(wordSearch, i+1, j-1, direction, word, index+1)
		}
	case directionWest:
		if west {
			return hasNextLetter(wordSearch, i, j-1, direction, word, index+1)
		}
	case directionNorthWest:
		if north && west {
			return hasNextLetter(wordSearch, i-1, j-1, direction, word, index+1)
		}
	}

	return false
}
