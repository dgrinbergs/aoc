package day3

import (
	"bufio"
	"strconv"
)

func Part2Solution(scanner *bufio.Scanner) int {
	numbers, symbols, lineLength := getNumberAndSymbolIndexesAndLineLength(scanner)
	sum := 0

	for index, symbol := range symbols {
		if symbol != '*' {
			continue
		}

		gears := make(map[int]int)
		for _, adjacentIndex := range getAdjacentSymbolIndexes(index, lineLength) {
			for numberIndex, number := range numbers {
				numberLength := len(strconv.Itoa(number))
				for i := 0; i < numberLength; i++ {
					if adjacentIndex == numberIndex+i {
						gears[numberIndex] = number
					}
				}
			}
		}
		if len(gears) == 2 {
			ratio := 1
			for _, gear := range gears {
				ratio *= gear
			}
			sum += ratio
		}
	}

	return sum
}

func getAdjacentSymbolIndexes(index, lineLength int) []int {
	indexes := make(map[int]bool)

	indexes[index-lineLength] = true // n
	indexes[index+lineLength] = true // s

	if index%lineLength != 0 {
		indexes[index-1-lineLength] = true // nw
		indexes[index-1] = true            // w
		indexes[index-1+lineLength] = true // sw
	}

	if index%lineLength != lineLength-1 {
		indexes[index+1-lineLength] = true // ne
		indexes[index+1] = true            // e
		indexes[index+1+lineLength] = true // se
	}

	uniqueIndexes := make([]int, len(indexes))
	i := 0
	for k, _ := range indexes {
		uniqueIndexes[i] = k
		i++
	}

	return uniqueIndexes
}
