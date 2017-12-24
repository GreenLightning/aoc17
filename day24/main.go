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

	components := make([]component, 0)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "/")
		components = append(components, component{ [2]int{ toInt(line[0]), toInt(line[1]) } })
	}

	max, longestMax, _ := getBestBridges(components, 0)

	fmt.Println("--- Part One ---")
	fmt.Println(max)

	fmt.Println("--- Part Two ---")
	fmt.Println(longestMax)
}

type component struct {
	ports [2]int
}

func getBestBridges(components []component, port int) (max int, longestMax int, longestLength int) {
	max, longestMax, longestLength = 0, 0, 0
	for c, comp := range components {
		for p := 0; p < 2; p++ {
			if comp.ports[p] == port {
				newComponents := make([]component, len(components) - 1)
				for i := 0; i < c; i++ {
					newComponents[i] = components[i]
				}
				for i := c + 1; i < len(components); i++ {
					newComponents[i - 1] = components[i]
				}
				newPort := comp.ports[1 - p]
				current, longestCurrent, currentLength := getBestBridges(newComponents, newPort)
				current += comp.ports[0] + comp.ports[1]
				longestCurrent += comp.ports[0] + comp.ports[1]
				currentLength++
				if current > max {
					max = current
				}
				if currentLength >= longestLength && longestCurrent > longestMax {
					longestMax, longestLength = longestCurrent, currentLength
				}
			}
		}
	}
	return max, longestMax, longestLength
}

func toInt(v string) int {
	i, e := strconv.Atoi(v)
	if e != nil { panic(e) }
	return i
}
