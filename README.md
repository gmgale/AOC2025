# Advent of Code 2025

## Project Structure

```
aoc2025/
├── main.go           # Entry point - runs solutions for any day
├── go.mod            # Go module definition
├── new_day.sh        # Script to generate boilerplate for new days
├── day_1/
│   ├── day1.go       # Solution code for day 1
│   └── input.txt     # Puzzle input for day 1
├── day_2/
│   ├── day2.go
│   └── input.txt
└── ...
```

## Usage

### Run a specific day
```bash
go run main.go -day 1
```

### Run with default (day 1)
```bash
go run main.go
```

### Build executable
```bash
go build -o aoc
./aoc -day 1
```

### Create a new day
```bash
./new_day.sh 2
```

This will:
- Create `day_2/` folder
- Generate `day2.go` with boilerplate
- Create empty `input.txt`

Then add the new day to the switch statement in `main.go`:
```go
case 2:
    day2.Solve()
```

## Adding A Solution

1. Put your puzzle input in `day_X/input.txt`
2. Implement `part1()` and `part2()` functions in `day_X/dayX.go`
3. Run with `go run main.go -day X`# AOC2025
