package handlers

import (
	"fmt"
	"strings"
	"vyh/aoc23/internal/helpers"
)

type Race struct {
	time   int
	record int
}

func (r Race) recordEffectiveSpeed() int {
	// ceiling of distance over time
	mod := 0
	if r.record%r.time > 0 {
		mod = 1
	}
	return r.record/r.time + mod
}

func (r Race) wins(holdTime int) bool {
	return holdTime*(r.time-holdTime) > r.record
}

func (r Race) numWinningHolds() int {
	hold := r.recordEffectiveSpeed()
	win := r.wins(hold)
	for ; !win && (hold < r.time); win = r.wins(hold) {
		hold++
	}
	fmt.Printf("first winning hold time: %d\n", hold)

	wins := 0
	for ; win && (hold < r.time); win = r.wins(hold) {
		wins += 1
		hold++
	}
	fmt.Println("last winning hold time: ", hold-1)
	return wins
}

type Races []Race

func (r Races) numWinningHolds() []int {
	result := make([]int, len(r))
	for i, race := range r {
		result[i] = race.numWinningHolds()
	}
	return result
}

type DaySix struct{}

func (d6 DaySix) SolvePartOne(filename string) int {
	races := d6.parseRaces(helpers.ReadLines(filename))
	// fmt.Printf("races: %+v\n", races)
	return helpers.Product(races.numWinningHolds()...)
}

func (d6 DaySix) SolvePartTwo(filename string) int {
	race := d6.parseBigRace(helpers.ReadLines(filename))
	return race.numWinningHolds()
}

func (d6 DaySix) parseRaces(lines []string) Races {
	// helpers.Require(len(lines) == 2)
	timeParts := helpers.SplitAndTrim(lines[0], ":")
	helpers.Require(timeParts[0] == "Time")
	distParts := helpers.SplitAndTrim(lines[1], ":")
	helpers.Require(distParts[0] == "Distance")
	times := helpers.ToInts(helpers.SplitAndTrim(timeParts[1], " "))
	distances := helpers.ToInts(helpers.SplitAndTrim(distParts[1], " "))
	helpers.Require(len(times) == len(distances))
	races := make(Races, len(times))
	for i := 0; i < len(times); i++ {
		races[i] = Race{
			time:   times[i],
			record: distances[i],
		}
	}
	return races
}

func (d6 DaySix) parseBigRace(lines []string) Race {
	timeParts := helpers.SplitAndTrim(lines[0], ":")
	helpers.Require(timeParts[0] == "Time")
	distParts := helpers.SplitAndTrim(lines[1], ":")
	helpers.Require(distParts[0] == "Distance")
	return Race{
		time:   helpers.ToInt(strings.ReplaceAll(timeParts[1], " ", "")),
		record: helpers.ToInt(strings.ReplaceAll(distParts[1], " ", "")),
	}
}
