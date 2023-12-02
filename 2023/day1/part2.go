package day1

import (
  "bufio"
  "sort"
  "strconv"
  "strings"
  "unicode"
)

var numbers = []string{"_", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func Part2Solution(scanner *bufio.Scanner) int {
  sum := 0

  for scanner.Scan() {
    text := scanner.Text()

    indexes := make(map[int][]int)

    for i, n := range numbers {
      if index := strings.Index(text, n); index != -1 {
        indexes[index] = append(indexes[index], i)
      }
      if index := strings.LastIndex(text, n); index != -1 {
        indexes[index] = append(indexes[index], i)
      }
    }

    for i, b := range text {
      if !unicode.IsDigit(b) {
        continue
      }
      indexes[i] = append(indexes[i], int(b)-'0')
    }

    keys := make([]int, len(indexes))
    i := 0
    for k := range indexes {
      keys[i] = k
      i++
    }

    sort.Slice(keys, func(i, j int) bool {
      return keys[i] < keys[j]
    })

    first := indexes[keys[0]]
    last := indexes[keys[len(keys)-1]]

    combined := strconv.Itoa(first[0]) + strconv.Itoa(last[0])
    v, _ := strconv.Atoi(combined)
    sum += v
  }

  return sum
}
