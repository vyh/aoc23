package handlers

import (
	"strings"
	"vyh/aoc23/internal/helpers"
)

type Card struct {
	id             int
	count          int
	winningNumbers map[int]struct{}
	presentNumbers map[int]struct{}
}

func newCard(s string) *Card {
	parts := helpers.SplitAndTrim(s, ":")
	id := helpers.ToInt(strings.TrimSpace(strings.TrimPrefix(parts[0], "Card")))
	parts = helpers.SplitAndTrim(parts[1], "|")
	winningNumbers := helpers.SplitAndTrim(parts[0], " ")
	winMap := map[int]struct{}{}
	for _, n := range winningNumbers {
		winMap[helpers.ToInt(n)] = struct{}{}
	}
	presentNumbers := helpers.SplitAndTrim(parts[1], " ")
	presentMap := map[int]struct{}{}
	for _, n := range presentNumbers {
		presentMap[helpers.ToInt(n)] = struct{}{}
	}
	return &Card{
		id:             id,
		count:          1,
		winningNumbers: winMap,
		presentNumbers: presentMap,
	}
}

func (c *Card) getWins() []int {
	wins := []int{}
	for n := range c.presentNumbers {
		if _, ok := c.winningNumbers[n]; ok {
			wins = append(wins, n)
		}
	}
	return wins
}

func (c *Card) getValue() int {
	return helpers.Pow(2, len(c.getWins())-1)
}

func (c *Card) winCopies(cards Cards) {
	for i := range c.getWins() {
		k := c.id + i + 1
		if copy, ok := cards[k]; ok {
			copy.count += c.count
		}
	}
	if next, ok := cards[c.id+1]; ok {
		next.winCopies(cards)
	}
}

type Cards map[int]*Card

func (c Cards) getValues() []int {
	values := make([]int, len(c))
	i := 0
	for _, card := range c {
		values[i] = card.getValue()
		i++
	}
	return values
}

func (c Cards) getTotalValue() int {
	return helpers.Sum(c.getValues())
}

func (c Cards) count() int {
	sum := 0
	for _, card := range c {
		sum += card.count
	}
	return sum
}

type DayFour struct{}

func (d4 DayFour) Solve(filename string) int {
	lines := helpers.ReadLines(filename)
	cards := make(Cards, len(lines))
	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		card := newCard(line)
		cards[card.id] = card
		// log.Info("card ", cards[i].id, " value: ", cards[i].getValue())
	}
	return cards.getTotalValue()
}

func (d4 DayFour) CountCopies(filename string) int {
	lines := helpers.ReadLines(filename)
	cards := make(Cards, len(lines))
	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		card := newCard(line)
		cards[card.id] = card
		// log.Info("card ", cards[i].id, " value: ", cards[i].getValue())
	}
	cards[1].winCopies(cards)
	return cards.count()
}
