package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDaySeven_SolvePartOne(t *testing.T) {
	filename := "../internal/testdata/7"
	expected := 6440
	actual := DaySeven{}.SolvePartOne(filename)
	assert.Equal(t, expected, actual)
}

func TestDaySeven_SolvePartTwo(t *testing.T) {
	filename := "../internal/testdata/7"
	expected := 5905
	actual := DaySeven{}.SolvePartTwo(filename)
	assert.Equal(t, expected, actual)
}
