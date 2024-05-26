package main

import (
	"flag"
	"fmt"
	"vyh/aoc23/handlers"
)

func main() {
	day := flag.Int("day", 1, "day whose puzzle to run")
	// file := flag.String("file", "", "file to use for input")
	flag.Parse()

	// if len(*file) == 0 {
	// 	log.Panic("file required")
	// }
	file := fmt.Sprintf("internal/data/%d", *day)
	fmt.Printf("running %d\n", *day)
	switch *day {
	case 1:
		fmt.Println(handlers.GetCalibrationSum(file))
	case 2:
		total := handlers.Cubes{
			Red:   12,
			Green: 13,
			Blue:  14,
		}
		fmt.Println(handlers.DayTwo{}.Solve(file, total))
		fmt.Println(handlers.DayTwo{}.SolvePowerSum(file))
	case 3:
		fmt.Println(handlers.DayThree{}.Solve(file))
		fmt.Println(handlers.DayThree{}.SumGearRatios(file))
	case 4:
		fmt.Println(handlers.DayFour{}.Solve(file))
		fmt.Println(handlers.DayFour{}.CountCopies(file))
	case 5:
		fmt.Println(handlers.DayFive{}.LowestLocationNumber(file))
		fmt.Println(handlers.DayFive{}.ReverseSolve(file))
	default:
		h := handlers.ForDay(*day)
		fmt.Println(h.SolvePartOne(file))
		fmt.Println(h.SolvePartTwo(file))
	}
}
