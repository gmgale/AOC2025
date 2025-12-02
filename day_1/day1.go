package day1

import (
	"aoc2025/aocutils"
	"fmt"
)

// Solve runs both parts of day 1
func Solve() {
	lines := aocutils.ReadInput("day_1")
	println("Day 1:")
	println("Part 1:", part1(lines))
	println("Part 2:", part2(lines))
}

func part1(lines []string) int {
	// Make an array of 0-99 ints
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	cArray := newCircularArray(data)

	// Count how many times we land on 0
	zeroCount := 0

	for _, line := range lines {
		// Split the first character and the rest as an integer
		direction := line[0]
		var steps int
		fmt.Sscanf(line[1:], "%d", &steps)

		switch direction {
		case 'R':
			for i := 0; i < steps; i++ {
				cArray.next()
			}
		case 'L':
			for i := 0; i < steps; i++ {
				cArray.prev()
			}
		}

		if cArray.current() == 0 {
			zeroCount++
		}
	}
	return zeroCount
}

func part2(lines []string) int {
	// Make an array of 0-99 ints
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	cArray := newCircularArray(data)

	// Count how many times we land on 0
	zeroCount := 0

	for _, line := range lines {
		// Split the first character and the rest as an integer
		direction := line[0]
		var steps int
		fmt.Sscanf(line[1:], "%d", &steps)

		switch direction {
		case 'R':
			for i := 0; i < steps; i++ {
				cArray.next()
				if cArray.current() == 0 {
					zeroCount++
				}
			}
		case 'L':
			for i := 0; i < steps; i++ {
				cArray.prev()
				if cArray.current() == 0 {
					zeroCount++
				}
			}
		}

		//if cArray.current() == 0 {
		//	zeroCount++
		//}
	}
	return zeroCount
}

type circularArray struct {
	data []int
	pos  int
}

func newCircularArray(data []int) *circularArray {
	return &circularArray{data: data, pos: 50}
}

func (c *circularArray) next() int {
	value := c.data[c.pos]
	c.pos = (c.pos + 1) % len(c.data)
	return value
}

func (c *circularArray) prev() int {
	value := c.data[c.pos]
	c.pos = (c.pos - 1 + len(c.data)) % len(c.data)
	return value
}

func (c *circularArray) current() int {
	return c.data[c.pos]
}
