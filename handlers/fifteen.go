package handlers

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"vyh/aoc23/internal/helpers"
)

// (label)(=|-)(focal length)
var opPattern = regexp.MustCompile(`^([a-z]+)(=|-)([0-9]+)?$`)

func lavaHash(s string) int {
	hash := 0
	for _, ch := range s {
		hash = ((hash + int(ch)) * 17) % 256
	}
	return hash
}

type lens struct {
	id          int
	focalLength int
}

type lenses []*lens

func (l lenses) lessThan(i, j int) bool {
	if l[j] == nil { // put nil slots at the end
		return true
	}
	if l[i] == nil {
		return false
	}
	return l[i].id < l[j].id
}

func (l lenses) sort() {
	sort.SliceStable(l, l.lessThan)
}

func (l lenses) focusingPower(idx int) int {
	sum := 0
	for i, obj := range l {
		if obj == nil {
			break
		}
		power := (idx + 1) * (i + 1) * obj.focalLength
		sum += power
		// fmt.Println("Box", idx, "lens", i, "focal length", obj.focalLength, "power", power)
	}
	return sum
}

func (l lenses) String() string {
	s := ""
	for i, slot := range l {
		s += fmt.Sprintf("lens %d: %+v\n", i, slot)
	}
	return s
}

type instruction struct {
	hash        int
	label       string
	op          string
	focalLength int
}

func newInstruction(s string) instruction {
	parts := opPattern.FindAllStringSubmatch(s, -1)[0][1:]
	focalLength := 0
	var err error
	if parts[1] == "=" {
		focalLength, err = strconv.Atoi(parts[2])
		helpers.Require(err == nil)
	}
	return instruction{
		hash:        lavaHash(parts[0]),
		label:       parts[0],
		op:          parts[1],
		focalLength: focalLength,
	}
}

type lensBox map[string]*lens

type DayFifteen struct{}

func (d DayFifteen) SolvePartOne(f string) int {
	lines := helpers.SplitAndTrim(helpers.ReadLines(f)[0], ",")
	hashes := []int{}
	for _, line := range lines {
		hashes = append(hashes, lavaHash(line))
	}
	return helpers.Sum(hashes)
}

func (d DayFifteen) SolvePartTwo(f string) int {
	lines := helpers.SplitAndTrim(helpers.ReadLines(f)[0], ",")
	boxes := make([]lensBox, 256)
	for i := range boxes {
		boxes[i] = lensBox{}
	}
	for i, line := range lines {
		op := newInstruction(line)
		if op.op == "-" {
			// hash == applicable box
			// - == remove from box if present & shift other lenses up
			delete(boxes[op.hash], op.label)
		} else if op.op == "=" {
			// = == set/override VALUE (but not POSITION if already there)
			if l, ok := boxes[op.hash][op.label]; ok {
				l.focalLength = op.focalLength
			} else {
				boxes[op.hash][op.label] = &lens{
					id:          i,
					focalLength: op.focalLength,
				}
			}
		}

	}
	// then get the values and sort by id at the end
	p := 0
	for i, b := range boxes {
		box := lenses{}
		for _, l := range b {
			if l != nil {
				// fmt.Println(label, l)
				box = append(box, l)
			}
		}
		if len(box) > 0 {
			box.sort()
			// fmt.Printf("%+v\n", box)
			p += box.focusingPower(i)
		}
	}
	return p
}
