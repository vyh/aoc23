package handlers

import (
	"strings"
	"vyh/aoc23/internal/helpers"

	log "github.com/sirupsen/logrus"
)

type galaxy struct { // node
	id int
	i  int
	j  int
}

func (g1 galaxy) distance(g2 galaxy) int {
	if g1.id >= g2.id { // > bc don't count same pair twice
		return 0
	}
	return helpers.Abs(g2.i-g1.i) + helpers.Abs(g2.j-g1.j)
	// maybe part 2 will be like um, actually, you can go diagonally
}

func (g1 galaxy) expandedDistance(g2 galaxy, emptyRows, emptyCols []int, multiplier int) int {
	if g1.id >= g2.id { // > bc don't count same pair twice
		return 0
	}
	additionalSpace := 0
	x1, x2 := helpers.Min(g1.j, g2.j), helpers.Max(g1.j, g2.j)
	for _, c := range emptyCols {
		if c > x1 && c < x2 {
			additionalSpace += (multiplier - 1) // -1 simplifies final calc below
		}
	}
	y1, y2 := helpers.Min(g1.i, g2.i), helpers.Max(g1.i, g2.i)
	for _, r := range emptyRows {
		if r > y1 && r < y2 {
			additionalSpace += (multiplier - 1)
		}
	}
	return x2 - x1 + y2 - y1 + additionalSpace
	// part 2 was actually like "that thing you thought would make your calculation
	// fiddly is exactly what you need to do"
}

// type galaxyPath struct { // edge
// 	a        galaxy
// 	b        galaxy
// 	distance int
// }

type DayEleven struct {
	multiplier int
}

func (d DayEleven) SolvePartOne(f string) int {
	d.multiplier = 2
	lines := helpers.ReadLines(f)
	if strings.TrimSpace(lines[len(lines)-1]) == "" {
		lines = lines[:len(lines)-1]
	}

	grid, eRows, eCols := d.parseGrid(lines)

	galaxies := []galaxy{}
	id := 1
	for i, row := range grid {
		for j, col := range row {
			if col == "#" {
				galaxies = append(galaxies, galaxy{
					id: id,
					i:  i,
					j:  j,
				})
				id++
			}
		}
	}
	// fmt.Printf("%+v\n", galaxies)

	distances := []int{}
	for _, g1 := range galaxies {
		for _, g2 := range galaxies {
			distances = append(distances, g1.expandedDistance(g2, eRows, eCols, d.multiplier))
		}
	}
	// fmt.Println(distances)

	return helpers.Sum(distances)
}

func (d DayEleven) parseGrid(lines []string) ([][]string, []int, []int) {
	grid := [][]string{}
	w := len(lines[0])
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		if len(line) != w {
			log.Panic("uneven line widths")
		}
		grid = append(grid, strings.Split(line, ""))
	}

	emptyRows := []int{}
	emptyCols := []int{}
	for i, row := range grid {
		if d.isEmpty(row) {
			emptyRows = append(emptyRows, i)
		}
	}
	for j := 0; j < w; j++ {
		if d.isEmpty(d.getCol(j, grid)) {
			emptyCols = append(emptyCols, j)
		}
	}
	// // unfeasible for part 2
	// // duplicate empty rows & cols starting from the end to not mess up other indexes
	// for i := len(emptyRows) - 1; i >= 0; i-- {
	// 	idx := emptyRows[i]
	// 	grid = append(grid[:idx], append([][]string{grid[idx]}, grid[idx:]...)...)
	// }
	// for j := len(emptyCols) - 1; j >= 0; j-- {
	// 	idx := emptyCols[j]
	// 	for i := range grid {
	// 		row := grid[i]
	// 		row = append(row[:idx], append([]string{row[idx]}, row[idx:]...)...)
	// 		grid[i] = row
	// 	}
	// }
	return grid, emptyRows, emptyCols
}

func (d DayEleven) isEmpty(row []string) bool {
	for _, c := range row {
		if c != "." {
			return false
		}
	}
	return true
}

func (d DayEleven) getCol(j int, g [][]string) []string {
	col := []string{} // don't assume it's square
	for i := range g {
		col = append(col, g[i][j])
	}
	return col
}

func (d DayEleven) SolvePartTwo(f string) int {
	lines := helpers.ReadLines(f)
	if strings.TrimSpace(lines[len(lines)-1]) == "" {
		lines = lines[:len(lines)-1]
	}

	grid, eRows, eCols := d.parseGrid(lines)

	galaxies := []galaxy{}
	id := 1
	for i, row := range grid {
		for j, col := range row {
			if col == "#" {
				galaxies = append(galaxies, galaxy{
					id: id,
					i:  i,
					j:  j,
				})
				id++
			}
		}
	}
	// fmt.Printf("%+v\n", galaxies)

	distances := []int{}
	for _, g1 := range galaxies {
		for _, g2 := range galaxies {
			distances = append(distances, g1.expandedDistance(g2, eRows, eCols, d.multiplier))
		}
	}
	// fmt.Println(distances)

	return helpers.Sum(distances)
}
