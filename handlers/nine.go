package handlers

import (
	"strings"
	"vyh/aoc23/internal/helpers"
)

type History struct {
	predictor  []int
	backfiller []int
}

func newHistory(line string) History {
	h := History{}
	derivatives := [][]int{helpers.ToInts(helpers.SplitAndTrim(line, " "))}
	for !helpers.AllZero(derivatives[len(derivatives)-1]) {
		row := derivatives[len(derivatives)-1]
		next := make([]int, len(row)-1)
		derivatives = append(derivatives, next)
		for i, j := 0, 1; j < len(row); i, j = i+1, j+1 {
			next[i] = row[j] - row[i]
		}
	}
	h.predictor = make([]int, len(derivatives)-1)
	h.backfiller = make([]int, len(derivatives)-1)
	for i, row := range derivatives[:len(derivatives)-1] {
		// ignore the all-zero row
		h.predictor[len(h.predictor)-i-1] = row[len(row)-1]
		h.backfiller[len(h.backfiller)-i-1] = row[0]
	}
	// fmt.Printf("%+v\n", h)
	return h
}

func (h History) next() int {
	for i, j := 0, 1; j < len(h.predictor); i, j = i+1, j+1 {
		h.predictor[j] += h.predictor[i]
	}
	return h.predictor[len(h.predictor)-1]
}

func (h History) previous() int {
	for i, j := 0, 1; j < len(h.backfiller); i, j = i+1, j+1 {
		h.backfiller[j] -= h.backfiller[i]
	}
	return h.backfiller[len(h.backfiller)-1]
}

type Histories []History

func (h Histories) next() []int {
	next := make([]int, len(h))
	for i := range h {
		next[i] = h[i].next()
	}
	return next
}

func (h Histories) previous() []int {
	previous := make([]int, len(h))
	for i := range h {
		previous[i] = h[i].previous()
	}
	return previous
}

type DayNine struct{}

func (d9 DayNine) SolvePartOne(f string) int {
	lines := helpers.ReadLines(f)
	histories := Histories{}
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		histories = append(histories, newHistory(line))
	}
	// fmt.Println(histories)
	next := histories.next()
	// fmt.Println(next)
	return helpers.Sum(next)
}

func (d9 DayNine) SolvePartTwo(f string) int {
	lines := helpers.ReadLines(f)
	histories := Histories{}
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		histories = append(histories, newHistory(line))
	}
	// fmt.Println(histories)
	previous := histories.previous()
	// fmt.Println(previous)
	return helpers.Sum(previous)
}
