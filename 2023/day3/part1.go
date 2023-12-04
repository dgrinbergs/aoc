package day3

import (
	"bufio"
	"strconv"
	"unicode"
)

func Part1Solution(scanner *bufio.Scanner) int {
	numbers, symbols, lineLength := getNumberAndSymbolIndexesAndLineLength(scanner)

	sum := 0

	for index, number := range numbers {
		for _, i := range getSymbolIndexes(index, number, lineLength) {
			if _, ok := symbols[i]; ok {
				sum += number
				break
			}
		}
	}

	return sum
}

func getIndex(row, lineLength, index int) int {
	return (row * lineLength) + index
}

func getSymbolIndexes(index, number, lineLength int) []int {
	indexes := make(map[int]bool)
	length := len(strconv.Itoa(number))

	for i := 0; i < length; i++ {
		trueIndex := index + i

		indexes[trueIndex-lineLength] = true // n
		indexes[trueIndex+lineLength] = true // s

		if trueIndex%lineLength != 0 {
			indexes[trueIndex-1-lineLength] = true // nw
			indexes[trueIndex-1] = true            // w
			indexes[trueIndex-1+lineLength] = true // sw
		}

		if trueIndex%lineLength != lineLength-1 {
			indexes[trueIndex+1-lineLength] = true // ne
			indexes[trueIndex+1] = true            // e
			indexes[trueIndex+1+lineLength] = true // se
		}
	}

	uniqueIndexes := make([]int, len(indexes))
	i := 0
	for k, _ := range indexes {
		uniqueIndexes[i] = k
		i++
	}

	return uniqueIndexes
}

func getNumberAndSymbolIndexesAndLineLength(scanner *bufio.Scanner) (map[int]int, map[int]rune, int) {
	symbols := make(map[int]rune)
	numbers := make(map[int]int)
	var lineLength int

	row := 0
	for scanner.Scan() {
		runes := []rune(scanner.Text())
		lineLength = len(runes)

		for i := 0; i < lineLength; i++ {
			r := runes[i]
			if r == '.' {
				continue
			}

			if !unicode.IsNumber(r) {
				symbols[getIndex(row, lineLength, i)] = r
			}

			if _, ok := symbols[getIndex(row, lineLength, i)]; ok {
				continue
			}

			number := make([]rune, 0)
			for j := i; j < lineLength; j++ {
				if unicode.IsNumber(runes[j]) {
					number = append(number, runes[j])

					if j == lineLength-1 {
						n, _ := strconv.Atoi(string(number))
						numbers[getIndex(row, lineLength, i)] = n
					}

					continue
				}

				n, _ := strconv.Atoi(string(number))
				numbers[getIndex(row, lineLength, i)] = n

				i = j - 1
				break
			}
		}

		row++
	}

	return numbers, symbols, lineLength
}
