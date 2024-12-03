package _024

import (
	solver "github.com/dgrinbergs/aoc"
	"github.com/dgrinbergs/aoc/2024/day1"
	"github.com/dgrinbergs/aoc/2024/day2"
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
