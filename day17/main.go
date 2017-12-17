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

		buffer := make([]int, 1, 2018)
		buffer[0] = 0

		pos := 0
		for i := 1; i <= 2017; i++ {
			pos = ((pos + input) % len(buffer)) + 1
			buffer = buffer[:len(buffer) + 1]
			for j := len(buffer) - 1; j > pos; j-- {
				buffer[j] = buffer[j-1]
			}
			buffer[pos] = i
		}

		fmt.Println(buffer[(pos + 1) % len(buffer)])
	}

	{
		fmt.Println("--- Part Two ---")

		pos, data, length := 0, -1, 1
		for i := 1; i <= 50000000; i++ {
			pos = ((pos + input) % length) + 1
			if pos == 1 { data = i }
			length++
		}

		fmt.Println(data)
	}
}

func toInt(v string) int {
	i, e := strconv.Atoi(v)
	if e != nil { panic(e) }
	return i
}
