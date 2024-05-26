package handlers

import (
	"fmt"
	"regexp"
	"strings"
	"vyh/aoc23/internal/helpers"

	log "github.com/sirupsen/logrus"
)

var forkPattern = regexp.MustCompile(`^([0-9A-Z]{3})\s+=\s+\(([0-9A-Z]{3}),\s+([0-9A-Z]{3})\)$`)

type Move int

const (
	L Move = iota
	R
)

type fork map[Move]string

type desertMap struct {
	moveSequence []Move
	forks        map[string]fork
	startPoints  []string
}

func newDesertMap(lines []string) desertMap {
	dm := desertMap{
		moveSequence: []Move{},
		forks:        map[string]fork{},
		startPoints:  []string{},
	}
	for _, c := range strings.Split(lines[0], "") {
		switch c {
		case "L":
			dm.moveSequence = append(dm.moveSequence, L)
		case "R":
			dm.moveSequence = append(dm.moveSequence, R)
		default:
			log.Panicf("unexpected move instruction %q", c)
		}
	}
	helpers.Require(strings.TrimSpace(lines[1]) == "")
	for _, line := range lines[2:] {
		if strings.TrimSpace(line) == "" {
			continue
		}
		matches := forkPattern.FindAllStringSubmatch(line, -1)[0]
		helpers.Require(len(matches) == 4)
		_, ok := dm.forks[matches[1]]
		helpers.Require(!ok)
		dm.forks[matches[1]] = fork{
			L: matches[2],
			R: matches[3],
		}
		if strings.HasSuffix(matches[1], "A") {
			dm.startPoints = append(dm.startPoints, matches[1])
		}
	}
	// fmt.Printf("%+v\n", dm)
	return dm
}

func (dm desertMap) isTerminal(loc string) bool {
	f := dm.forks[loc]
	return loc == f[L] && loc == f[R]
}

func (dm desertMap) countSteps(start, end string) int {
	steps := 0
	current := start
	endPtn := regexp.MustCompile("^" + end + "$")

	for i := 0; !endPtn.MatchString(current); i = (i + 1) % len(dm.moveSequence) {
		fork := dm.forks[current]
		if dm.isTerminal(current) {
			break
		}
		current = fork[dm.moveSequence[i]]
		steps++
	}

	if !endPtn.MatchString(current) {
		log.Panicf("reached terminal %q is not destination %q", current, end)
	}
	return steps
}

func (dm desertMap) cycleLength(start string) (int, int) {
	steps := 0
	current := start
	moves := map[string]bool{}
	cycle := 0
	countingCycle := false
	key := func(s string, n int) string { return fmt.Sprintf("%s:%d", s, n) }

	for i := 0; !moves[key(current, i)] || cycle == 0; i = (i + 1) % len(dm.moveSequence) {
		if moves[key(current, i)] {
			countingCycle = true
			moves = map[string]bool{}
		}
		moves[key(current, i)] = true
		current = dm.forks[current][dm.moveSequence[i]]
		if !countingCycle {
			steps++
		} else {
			cycle++
		}
	}

	return steps, cycle
}

func (dm desertMap) cycleLengths() ([]int, []int) {
	lengths := make([]int, len(dm.startPoints))
	cycleLengths := make([]int, len(dm.startPoints))
	for i, start := range dm.startPoints {
		lengths[i], cycleLengths[i] = dm.cycleLength(start)
	}
	return lengths, cycleLengths
}

func (dm desertMap) countStepsAllPaths() int {
	steps := 0
	current := dm.startPoints[:]
	fmt.Println(len(current), " starting points")

	z := map[int]struct{}{}
	for i := 0; len(z) != len(current); i = (i + 1) % len(dm.moveSequence) {
		z = map[int]struct{}{}
		for j, loc := range current {
			current[j] = dm.forks[loc][dm.moveSequence[i]]
			if strings.HasSuffix(current[j], "Z") {
				z[j] = struct{}{}
			}
		}
		steps++
		// if i+1 == len(dm.moveSequence) {
		// 	fmt.Println("reached end of sequence; current: ", current)
		// }
	}

	return steps
}

// func (dm desertMap) reverseWalk() int {
// 	endPoints := map[string]
// 	return -1
// }

type DayEight struct{}

func (d8 DayEight) SolvePartOne(f string) int {
	lines := helpers.ReadLines(f)
	dm := newDesertMap(lines)
	return dm.countSteps("AAA", "ZZZ")
}

func (d8 DayEight) SolvePartTwo(f string) int {
	lines := helpers.ReadLines(f)
	dm := newDesertMap(lines)
	lengths, cycleLengths := dm.cycleLengths()
	fmt.Println(lengths, cycleLengths)
	steps := []int{}
	for _, s := range dm.startPoints {
		steps = append(steps, dm.countSteps(s, "..Z"))
	}
	fmt.Println(steps)
	// lcm := -1
	if len(steps) == 1 {
		return steps[0]
	}
	if len(steps) > 2 {
		return helpers.LCM(steps[0], steps[1], steps[2:]...)
	} else {
		return helpers.LCM(steps[0], steps[1])
	}
	// fmt.Println(lcm)
	// return lcm
}
