package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayTwo_Solve(t *testing.T) {
	filename := "../internal/testdata/2"
	total := Cubes{
		Red:   12,
		Green: 13,
		Blue:  14,
	}
	expected := 8
	assert.Equal(t, expected, DayTwo{}.Solve(filename, total))
}

func TestDayTwo_SolvePowerSum(t *testing.T) {
	filename := "../internal/testdata/2"
	expected := 2286
	assert.Equal(t, expected, DayTwo{}.SolvePowerSum(filename))
}
