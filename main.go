package main

import (
	"flag"
	"fmt"
	"os"

	"aoc2025/day_1"
)

func main() {
	day := flag.Int("day", 1, "Which day to run (1-25)")
	flag.Parse()

	switch *day {
	case 1:
		day1.Solve()
	default:
		fmt.Fprintf(os.Stderr, "Day %d not implemented yet\n", *day)
		os.Exit(1)
	}
}
