package handlers

import (
	"regexp"
	"strings"
	"vyh/aoc23/internal/helpers"

	log "github.com/sirupsen/logrus"
)

var p = regexp.MustCompile(`([0-9]|one|two|three|four|five|six|seven|eight|nine)`)
var rp = regexp.MustCompile(`([0-9]|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin)`)

func parseCalibration(line string) int {
	left := p.FindString(line)
	right := helpers.ReverseString(rp.FindString(helpers.ReverseString(line)))
	if len(left) == 0 {
		return 0
	}
	tens := helpers.StringToDigit(left)
	ones := helpers.StringToDigit(right)
	return tens*10 + ones
}

func parseCalibrations(lines []string) []int {
	calibrations := []int{}
	for _, line := range lines {
		calibrations = append(calibrations, parseCalibration(line))
	}
	return calibrations
}

func GetCalibrationSum(filename string) int {
	contents, err := helpers.ReadText(filename)
	if err != nil {
		log.Error(err)
		return 0
	}
	lines := strings.Split(contents, "\n")
	sum := 0
	for _, calibration := range parseCalibrations(lines) {
		sum += calibration
	}
	return sum
}
