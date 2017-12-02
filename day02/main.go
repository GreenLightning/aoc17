package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func main() {
	file, err := os.Open("input.txt")
	if err != nil { panic(err) }

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var data [][]int
	for scanner.Scan() {
		line := scanner.Text()
		var values []int
		for line != "" {
			var value int
			value, line = grabInt(line)
			values = append(values, value)
		}
		data = append(data, values)
	}

	{
		fmt.Println("--- Part One ---")
		checkSum := 0
		for _, values := range data {
			min, max := MaxInt, MinInt
			for _, value := range values {
				if value > max { max = value }
				if value < min { min = value }
			}
			checkSum += max - min
		}
		fmt.Println(checkSum)
	}

	{
		fmt.Println("--- Part Two ---")
		checkSum := 0
		for _, values := range data {
			for i, a := range values {
				for j, b := range values {
					if i != j && a % b == 0 {
						checkSum += a / b
					}
				}
			}
		}
		fmt.Println(checkSum)
	}
}

func grabInt(line string) (int, string) {
	index := strings.Index(line, "\t")
	if index == -1 {
		return toInt(line), ""
	} else {
		return toInt(line[:index]), line[index+1:]
	}
}

func toInt(v string) int {
	i, e := strconv.Atoi(v)
	if e != nil { panic(e) }
	return i
}
