package handlers

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"vyh/aoc23/internal/helpers"
)

func ch(s string, i int) string {
	return string([]byte{s[i]})
}

type springLine struct {
	springsStr string
	// springs      []string
	groupLengths []int
	pattern      *regexp.Regexp
}

func (l springLine) isSpringSolution(s string) bool {
	result := l.pattern.MatchString(s)
	// fmt.Println("is ", s, " a match: ", result)
	return result
}

func strReplace(i int, s, c string) string {
	return s[:i] + c + s[i+1:]
}

// may just have to go recursive - try it, fail out when the tried thing does not work
func (l springLine) springSolutions(s string, i int, groupsRemaining []int) int {
	if len(groupsRemaining) == 0 {
		for ; i < len(s); i++ {
			if ch(s, i) == "?" {
				s = strReplace(i, s, ".")
			}
		}
	}
	if i >= len(s) {
		if l.isSpringSolution(s) {
			return 1
		}
		return 0
	}

	switch ch(s, i) {
	case "#":
		j := 1
		for ; j < groupsRemaining[0] && i+j < len(s); j++ {
			idx := i + j
			switch ch(s, idx) {
			case "#":
				continue
			case "?":
				s = strReplace(idx, s, "#") // hello inefficient
			default:
				return 0
			}
		}
		if i+j < len(s) {
			char := ch(s, i+j)
			if char == "#" {
				return 0
			} else if char == "?" {
				s = strReplace(i+j, s, ".")
				j++
			}
		}
		return l.springSolutions(s, i+j, groupsRemaining[1:])
	case ".":
		return l.springSolutions(s, i+1, groupsRemaining)
	case "?":
		return l.springSolutions(strReplace(i, s, "."), i, groupsRemaining) +
			l.springSolutions(strReplace(i, s, "#"), i, groupsRemaining)
	default:
		log.Panic("unexpected character in ", s, ch(s, i))
	}

	return 0
}

func newSpringLine(s string, multiply bool) springLine {
	parts := helpers.SplitAndTrim(s, " ")
	groupLengths := helpers.ToInts(helpers.SplitAndTrim(parts[1], ","))
	s = strings.TrimSpace(parts[0])
	if multiply {
		s = fmt.Sprintf("%s?%s?%s?%s?%s", s, s, s, s, s)
		groups := groupLengths[:]
		for i := 1; i < 5; i++ {
			groups = append(groups, groupLengths...)
		}
		helpers.Require(len(groups) == 5*len(groupLengths))
		groupLengths = groups
	}

	ps := `^\.*`
	for i, gl := range groupLengths {
		if i < len(groupLengths)-1 {
			ps = ps + fmt.Sprintf("(#{%d})", gl) + `\.+`
		} else {
			ps = ps + fmt.Sprintf("(#{%d})", gl) + `\.*$`
		}
	}
	// rejoin in case of spaces
	return springLine{s, groupLengths, regexp.MustCompile(ps)}
}

type DayTwelve struct{}

func (d DayTwelve) SolvePartOne(f string) int {
	lines := helpers.ReadLines(f)
	springLines := []springLine{}
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		springLines = append(springLines, newSpringLine(line, false))
	}
	// fmt.Println(springLines)
	counts := make([]int, len(springLines))
	for i, sl := range springLines {
		counts[i] = sl.springSolutions(sl.springsStr, 0, sl.groupLengths)
	}
	// fmt.Println(counts)
	return helpers.Sum(counts)
}

func (d DayTwelve) SolvePartTwo(f string) int {
	lines := helpers.ReadLines(f)
	springLines := []springLine{}
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		springLines = append(springLines, newSpringLine(line, true))
	}
	// fmt.Println(springLines)
	counts := make([]int, len(springLines))
	for i, sl := range springLines {
		counts[i] = sl.springSolutions(sl.springsStr, 0, sl.groupLengths)
		fmt.Println("line ", i, " count ", counts[i])
	}
	// fmt.Println(counts)
	return helpers.Sum(counts)
}
