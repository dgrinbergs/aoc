package _024

import (
	solver "github.com/dgrinbergs/aoc"
	"github.com/dgrinbergs/aoc/2024/day1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay1(t *testing.T) {
	assert.Equal(t, 3508942, solver.Solve("day1/input.txt", day1.Part1Solution))
}
