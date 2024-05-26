package handlers

import (
	"strconv"
	"strings"
	"vyh/aoc23/internal/helpers"

	log "github.com/sirupsen/logrus"
)

type Cubes struct {
	Red   int
	Green int
	Blue  int
}

func (c Cubes) isPossibleGame(total Cubes) bool {
	return c.Red <= total.Red && c.Green <= total.Green && c.Blue <= total.Blue
}

func (c Cubes) power() int {
	return c.Red * c.Green * c.Blue
}

type CubeGame []Cubes

func (g CubeGame) minTotal() Cubes {
	total := Cubes{}
	for _, cubes := range g {
		total.Red = helpers.Max(total.Red, cubes.Red)
		total.Green = helpers.Max(total.Green, cubes.Green)
		total.Blue = helpers.Max(total.Blue, cubes.Blue)
	}
	return total
}

func (g CubeGame) powerMin() int {
	return g.minTotal().power()
}

type DayTwo struct{}

func (d2 DayTwo) Solve(filename string, total Cubes) int {
	sum := 0
	for id, cubesGames := range d2.parse(filename) {
		possible := true
		for _, cubes := range cubesGames {
			if !cubes.isPossibleGame(total) {
				possible = false
				break
			}
		}
		if possible {
			sum += id
		}
	}
	return sum
}

func (d2 DayTwo) SolvePowerSum(filename string) int {
	sum := 0
	for _, game := range d2.parse(filename) {
		sum += game.powerMin()
	}
	return sum
}

func (d2 DayTwo) parse(filename string) map[int]CubeGame {
	result := map[int]CubeGame{}
	contents, err := helpers.ReadText(filename)
	if err != nil {
		log.Panic(err)
	}
	lines := strings.Split(contents, "\n")
	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		id, games := d2.parseGame(line)
		result[id] = games
	}
	return result
}

func (d2 DayTwo) parseGame(line string) (int, CubeGame) {
	parts := strings.Split(line, ":")
	id, err := strconv.Atoi(strings.TrimLeft(parts[0], "Game "))
	if err != nil {
		log.Panic(err)
	}
	game := CubeGame{}
	for _, outcome := range strings.Split(parts[1], ";") {
		game = append(game, d2.parseCubes(strings.TrimSpace(outcome)))
	}
	if len(game) == 0 {
		log.Panicf("game %d has 0 outcomes", id)
	}
	return id, game
}

func (d2 DayTwo) parseCubes(cubesStr string) Cubes {
	cubes := Cubes{}
	for _, color := range strings.Split(cubesStr, ",") {
		parts := strings.Split(strings.TrimSpace(color), " ")
		n, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			log.Panicf("failed to parse n for %v: %v", parts, err)
		}
		switch parts[1] {
		case "blue":
			cubes.Blue += n
		case "green":
			cubes.Green += n
		case "red":
			cubes.Red += n
		default:
			log.Panic("unrecognized color")
		}
	}
	return cubes
}
