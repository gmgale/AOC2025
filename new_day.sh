#!/bin/bash

if [ -z "$1" ]; then
    echo "Usage: ./new_day.sh <day_number>"
    exit 1
fi

DAY=$1
DIR="day_$DAY"

if [ -d "$DIR" ]; then
    echo "Day $DAY already exists!"
    exit 1
fi

mkdir -p "$DIR"
touch "$DIR/input.txt"

cat > "$DIR/day$DAY.go" << EOF
package day$DAY

import (
	"aoc2025/aocutils"
)

// Solve runs both parts of day $DAY
func Solve() {
	lines := aocutils.ReadInput("day_$DAY")
	println("Day $DAY:")
	println("Part 1:", part1(lines))
	println("Part 2:", part2(lines))
}

func part1(lines []string) int {
	// TODO: Implement part 1
	return 0
}

func part2(lines []string) int {
	// TODO: Implement part 2
	return 0
}
EOF

echo "Created day_$DAY/"
echo "Don't forget to:"
echo "1. Add your input to day_$DAY/input.txt"
echo "2. Add the case to main.go switch statement"