package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil { panic(err) }

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	programs := make(map[int]program)

	for scanner.Scan() {
		line := scanner.Text()
		program := program{}
		index := strings.Index(line, " <-> ")
		id, line := toInt(line[:index]), line[index+5:]
		for line != "" {
			cid := 0
			index := strings.Index(line, ", ")
			if index != -1 {
				cid, line = toInt(line[:index]), line[index+2:]
			} else {
				cid, line = toInt(line), ""
			}
			program.connected = append(program.connected, cid)
		}
		programs[id] = program
	}

	visited := make(map[int]bool)
	visitAll(programs, visited, 0)

	fmt.Println("--- Part One ---")
	fmt.Println(len(visited))

	groups := 1
	for len(visited) < len(programs) {
		groups++
		start := 0
		for id, _ := range programs {
			if !visited[id] {
				start = id
				break
			}
		}
		visitAll(programs, visited, start)
	}

	fmt.Println("--- Part Two ---")
	fmt.Println(groups)
}

type program struct {
	connected []int
}

func visitAll(programs map[int]program, visited map[int]bool, start int) {
	open := []int{ start }
	for len(open) > 0 {
		next := make([]int, 0)
		for _, id := range open {
			if !visited[id] {
				visited[id] = true
				for _, cid := range programs[id].connected {
					next = append(next, cid)
				}
			}
		}
		open = next
	}
}

func toInt(v string) int {
	i, e := strconv.Atoi(v)
	if e != nil { panic(e) }
	return i
}
