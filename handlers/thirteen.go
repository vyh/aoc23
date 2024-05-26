package handlers

import (
	"fmt"
	"strings"
	"vyh/aoc23/internal/helpers"
)

func isReflectedAcross(midx int, lines []string) bool {
	j := midx + 1
	for i := midx; i >= 0 && j < len(lines); i, j = i-1, j+1 {
		if lines[i] != lines[j] {
			return false
		}
	}
	return true
}

type DayThirteen struct{}

func (d DayThirteen) SolvePartOne(f string) int {
	lines := helpers.ReadLines(f)
	grids := [][]string{}
	i, j := 0, 0
	for j < len(lines) {
		for ; j < len(lines) && strings.TrimSpace(lines[j]) != ""; j++ {
		}
		grids = append(grids, lines[i:j])
		for i = j; i < len(lines) && strings.TrimSpace(lines[i]) == ""; i++ {
		}
		j = i
	}
	fmt.Println("n grids:", len(grids))

	rowMirrorSizes := []int{}
	colMirrorSizes := []int{}

	for gi, grid := range grids {
		fmt.Println("grid #", gi)
		// for _, r := range grid {
		// 	fmt.Println(r)
		// }
		// fmt.Println()
		t := helpers.Transpose(helpers.ToGrid(grid))
		cols := make([]string, len(t))
		for i, row := range t {
			cols[i] = strings.Join(row, "")
		}
		// fmt.Println(grid)
		// fmt.Println(cols)

		for i := 0; i < len(grid)-1; i++ {
			if grid[i] == grid[i+1] && isReflectedAcross(i, grid) {
				rowMirrorSizes = append(rowMirrorSizes, i+1)
				break
			}
		}

		for i := 0; i < len(cols)-1; i++ {
			if cols[i] == cols[i+1] && isReflectedAcross(i, cols) {
				colMirrorSizes = append(colMirrorSizes, i+1)
				break
			}
		}
	}

	return (100 * helpers.Sum(rowMirrorSizes)) + helpers.Sum(colMirrorSizes)
}

func (d DayThirteen) SolvePartTwo(f string) int {
	return -1
}
