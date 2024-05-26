package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayEight_SolvePartOne(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{
			filename: "../internal/testdata/8a",
			expected: 2,
		},
		{
			filename: "../internal/testdata/8b",
			expected: 6,
		},
	}

	for _, tc := range tests {
		actual := DayEight{}.SolvePartOne(tc.filename)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestDayEight_SolvePartTwo(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{
			filename: "../internal/testdata/8a",
			expected: 2,
		},
		{
			filename: "../internal/testdata/8b",
			expected: 6,
		},
		{
			filename: "../internal/testdata/8c",
			expected: 6,
		},
	}

	for _, tc := range tests {
		actual := DayEight{}.SolvePartTwo(tc.filename)
		assert.Equal(t, tc.expected, actual)
	}
}
