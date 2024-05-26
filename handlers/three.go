package handlers

import (
	"fmt"
	"regexp"
	"strconv"
	"vyh/aoc23/internal/helpers"

	log "github.com/sirupsen/logrus"
)

var (
	numberPattern = regexp.MustCompile(`[0-9]+`)
	symbolPattern = regexp.MustCompile(`[^a-zA-Z0-9.]`)
	starPattern   = regexp.MustCompile(`\*`)
)

type indexMap map[int]map[int]bool
type indexRanges [][]int
type starMap map[string][]int

type DayThree struct{}

func (d3 DayThree) Solve(filename string) int {
	lines := helpers.ReadLines(filename)
	symbols := d3.getIndexMap(*symbolPattern, lines)
	// fmt.Println("symbols: ", symbols)
	numbers := d3.getAllIndexes(*numberPattern, lines)
	// fmt.Println("number indices: ", numbers)
	filteredNumbers := d3.filterAdjacentIndexes(numbers, symbols)
	// fmt.Println("filtered number indices: ", filteredNumbers)
	numberValues := d3.toIntegers(d3.valuesAtRanges(lines, filteredNumbers))
	// fmt.Println("numbers: ", numberValues)
	return helpers.Sum(numberValues)
}

func (d3 DayThree) SumGearRatios(filename string) int {
	lines := helpers.ReadLines(filename)
	stars := d3.getIndexMap(*starPattern, lines)
	// fmt.Println("stars: ", stars)
	numbers := d3.getAllIndexes(*numberPattern, lines)
	// fmt.Println("number indices: ", numbers)
	possibleGears := d3.mapAdjacentNumbers(numbers, stars, lines)
	// fmt.Println("star map: ", possibleGears)
	ratios := d3.gearRatios(possibleGears)
	// fmt.Println("gear ratios: ", ratios)
	return helpers.Sum(ratios)
}

func (d3 DayThree) getIndexes(p regexp.Regexp, s string) indexRanges {
	return p.FindAllStringIndex(s, -1)
}

func (d3 DayThree) getAllIndexes(p regexp.Regexp, lines []string) []indexRanges {
	indexes := []indexRanges{}
	for _, line := range lines {
		indexes = append(indexes, d3.getIndexes(p, line))
	}
	return indexes
}

func (d3 DayThree) getIndexMap(p regexp.Regexp, lines []string) indexMap {
	im := indexMap{}
	for i, line := range lines {
		im[i] = map[int]bool{}
		indexes := d3.getIndexes(p, line)
		for _, idxRange := range indexes {
			for j := idxRange[0]; j < idxRange[1]; j++ {
				im[i][j] = true
			}
		}
	}
	return im
}

func (d3 DayThree) isSymbolAdjacent(row int, colRange []int, symbols indexMap) bool {
	for _, cell := range helpers.GetAdjacentCells(row, colRange) {
		if _, ok := symbols[cell[0]][cell[1]]; ok {
			return true
		}
	}
	return false
}

func (d3 DayThree) adjacentSymbols(row int, colRange []int, symbols indexMap) []string {
	adj := []string{}
	for _, cell := range helpers.GetAdjacentCells(row, colRange) {
		if _, ok := symbols[cell[0]][cell[1]]; ok {
			adj = append(adj, fmt.Sprintf("%d,%d", cell[0], cell[1]))
		}
	}
	return adj
}

func (d3 DayThree) filterAdjacentIndexes(indexes []indexRanges, symbols indexMap) []indexRanges {
	filtered := []indexRanges{}
	for row, colRanges := range indexes {
		idxRanges := indexRanges{}
		for _, colRange := range colRanges {
			if d3.isSymbolAdjacent(row, colRange, symbols) {
				idxRanges = append(idxRanges, colRange)
			}
		}
		filtered = append(filtered, idxRanges)
	}
	return filtered
}

func coordKey(i, j int) string {
	return fmt.Sprintf("%d,%d", i, j)
}

func (d3 DayThree) mapAdjacentNumbers(indexes []indexRanges, symbols indexMap, lines []string) starMap {
	sMap := starMap{}
	for i, row := range symbols {
		for j := range row {
			sMap[coordKey(i, j)] = []int{}
		}
	}
	for row, colRanges := range indexes {
		for _, colRange := range colRanges {
			n := d3.numberAt(lines, row, colRange)
			for _, star := range d3.adjacentSymbols(row, colRange, symbols) {
				sMap[star] = append(sMap[star], n)
			}
		}
	}
	return sMap
}

func (d3 DayThree) gearRatios(sm starMap) []int {
	ratios := []int{}
	for _, numbers := range sm {
		if len(numbers) == 2 {
			// log.Infof("appending %d * %d", numbers[0], numbers[1])
			ratios = append(ratios, numbers[0]*numbers[1])
		}
	}
	return ratios
}

func (d3 DayThree) numberAt(lines []string, row int, colRange []int) int {
	s := lines[row][colRange[0]:colRange[1]]
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Panicf(
			"failed to parse number at %d,%d:%d : %v",
			row, colRange[0], colRange[1], err,
		)
	}
	return n
}

func (d3 DayThree) valuesAtRanges(lines []string, indexes []indexRanges) []string {
	values := []string{}
	if len(lines) != len(indexes) {
		log.Panicf("%d lines != %d indexes", len(lines), len(indexes))
	}
	for i, line := range lines {
		for _, idxRange := range indexes[i] {
			values = append(values, line[idxRange[0]:idxRange[1]])
		}
	}
	return values
}

func (d3 DayThree) toIntegers(values []string) []int {
	integers := []int{}
	for _, v := range values {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Panicf("failed to convert %s: %v", v, err)
		}
		integers = append(integers, n)
	}
	return integers
}
