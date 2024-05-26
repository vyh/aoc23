package handlers

import (
	"fmt"
	"vyh/aoc23/internal/helpers"
)

type DayFourteen struct{}

func (d DayFourteen) SolvePartOne(f string) int {
	grid := helpers.ToGrid(helpers.ReadLines(f))
	w := len(grid[0])
	rowRockCount := make([]int, len(grid))
	for j := 0; j < w; j++ {
		blockAt := -1
		rocks := 0
		for i := 0; i <= len(grid); i++ {
			if i == len(grid) || grid[i][j] == "#" {
				for r := 0; r < rocks; r++ {
					row := blockAt + r + 1
					// fmt.Println("rock at column", j, "row", row)
					rowRockCount[row]++
				}
				blockAt = i
				rocks = 0
			} else if grid[i][j] == "O" {
				rocks++
			}
		}
	}
	fmt.Println(rowRockCount)
	weighted := make([]int, len(rowRockCount))
	for i, c := range rowRockCount {
		weighted[i] = c * (len(rowRockCount) - i)
	}
	return helpers.Sum(weighted)
}

func (d DayFourteen) SolvePartTwo(f string) int {
	return -1
}
