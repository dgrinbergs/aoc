package day4

import (
	"bufio"
)

func Part2Solution(scanner *bufio.Scanner) int {
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

	for i := 1; i < len(wordSearch)-1; i++ {
		for j := 1; j < len(wordSearch[i])-1; j++ {
			current := wordSearch[i][j]
			if current != 'A' {
				continue
			}

			northEast := wordSearch[i-1][j+1]
			northWest := wordSearch[i-1][j-1]
			southEast := wordSearch[i+1][j+1]
			southWest := wordSearch[i+1][j-1]

			if ((northWest == 'M' && southEast == 'S') || (northWest == 'S' && southEast == 'M')) &&
				((northEast == 'M' && southWest == 'S') || (northEast == 'S' && southWest == 'M')) {
				found++
			}
		}
	}

	return found
}
