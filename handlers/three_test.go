package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayThree_Solve(t *testing.T) {
	filename := "../internal/testdata/3"
	expected := 4361
	actual := DayThree{}.Solve(filename)
	assert.Equal(t, expected, actual)
}

func TestDayThree_SumGearRatios(t *testing.T) {
	tests := []struct {
		description string
		filename    string
		expected    int
	}{
		{
			description: "provided test",
			filename:    "../internal/testdata/3",
			expected:    467835,
		},
		{
			description: "three numbers by star is not included",
			filename:    "../internal/testdata/3b",
			expected:    16345,
		},
		{
			description: "a number in >= 1 gear",
			filename:    "../internal/testdata/3c",
			expected:    489430,
		},
		{
			description: "uneven line lengths",
			filename:    "../internal/testdata/3d",
			expected:    467835,
		},
	}

	for _, tc := range tests {
		actual := DayThree{}.SumGearRatios(tc.filename)
		assert.Equal(t, tc.expected, actual, tc.description)
	}
}
