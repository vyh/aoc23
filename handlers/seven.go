package handlers

import (
	"sort"
	"strings"
	"vyh/aoc23/internal/helpers"
)

var CardValues = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

type Hand struct {
	cards []int
	kind  float64 // <highest # matching>.<2nd highest>
	bid   int
}

func newHand(line string, joker bool) Hand {
	if joker {
		CardValues["J"] = 1
	}
	parts := helpers.SplitAndTrim(line, " ")
	helpers.Require(len(parts) == 2)
	cards := make([]int, 5)
	counts := map[int]int{}
	for i, c := range strings.Split(parts[0], "") {
		cards[i] = CardValues[c]
		counts[cards[i]] += 1
	}
	helpers.Require(cards[4] > 0)
	first := 0.0
	second := 0.0
	for v, count := range counts {
		if v == 1 {
			continue // don't count joker in initial hand value
		}
		c := float64(count)
		if c > first {
			first, second = c, first
		} else if c > second {
			second = c
		}
	}
	if jokers, ok := counts[1]; joker && ok {
		// I think it's always strictly better to increase the first #
		// as much as possible; (3,2 or 2,2): 4,1 > 3,2; 3,1 > 2,2
		first += float64(jokers)
	}
	bid := helpers.ToInt(parts[1])
	return Hand{
		cards: cards,
		kind:  first + second/10.0,
		bid:   bid,
	}
}

type Hands []Hand

func (h Hands) lessThan(i, j int) bool {
	if h[i].kind == h[j].kind {
		a, b := h[i].cards, h[j].cards
		for card := 0; card < 5; card++ {
			if a[card] == b[card] {
				continue
			}
			return a[card] < b[card]
		}
		return false
	}
	return h[i].kind < h[j].kind
}

func (h Hands) sort() {
	sort.SliceStable(h, h.lessThan)
}

func (h Hands) rank(i int) int {
	return i + 1
}

func (h Hands) winnings() []int {
	// assumes hands have already been sorted
	winnings := make([]int, len(h))
	for i, hand := range h {
		winnings[i] = hand.bid * h.rank(i)
	}
	return winnings
}

type DaySeven struct{}

func (d7 DaySeven) SolvePartOne(filename string) int {
	lines := helpers.ReadLines(filename)
	hands := Hands{}
	for _, line := range lines {
		if len(strings.TrimSpace(line)) > 0 {
			hands = append(hands, newHand(line, false))
		}
	}
	hands.sort()
	// fmt.Printf("hands: %+v\n", hands)
	winnings := hands.winnings()
	// fmt.Printf("winnings: %+v\n", winnings)
	totalWinnings := helpers.Sum(winnings)
	return totalWinnings
}

func (d7 DaySeven) SolvePartTwo(filename string) int {
	lines := helpers.ReadLines(filename)
	hands := Hands{}
	for _, line := range lines {
		if len(strings.TrimSpace(line)) > 0 {
			hands = append(hands, newHand(line, true))
		}
	}
	hands.sort()
	// fmt.Printf("hands: %+v\n", hands)
	winnings := hands.winnings()
	// fmt.Printf("winnings: %+v\n", winnings)
	totalWinnings := helpers.Sum(winnings)
	return totalWinnings
}
