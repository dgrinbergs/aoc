package _024

import (
	solver "github.com/dgrinbergs/aoc"
	"github.com/dgrinbergs/aoc/2024/day1"
	"github.com/dgrinbergs/aoc/2024/day2"
	"github.com/dgrinbergs/aoc/2024/day3"
	"github.com/dgrinbergs/aoc/2024/day4"
	"github.com/dgrinbergs/aoc/2024/day5"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay1(t *testing.T) {
	assert.Equal(t, 3508942, solver.Solve("day1/input.txt", day1.Part1Solution))
	assert.Equal(t, 26593248, solver.Solve("day1/input.txt", day1.Part2Solution))
}

func TestDay2(t *testing.T) {
	assert.Equal(t, 2, solver.Solve("day2/example.txt", day2.Part1Solution))
	assert.Equal(t, 252, solver.Solve("day2/input.txt", day2.Part1Solution))
	assert.Equal(t, 4, solver.Solve("day2/example.txt", day2.Part2Solution))
	assert.Equal(t, 324, solver.Solve("day2/input.txt", day2.Part2Solution))
}

func TestDay3(t *testing.T) {
	assert.Equal(t, 182780583, solver.Solve("day3/input.txt", day3.Part1Solution))
	assert.Equal(t, 48, solver.Solve("day3/example.txt", day3.Part2Solution))
	assert.Equal(t, 90772405, solver.Solve("day3/input.txt", day3.Part2Solution))
}

func TestDay4(t *testing.T) {
	assert.Equal(t, 18, solver.Solve("day4/example.txt", day4.Part1Solution))
	assert.Equal(t, 2646, solver.Solve("day4/input.txt", day4.Part1Solution))
	assert.Equal(t, 9, solver.Solve("day4/example.txt", day4.Part2Solution))
	assert.Equal(t, 2000, solver.Solve("day4/input.txt", day4.Part2Solution))
}

func TestDay5(t *testing.T) {
	assert.Equal(t, 143, solver.Solve("day5/example.txt", day5.Part1Solution))
	assert.Equal(t, 6242, solver.Solve("day5/input.txt", day5.Part1Solution))
	assert.Equal(t, 123, solver.Solve("day5/example.txt", day5.Part2Solution))
	assert.Equal(t, 5169, solver.Solve("day5/input.txt", day5.Part2Solution))
}
