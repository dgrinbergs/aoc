package day5

import (
	"bufio"
	"math"
	"strconv"
	"strings"
)

func Part2Solution(scanner *bufio.Scanner) int {
	locations := getSeedsPart2(scanner)
	groupedBlockMappings := getSectionsWithBlocks(scanner)

	mapLocations(locations, groupedBlockMappings)

	minimum := math.MaxInt
	for seed, _ := range locations {
		if seed < minimum {
			minimum = seed
		}
	}

	return minimum
}

func getSeedsPart2(scanner *bufio.Scanner) map[int]int {
	if !scanner.Scan() {
		panic(scanner.Err())
	}

	seeds := make(map[int]int)
	text := scanner.Text()[7:]

	split := strings.Split(text, " ")

	for i := 0; i < len(split)-1; i += 2 {
		start, _ := strconv.Atoi(split[i])
		width, _ := strconv.Atoi(split[i+1])
		seeds[start] = width
	}

	return seeds
}
