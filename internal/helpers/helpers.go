package helpers

import (
	"fmt"
	"io"
	"math"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

// files

func ReadText(filename string) (string, error) {
	dir, _ := os.Getwd()
	fmt.Println(path.Join(dir, filename))
	r, err := os.Open(path.Join(dir, filename))
	if err != nil {
		return "", err
	}
	b, err := io.ReadAll(r)
	return string(b), err
}

func ReadLines(filename string) []string {
	content, err := ReadText(filename)
	if err != nil {
		log.Panicf("failed to read %s: %v", filename, err)
	}
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	return lines
}

// strings

func SplitAndTrim(s, delim string) []string {
	result := []string{}
	for _, part := range strings.Split(s, delim) {
		element := strings.TrimSpace(part)
		if len(element) > 0 {
			result = append(result, element)
		}
	}
	return result
}

func StringToDigit(s string) int {
	wordToInt := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	if n, ok := wordToInt[s]; ok {
		return n
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Error(err)
	}
	return n
}

func ReverseString(s string) string {
	rs := ""
	for _, c := range s {
		rs = string(c) + rs
	}
	return rs
}

// grid

func GetAdjacentCells(row int, colRange []int) [][]int {
	adjacentCells := [][]int{}
	for _, r := range []int{row - 1, row, row + 1} {
		if r < 0 {
			continue
		}
		for c := colRange[0] - 1; c < colRange[1]+1; c++ {
			if c < 0 {
				continue
			}
			if r == row && c >= colRange[0] && c < colRange[1] {
				continue
			}
			adjacentCells = append(adjacentCells, []int{r, c})
		}
	}
	return adjacentCells
}

func IsAdjacent(row, col int, cell []int) bool {
	return len(cell) == 2 && row == cell[0] && col == cell[1]
}

func AllMatch(p *regexp.Regexp, ss ...string) bool {
	for _, s := range ss {
		if !p.MatchString(s) {
			return false
		}
	}
	return true
}

func ToGrid(lines []string) [][]string {
	grid := [][]string{}
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		grid = append(grid, SplitAndTrim(line, ""))
	}
	return grid
}

func Transpose(grid [][]string) [][]string {
	w := len(grid[0])
	t := make([][]string, w)
	for i := range t {
		t[i] = make([]string, len(grid))
	}
	for i, row := range grid {
		if len(row) != w {
			log.Panic("grid not rectangular; unexpected row length", row)
		}
		for j, c := range row {
			t[j][i] = c
		}
	}
	return t
}

// int math

func Max(x, y int) int {
	if y > x {
		return y
	}
	return x
}

func Min(xs ...int) int {
	min := math.MaxInt
	for _, x := range xs {
		if x < min {
			min = x
		}
	}
	return min
}

func Sum(xs []int) int {
	sum := 0
	for _, x := range xs {
		sum += x
	}
	return sum
}

func Product(xs ...int) int {
	p := 1
	for _, x := range xs {
		p *= x
	}
	return p
}

func Pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func ToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return n
}

func ToInts(ss []string) []int {
	result := make([]int, len(ss))
	for i, s := range ss {
		result[i] = ToInt(s)
	}
	return result
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func AllZero(ns []int) bool {
	for _, n := range ns {
		if n != 0 {
			return false
		}
	}
	return true
}

func IntsEqual(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// misc

func Require(condition bool) {
	if !condition {
		log.Panic("failed condition")
	}
}
