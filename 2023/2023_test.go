package _023

import (
	solver "github.com/dgrinbergs/aoc"
	"github.com/dgrinbergs/aoc/2023/day1"
	"github.com/dgrinbergs/aoc/2023/day2"
	"github.com/dgrinbergs/aoc/2023/day3"
	"github.com/dgrinbergs/aoc/2023/day4"
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

func TestDay3(t *testing.T) {
	assert.Equal(t, 4361, solver.Solve("day3/example.txt", day3.Part1Solution))
	assert.Equal(t, 539433, solver.Solve("day3/input.txt", day3.Part1Solution))
	assert.Equal(t, 467835, solver.Solve("day3/example.txt", day3.Part2Solution))
	assert.Equal(t, 75847567, solver.Solve("day3/input.txt", day3.Part2Solution))
}

func TestDay4(t *testing.T) {
	assert.Equal(t, 13, solver.Solve("day4/example.txt", day4.Part1Solution))
	assert.Equal(t, 28538, solver.Solve("day4/input.txt", day4.Part1Solution))
	assert.Equal(t, 30, solver.Solve("day4/example.txt", day4.Part2Solution))
}
