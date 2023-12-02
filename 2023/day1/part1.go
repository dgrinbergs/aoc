package day1

import (
  "bufio"
  "strconv"
  "unicode"
)

func Part1Solution(scanner *bufio.Scanner) int {
  sum := 0

  for scanner.Scan() {
    text := scanner.Text()

    digits := make([]string, 0)

    for _, b := range text {
      if unicode.IsNumber(b) {
        digits = append(digits, string(b))
      }
    }

    number, _ := strconv.Atoi(digits[0] + digits[len(digits)-1])
    sum += number
  }

  return sum
}
