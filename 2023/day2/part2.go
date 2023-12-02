package day2

import (
  "bufio"
)

func Part2Solution(scanner *bufio.Scanner) int {
  sum := 0

  for scanner.Scan() {
    game := scanner.Text()

    pow := 1
    for _, v := range GetTurns(game) {
      pow *= v
    }

    sum += pow
  }

  return sum
}
