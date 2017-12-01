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
	line := scanner.Text()
	
	fmt.Println(matchingSum(line, 1))
	fmt.Println(matchingSum(line, len(line) / 2))
}

func matchingSum(line string, offset int) int {
	sum := 0
	length := len(line)
	for i := 0; i < length; i++ {
		if line[i] == line[(i + offset) % length] {
			sum += toInt(line[i:i+1])
		}
	}
	return sum
}

func toInt(v string) int {
	i, e := strconv.Atoi(v)
	if e != nil { panic(e) }
	return i
}
