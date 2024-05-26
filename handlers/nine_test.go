package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayNine_SolvePartOne(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{
			filename: "../internal/testdata/9",
			expected: 114,
		},
	}

	for _, tc := range tests {
		actual := DayNine{}.SolvePartOne(tc.filename)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestDayNine_SolvePartTwo(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{
			filename: "../internal/testdata/9",
			expected: 2,
		},
	}

	for _, tc := range tests {
		actual := DayNine{}.SolvePartTwo(tc.filename)
		assert.Equal(t, tc.expected, actual)
	}
}
