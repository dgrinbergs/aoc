package day5

import (
	"bufio"
	"fmt"
)

func Part2Solution(scanner *bufio.Scanner) int {
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return 0
}
