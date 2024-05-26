package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayTen_SolvePartOne(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{
			filename: "../internal/testdata/10a",
			expected: 4,
		},
		{
			filename: "../internal/testdata/10b",
			expected: 8,
		},
	}

	for _, tc := range tests {
		actual := DayTen{}.SolvePartOne(tc.filename)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestDayTen_SolvePartTwo(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{
			filename: "../internal/testdata/10a",
			expected: 4,
		},
		{
			filename: "../internal/testdata/10b",
			expected: 8,
		},
	}

	for _, tc := range tests {
		actual := DayTen{}.SolvePartTwo(tc.filename)
		assert.Equal(t, tc.expected, actual)
	}
}
