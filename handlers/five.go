package handlers

import (
	"fmt"
	"math"
	"strings"
	"vyh/aoc23/internal/helpers"

	log "github.com/sirupsen/logrus"
)

type seedRange struct {
	start  int
	length int
}

type seedRanges struct {
	ranges []seedRange
}

func (s seedRanges) contain(v int) bool {
	for _, r := range s.ranges {
		if v >= r.start && v < r.start+r.length {
			return true
		}
	}
	return false
}

type mapRange struct {
	sourceStart int
	destStart   int
	length      int
}

func (mr mapRange) isInRange(input int) bool {
	return input >= mr.sourceStart && input < mr.sourceStart+mr.length
}

func (mr mapRange) getDest(input int) (int, error) {
	if !mr.isInRange(input) {
		return 0, fmt.Errorf("out of range")
	}
	return mr.destStart + (input - mr.sourceStart), nil
}

type farmMapper struct {
	destination string
	ranges      []mapRange
}

func (fm farmMapper) getDest(input int) int {
	for _, r := range fm.ranges {
		if d, err := r.getDest(input); err == nil {
			return d
		}
	}
	return input
}

type farmMappers map[string]farmMapper

func (fm farmMappers) validateReachableLocationMapper(source string) {
	for m, ok := fm[source]; ok; m, ok = fm[m.destination] {
		if m.destination == "location" {
			return
		}
	}
	log.Panicf("cannot reach destination from %q", source)
}

func (fm farmMappers) validateReachableMapper(source, destination string) {
	for m, ok := fm[source]; ok; m, ok = fm[m.destination] {
		if m.destination == destination {
			return
		}
	}
	log.Panicf("cannot reach %q from %q", destination, source)
}

func (fm farmMappers) getDest(source string, value int) (string, int) {
	if mapper, ok := fm[source]; ok {
		return mapper.destination, mapper.getDest(value)
	}
	log.Panicf("no mapper for source %q", source)
	return "", -1
}

func (fm farmMappers) getSeedLocation(seed int) int {
	source := "seed"
	value := seed
	fm.validateReachableLocationMapper(source)
	for source != "location" {
		source, value = fm.getDest(source, value)
	}
	return value
}

func (fm farmMappers) getFinalDestination(value int, source, destination string) int {
	fm.validateReachableMapper(source, destination)
	// fmt.Println("source ", source, "value ", value)
	for source != destination {
		source, value = fm.getDest(source, value)
		// fmt.Printf("type %v, value %v\n", source, value)
	}
	return value
}

func (fm farmMappers) getLeastLocationWithSeed(seeds seedRanges) int {
	for loc := 0; loc < 200_000_000; loc++ {
		seed := fm.getFinalDestination(loc, "location", "seed")
		if seeds.contain(seed) {
			fmt.Printf("min at seed %d: %d\n", seed, loc)
			return loc
		}
	}
	log.Panic("failed to solve")
	return -1
}

func (fm farmMappers) getLeastSeedLocation(seeds seedRanges) int {
	min := math.MaxInt
	for _, r := range seeds.ranges {
		for offset := 0; offset < r.length; offset++ {
			// fmt.Println("seed: ", r.start+offset)
			loc := fm.getSeedLocation(r.start + offset)
			if loc < min {
				// fmt.Printf("new min at seed %d: %d\n", r.start+offset, loc)
				min = loc
			}
		}
	}
	return min
}

type DayFive struct{}

func (d5 DayFive) LowestLocationNumber(filename string) int {
	allLines := helpers.ReadLines(filename)
	lines, j := d5.sliceLineSet(allLines)
	seeds := d5.parseSeeds(lines[0])
	mappers := farmMappers{}
	for next := j; next >= 0 && next < len(allLines); {
		lines, j = d5.sliceLineSet(allLines[next:])
		next += j
		// fmt.Println("sliced lines:", len(lines), lines, next)
		source, mapper := d5.parseMapper(lines)
		mappers[source] = mapper
	}
	// fmt.Println(seeds)
	// fmt.Printf("%+v\n", mappers)
	locations := make([]int, len(seeds))
	for i, seed := range seeds {
		locations[i] = mappers.getSeedLocation(seed)
	}
	return helpers.Min(locations...)
}

