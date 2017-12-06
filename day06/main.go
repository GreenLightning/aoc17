package main

import (
	"bufio"
	"bytes"
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

	scanner.Scan()
	input := scanner.Text()

	data := make([]int, 0)

	for input != "" {
		var value int
		value, input = grabInt(input)
		data = append(data, value)
	}

	count := 0
	counts := make(map[string]int)
	state := stringify(data)
	length := len(data)
	for {
		if _, ok := counts[state]; ok {
			break
		}
		counts[state] = count
		max, index := data[0], 0
		for i := 1; i < length; i++ {
			if data[i] > max {
				max, index = data[i], i
			}
		}
		data[index] = 0
		for i := (index + 1) % length; max > 0; i = (i + 1) % length {
			data[i]++
			max--
		}
		state = stringify(data)
		count++
	}

	fmt.Println("--- Part One ---")
	fmt.Println(count)

	fmt.Println("--- Part Two ---")
	fmt.Println(count - counts[state])
}

func stringify(data []int) string {
	var buffer bytes.Buffer
	length := len(data)
	if length > 0 {
		buffer.WriteString(strconv.Itoa(data[0]))
	}
	for i := 1; i < length; i++ {
		buffer.WriteRune(' ')
		buffer.WriteString(strconv.Itoa(data[i]))
	}
	return buffer.String()
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
