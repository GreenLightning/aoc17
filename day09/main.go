package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil { panic(err) }

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	line := scanner.Text()

	_, score, count := parseGroup(line, 1)

	fmt.Println("--- Part One ---")
	fmt.Println(score)

	fmt.Println("--- Part Two ---")
	fmt.Println(count)
}

func parseElement(line string, score int) (string, int, int) {
	if line[0] == '{' {
		return parseGroup(line, score)
	} else {
		return parseGarbage(line)
	}
}

func parseGroup(line string, score int) (string, int, int) {
	total, count := score, 0
	line = line[1:]
	for line[0] != '}' {
		var s, c int
		line, s, c = parseElement(line, score + 1)
		total += s
		count += c
		if line[0] == ',' {
			line = line[1:]
		}
	}
	line = line[1:]
	return line, total, count
}

func parseGarbage(line string) (string, int, int) {
	count := 0
	line = line[1:]
	for line[0] != '>' {
		if line[0] == '!' {
			line = line[2:]
		} else {
			line = line[1:]
			count++
		}
	}
	line = line[1:]
	return line, 0, count
}
