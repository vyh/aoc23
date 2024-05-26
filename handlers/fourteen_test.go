package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayFourteen_SolvePartOne(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{
			filename: "../internal/testdata/14",
			expected: 136,
		},
	}

	for _, tc := range tests {
		actual := DayFourteen{}.SolvePartOne(tc.filename)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestDayFourteen_SolvePartTwo(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{
			filename: "../internal/testdata/14",
			expected: -1,
		},
	}

	for _, tc := range tests {
		actual := DayFourteen{}.SolvePartTwo(tc.filename)
		assert.Equal(t, tc.expected, actual)
	}
}
