package handlers

import (
	"strings"
	"vyh/aoc23/internal/helpers"

	log "github.com/sirupsen/logrus"
)

type move []int

var (
	north     = move{-1, 0}
	east      = move{0, 1}
	south     = move{1, 0}
	west      = move{0, -1}
	opposites = map[string]string{
		"north": "south",
		"east":  "west",
		"south": "north",
		"west":  "east",
	}
)

type pipe map[string]move

func (p pipe) moveFrom(dir string) (string, move) {
	for d, m := range p {
		if d != dir {
			return d, m
		}
	}
	log.Panic("didn't find a move!")
	return "", nil
}

var pipes = map[string]pipe{
	"F": pipe{"east": east, "south": south},
	"7": pipe{"west": west, "south": south},
	"J": pipe{"north": north, "west": west},
	"L": pipe{"north": north, "east": east},
	"-": pipe{"west": west, "east": east},
	"|": pipe{"north": north, "south": south},
	"S": pipe{}, // moves unknown at init
}

type pipeCell []int

func (c pipeCell) equals(other pipeCell) bool {
	return c[0] == other[0] && c[1] == other[1]
}

func (c pipeCell) plus(other move) pipeCell {
	return pipeCell{c[0] + other[0], c[1] + other[1]}
}

type pipeGrid struct {
	start       pipeCell
	height      int
	width       int
	cycleLength int
	rows        [][]*pipe
	cameFrom    string
}

func newPipeGrid(rows [][]string) *pipeGrid {
	g := pipeGrid{
		height: len(rows),
	}
	if g.height == 0 {
		g.width = 0
	} else {
		g.width = len(rows[0])
	}
	grid := make([][]*pipe, g.height)
	for i, row := range rows {
		if len(row) != g.width {
			log.Panic("not a rectangular grid!", rows)
		}
		grid[i] = make([]*pipe, g.width)
		for j, cell := range row {
			if cell == "S" {
				g.start = pipeCell{i, j}
				cardinalMoves := map[string]move{
					"north": north,
					"south": south,
					"east":  east,
					"west":  west,
				}
				for name, dir := range cardinalMoves {
					other := g.start.plus(dir)
					if !g.isInBounds(other[0], other[1]) {
						continue
					}
					if p, ok := pipes[rows[other[0]][other[1]]]; ok {
						if _, ok := p[opposites[name]]; ok {
							pipes["S"][name] = dir
						}
					}
				}
				helpers.Require(len(pipes["S"]) == 2)
			}
			p := pipes[cell]
			grid[i][j] = &p
		}
	}
	g.rows = grid
	// fmt.Println(g)
	return &g
}

func (g *pipeGrid) move(cell pipeCell) pipeCell {
	pipe := *g.rows[cell[0]][cell[1]]
	dir, m := "", move{}
	if g.cameFrom == "" {
		// just go whatever you find first, stop trying to go clockwise
		for dir, m = range pipe {
			break
		}
	} else {
		dir, m = pipe.moveFrom(g.cameFrom)
	}
	g.cameFrom = opposites[dir]
	return cell.plus(m)
}

func (g *pipeGrid) findCycleLength() {
	steps := 0
	for current := g.start; steps == 0 || !current.equals(g.start); steps++ {
		// fmt.Println(current)
		current = g.move(current)
	}
	// fmt.Println(steps)
	g.cycleLength = steps
}

func (g *pipeGrid) isInBounds(i, j int) bool {
	return i >= 0 && j >= 0 && i < g.height && j < g.width
}

func (g *pipeGrid) farthestDistance() int {
	// fmt.Println(g.cycleLength)
	return g.cycleLength/2 + g.cycleLength%2
}

// func (g pipeGrid) at(c pipeCell) string {
// 	if g.isInBounds(c[0], c[1]) {
// 		return g.rows[c[0]][c[1]]
// 	}
// 	return ""
// }

type DayTen struct{}

func (d DayTen) SolvePartOne(f string) int {
	lines := helpers.ReadLines(f)
	if strings.TrimSpace(lines[len(lines)-1]) == "" {
		lines = lines[:len(lines)-1]
	}
	grid := make([][]string, len(lines))
	for i, line := range lines {
		grid[i] = strings.Split(line, "")
	}
	g := newPipeGrid(grid)
	g.findCycleLength()
	return g.farthestDistance()
}

func (d DayTen) SolvePartTwo(f string) int {
	return -1
}
