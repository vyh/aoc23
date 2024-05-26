package handlers

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCalibration(t *testing.T) {
	lines := strings.Split(`two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen
	eightwothre`, "\n")

	expected := []int{29, 83, 13, 24, 42, 14, 76, 82}

	for i, line := range lines {
		actual := parseCalibration(line)
		assert.Equal(t, expected[i], actual)
	}
}

func TestParseCalibrations(t *testing.T) {
	lines := strings.Split(`two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen
	eightwothre`, "\n")

	expected := []int{29, 83, 13, 24, 42, 14, 76, 82}

	actual := parseCalibrations(lines)
	assert.Equal(t, expected, actual)
}
