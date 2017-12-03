package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil { panic(err) }

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	input := toInt(scanner.Text())

	{
		fmt.Println("--- Part One ---")
		pos := 1
		x, y := loop(input, func (x, y, input int) bool {
			pos++
			return pos < input
		})
		fmt.Println(abs(x) + abs(y))
	}

	{
		fmt.Println("--- Part Two ---")
		const n = 1000
		grid := [2*n+1][2*n+1]int{}
		grid[n][n] = 1
		x, y := loop(input, func (x, y, input int) bool {
			x, y = x+n, y+n
			grid[x][y] = grid[x-1][y] + grid[x+1][y] +
			             grid[x][y-1] + grid[x][y+1] +
			             grid[x-1][y-1] + grid[x-1][y+1] +
			             grid[x+1][y-1] + grid[x+1][y+1]
			return grid[x][y] <= input
		})
		fmt.Println(grid[x+n][y+n])
	}
}

func loop(input int, function func(int,int,int)bool) (int, int) {
	x, y := 0, 0
	step := 1
	for {
		for j := 0; j < step; j++ {
			x += 1
			if !function(x, y, input) {
				return x, y
			}
		}
		for j := 0; j < step; j++ {
			y += 1
			if !function(x, y, input) {
				return x, y
			}
		}
		step++
		for j := 0; j < step; j++ {
			x -= 1
			if !function(x, y, input) {
				return x, y
			}
		}
		for j := 0; j < step; j++ {
			y -= 1
			if !function(x, y, input) {
				return x, y
			}
		}
		step++
	}
}

func toInt(v string) int {
	i, e := strconv.Atoi(v)
	if e != nil { panic(e) }
	return i
}

func abs(v int) int {
	if v < 0 { return -v }
	return v
}
