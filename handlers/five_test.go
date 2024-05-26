package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayFive_LowestLocationNumber(t *testing.T) {
	filename := "../internal/testdata/5"
	expected := 35
	actual := DayFive{}.LowestLocationNumber(filename)
	assert.Equal(t, expected, actual)
}

func TestDayFive_Solve(t *testing.T) {
	filename := "../internal/testdata/5"
	expected := 46
	actual := DayFive{}.Solve(filename)
	assert.Equal(t, expected, actual)
	actual = DayFive{}.ReverseSolve(filename)
	assert.Equal(t, expected, actual)
}
