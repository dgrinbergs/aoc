package _023

import (
	solver "github.com/dgrinbergs/aoc"
	"github.com/dgrinbergs/aoc/2023/day1"
	"github.com/dgrinbergs/aoc/2023/day2"
	"github.com/dgrinbergs/aoc/2023/day3"
	"github.com/dgrinbergs/aoc/2023/day4"
	"github.com/dgrinbergs/aoc/2023/day5"
	"github.com/dgrinbergs/aoc/2023/day6"
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
	assert.Equal(t, 9425061, solver.Solve("day4/input.txt", day4.Part2Solution))
}

func TestDay5(t *testing.T) {
	assert.Equal(t, 35, solver.Solve("day5/example.txt", day5.Part1Solution))
	assert.Equal(t, 486613012, solver.Solve("day5/input.txt", day5.Part1Solution))
	assert.Equal(t, 46, solver.Solve("day5/example.txt", day5.Part2Solution))
	assert.Equal(t, 56931769, solver.Solve("day5/input.txt", day5.Part2Solution))
}

func TestDay6(t *testing.T) {
	assert.Equal(t, 288, solver.Solve("day6/example.txt", day6.Part1Solution))
	assert.Equal(t, 1195150, solver.Solve("day6/input.txt", day6.Part1Solution))
	assert.Equal(t, 71503, solver.Solve("day6/example.txt", day6.Part2Solution))
	assert.Equal(t, -1, solver.Solve("day6/input.txt", day6.Part2Solution))
}
