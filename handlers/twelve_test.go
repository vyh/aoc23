package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayTwelve_SolvePartOne(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{
			filename: "../internal/testdata/12",
			expected: 21,
		},
	}

	for _, tc := range tests {
		actual := DayTwelve{}.SolvePartOne(tc.filename)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestDayTwelve_SolvePartTwo(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{
			filename: "../internal/testdata/12",
			expected: 525152,
		},
	}

	for _, tc := range tests {
		actual := DayTwelve{}.SolvePartTwo(tc.filename)
		assert.Equal(t, tc.expected, actual)
	}
}
