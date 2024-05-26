package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayEleven_SolvePartOne(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{
			filename: "../internal/testdata/11",
			expected: 374,
		},
	}

	for _, tc := range tests {
		actual := DayEleven{}.SolvePartOne(tc.filename)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestDayEleven_SolvePartTwo(t *testing.T) {
	tests := []struct {
		filename   string
		expected   int
		multiplier int
	}{
		{
			filename:   "../internal/testdata/11",
			expected:   374,
			multiplier: 2,
		},
		{
			filename:   "../internal/testdata/11",
			expected:   1030,
			multiplier: 10,
		},
		{
			filename:   "../internal/testdata/11",
			expected:   8410,
			multiplier: 100,
		},
	}

	for _, tc := range tests {
		actual := DayEleven{tc.multiplier}.SolvePartTwo(tc.filename)
		assert.Equal(t, tc.expected, actual)
	}
}
