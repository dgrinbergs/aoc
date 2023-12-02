package solver

import (
  "bufio"
  "os"
)

type GetAnswer[T any] func(scanner *bufio.Scanner) T

func Solve[T any](path string, getAnswer GetAnswer[T]) T {
  file, err := os.Open(path)
  if err != nil {
    panic(err)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)
  return getAnswer(scanner)
}
