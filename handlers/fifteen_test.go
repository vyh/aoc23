package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayFifteen_SolvePartOne(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{
			filename: "../internal/testdata/15",
			expected: 1320,
		},
	}

	for _, tc := range tests {
		actual := DayFifteen{}.SolvePartOne(tc.filename)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestDayFifteen_SolvePartTwo(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{
			filename: "../internal/testdata/15",
			expected: 145,
		},
	}

	for _, tc := range tests {
		actual := DayFifteen{}.SolvePartTwo(tc.filename)
		assert.Equal(t, tc.expected, actual)
	}
}
