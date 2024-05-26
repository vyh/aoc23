package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDaySix_SolvePartOne(t *testing.T) {
	filename := "../internal/testdata/6"
	expected := 288
	actual := DaySix{}.SolvePartOne(filename)
	assert.Equal(t, expected, actual)
}

func TestDaySix_SolvePartTwo(t *testing.T) {
	filename := "../internal/testdata/6"
	expected := 71503
	actual := DaySix{}.SolvePartTwo(filename)
	assert.Equal(t, expected, actual)
}
