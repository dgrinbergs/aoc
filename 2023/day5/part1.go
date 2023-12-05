package day5

import (
	"bufio"
	"math"
	"strconv"
	"strings"
)

func Part1Solution(scanner *bufio.Scanner) int {
	locations := getSeedsPart1(scanner)
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

func getSeedsPart1(scanner *bufio.Scanner) map[int]int {
	if !scanner.Scan() {
		panic(scanner.Err())
	}

	seeds := make(map[int]int)

	for _, str := range strings.Split(scanner.Text()[7:], " ") {
		seed, _ := strconv.Atoi(str)
		seeds[seed] = 1
	}

	return seeds
}