func (d5 DayFive) Solve(filename string) int {
	allLines := helpers.ReadLines(filename)
	lines, j := d5.sliceLineSet(allLines)
	seeds := d5.parseSeedRanges(lines[0])
	mappers := farmMappers{}
	for next := j; next >= 0 && next < len(allLines[1:]); {
		lines, j = d5.sliceLineSet(allLines[next:])
		next += j
		// fmt.Println("sliced lines:", len(lines), lines, next)
		source, mapper := d5.parseMapper(lines)
		mappers[source] = mapper
	}
	// fmt.Println("seed ranges: ", seeds.ranges, j)
	// return -1
	return mappers.getLeastSeedLocation(seeds)
}

func (d5 DayFive) ReverseSolve(filename string) int {
	allLines := helpers.ReadLines(filename)
	lines, j := d5.sliceLineSet(allLines)
	seeds := d5.parseSeedRanges(lines[0])
	mappers := farmMappers{}
	for next := j; next >= 0 && next < len(allLines[1:]); {
		lines, j = d5.sliceLineSet(allLines[next:])
		next += j
		// fmt.Println("sliced lines:", len(lines), lines, next)
		source, mapper := d5.reverseMap(d5.parseMapper(lines))
		mappers[source] = mapper
	}
	// fmt.Printf("%+v\n", seeds)
	// fmt.Printf("%+v\n", mappers)
	return mappers.getLeastLocationWithSeed(seeds)
}

func (d5 DayFive) sliceLineSet(lines []string) ([]string, int) {
	i := 0
	for ; strings.TrimSpace(lines[i]) == ""; i++ {
	}
	for j, line := range lines[i:] {
		if strings.TrimSpace(line) == "" {
			return lines[i : i+j], j + 1
		}
	}
	return lines[i:], len(lines)
}

func (d5 DayFive) parseSeeds(s string) []int {
	parts := helpers.SplitAndTrim(s, ":")
	if parts[0] != "seeds" {
		log.Panicf("expected 'seeds', got %s", parts[0])
	}
	parts = helpers.SplitAndTrim(parts[1], " ")
	seeds := make([]int, len(parts))
	for i, p := range parts {
		seeds[i] = helpers.ToInt(p)
	}
	return seeds
}

func (d5 DayFive) parseSeedRanges(s string) seedRanges {
	parts := helpers.SplitAndTrim(s, ":")
	if parts[0] != "seeds" {
		log.Panicf("expected 'seeds', got %s", parts[0])
	}
	allParts := helpers.SplitAndTrim(parts[1], " ")
	if len(allParts)%2 != 0 {
		log.Panic("uneven number of seed parts: ", parts)
	}
	seeds := seedRanges{
		ranges: []seedRange{},
	}
	for i := 0; i < len(allParts)-1; i += 2 {
		start, len := allParts[i], allParts[i+1]
		seeds.ranges = append(seeds.ranges, seedRange{
			start:  helpers.ToInt(start),
			length: helpers.ToInt(len),
		})
	}
	// fmt.Printf("seeds: %+v\n", seeds)
	return seeds
}

func (d5 DayFive) parseMapper(lines []string) (string, farmMapper) {
	parts := helpers.SplitAndTrim(strings.TrimSpace(strings.TrimSuffix(lines[0], "map:")), "-")
	if len(parts) != 3 || parts[1] != "to" {
		log.Panicf("unexpected split result on map header: %v", parts)
	}
	source, destination := parts[0], parts[2]
	mapper := farmMapper{
		destination: destination,
		ranges:      make([]mapRange, len(lines[1:])),
	}
	for i, line := range lines[1:] {
		mapper.ranges[i] = d5.parseRange(line)
	}
	return source, mapper
}

func (d5 DayFive) reverseMap(source string, fm farmMapper) (string, farmMapper) {
	dest := fm.destination
	fm.destination = source
	ranges := []mapRange{}
	for _, r := range fm.ranges {
		r.destStart, r.sourceStart = r.sourceStart, r.destStart
		ranges = append(ranges, r)
	}
	fm.ranges = ranges
	return dest, fm
}

func (d5 DayFive) parseRange(line string) mapRange {
	parts := helpers.SplitAndTrim(line, " ")
	if len(parts) != 3 {
		log.Panicf("expected 3 values in line: %s, got: %d", line, len(parts))
	}
	return mapRange{
		destStart:   helpers.ToInt(parts[0]),
		sourceStart: helpers.ToInt(parts[1]),
		length:      helpers.ToInt(parts[2]),
	}
}
