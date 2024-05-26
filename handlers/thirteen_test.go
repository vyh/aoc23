package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayThirteen_SolvePartOne(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{
			filename: "../internal/testdata/13",
			expected: 405,
		},
	}

	for _, tc := range tests {
		actual := DayThirteen{}.SolvePartOne(tc.filename)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestDayThirteen_SolvePartTwo(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{
			filename: "../internal/testdata/13",
			expected: -1,
		},
	}

	for _, tc := range tests {
		actual := DayThirteen{}.SolvePartTwo(tc.filename)
		assert.Equal(t, tc.expected, actual)
	}
}
