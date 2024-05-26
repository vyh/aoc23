package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayFour_Solve(t *testing.T) {
	filename := "../internal/testdata/4"
	expected := 13
	actual := DayFour{}.Solve(filename)
	assert.Equal(t, expected, actual)
}

func TestDayFour_CountCopies(t *testing.T) {
	filename := "../internal/testdata/4"
	expected := 30
	actual := DayFour{}.CountCopies(filename)
	assert.Equal(t, expected, actual)
}
