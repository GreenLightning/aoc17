package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	programs := parsePrograms("input.txt")
	connectPrograms(programs)
	root := findRoot(programs)

	fmt.Println("--- Part One ---")
	fmt.Println(root.name)

	fmt.Println("--- Part Two ---")
	calculateTotalWeight(root)
	checkBalance(root, 0)
}

type program struct {
	name string
	weight int
	childNames []string

	parent *program
	children []*program
	totalWeight int
}

func parsePrograms(filename string) []*program {
	programs := make([]*program, 0)

	file, err := os.Open(filename)
	if err != nil { panic(err) }

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	programRule := regexp.MustCompile(`(\w+) \((\d+)\)(?: -> (.*))?`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := programRule.FindStringSubmatch(line)

		prog := &program{}
		prog.name = matches[1]
		prog.weight = toInt(matches[2])

		childList := matches[3]
		for childList != "" {
			index := strings.Index(childList, ", ")
			childName := ""
			if index == -1 {
				childName, childList = childList, ""
			} else {
				childName, childList = childList[:index], childList[index+2:]
			}
			prog.childNames = append(prog.childNames, childName)
		}

		programs = append(programs, prog)
	}

	return programs
}

func connectPrograms(programs []*program) {
	programsByName := make(map[string]*program)
	for _, prog := range programs {
		programsByName[prog.name] = prog
	}
	for _, prog := range programs {
		for _, name := range prog.childNames {
			child := programsByName[name]
			child.parent = prog
			prog.children = append(prog.children, child)
		}
	}
}

func findRoot(programs []*program) *program {
	root := programs[0]
	for root.parent != nil {
		root = root.parent
	}
	return root
}

func calculateTotalWeight(prog *program) int {
	weight := prog.weight
	for _, child := range prog.children {
		weight += calculateTotalWeight(child)
	}
	prog.totalWeight = weight
	return weight
}

func checkBalance(prog *program, expectedWeight int) {
	switch len(prog.children) {
	case 0:
		if prog.weight != expectedWeight {
			fmt.Println(expectedWeight)
		}

	case 1:
		checkBalance(prog.children[0], expectedWeight - prog.weight)

	default:
		childWeight := prog.children[0].totalWeight
		allSame := true
		for i := 1; i < len(prog.children); i++ {
			if prog.children[i].totalWeight != childWeight {
				allSame = false
				break
			}
		}
		if allSame {
			fmt.Println(expectedWeight - len(prog.children) * childWeight)
		} else {
			// If we cannot calculate the expected weight using the child
			// weight we got from our first child, that means that this first
			// child must be imbalanced itself, so we get the correct child
			// weight from our second child.
			if prog.weight + len(prog.children) * childWeight != expectedWeight {
				childWeight = prog.children[1].totalWeight
			}
			for _, child := range prog.children {
				if child.totalWeight != childWeight  {
					checkBalance(child, childWeight)
				}
			}
		}
	}
}

func toInt(v string) int {
	i, e := strconv.Atoi(v)
	if e != nil { panic(e) }
	return i
}
