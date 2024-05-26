package handlers

type Handler interface {
	SolvePartOne(string) int
	SolvePartTwo(string) int
}

func ForDay(day int) Handler {
	return map[int]Handler{
		6:  DaySix{},
		7:  DaySeven{},
		8:  DayEight{},
		9:  DayNine{},
		10: DayTen{},
		11: DayEleven{multiplier: 1000000},
		12: DayTwelve{},
		13: DayThirteen{},
		14: DayFourteen{},
		15: DayFifteen{},
	}[day]
}
