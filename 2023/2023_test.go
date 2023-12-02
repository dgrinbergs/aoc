package _023

import (
  solver "github.com/dgrinbergs/aoc"
  "github.com/dgrinbergs/aoc/2023/day1"
  "github.com/dgrinbergs/aoc/2023/day2"
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestDay1(t *testing.T) {
  assert.Equal(t, 55477, solver.Solve("day1/input.txt", day1.Part1Solution))
  assert.Equal(t, 54431, solver.Solve("day1/input.txt", day1.Part2Solution))
}

func TestDay2(t *testing.T) {
  assert.Equal(t, 2551, solver.Solve("day2/input.txt", day2.Part1Solution))
  assert.Equal(t, 62811, solver.Solve("day2/input.txt", day2.Part2Solution))
}
