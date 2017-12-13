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

	layers := make(map[int]*layer)
	totalDepth := -1

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), ": ")
		depth, size := toInt(values[0]), toInt(values[1])
		layers[depth] = &layer{ size, 0, 1 }
		if depth > totalDepth { totalDepth = depth }
	}

	fmt.Println("--- Part One ---")

	severity := 0

	for depth := 0; depth <= totalDepth; depth++ {
		if layer, found := layers[depth]; found {
			if layer.pos == 0 {
				severity += depth * layer.size
			}
		}

		for _, layer := range layers {
			layer.advance()
		}
	}

	fmt.Println(severity)

	// Reset layers.
	for _, layer := range layers  {
		layer.pos, layer.dir = 0, 1
	}

	fmt.Println("--- Part Two ---")

	// Skew layers.
	for depth, layer := range layers {
		for i := 0; i < depth; i++ {
			layer.advance()
		}
	}

	delay := 0

	for {
		good := true
		for _, layer := range layers {
			if layer.pos == 0 {
				good = false
				break
			}
		}
		if good {
			break
		}
		for _, layer := range layers {
			layer.advance()
		}
		delay++
	}

	fmt.Println(delay)
}

type layer struct {
	size, pos, dir int
}

func (layer *layer) advance() {
	layer.pos += layer.dir
	if layer.dir == 1 && layer.pos == layer.size - 1 {
		layer.dir = -1
	} else if layer.dir == -1 && layer.pos == 0 {
		layer.dir = 1
	}
}

func toInt(v string) int {
	i, e := strconv.Atoi(v)
	if e != nil { panic(e) }
	return i
}
