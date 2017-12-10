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

	scanner.Scan()
	input := scanner.Text()

	{
		fmt.Println("--- Part One ---")
		lengths := make([]int, 0)
		for _, str := range strings.Split(input, ",") {
			lengths = append(lengths, toInt(str))
		}
		list := makeList()
		doOneRound(list, lengths, 0, 0)
		fmt.Println(list[0] * list[1])
	}

	{
		fmt.Println("--- Part Two ---")
		lengths := make([]int, 0)
		for _, b := range []byte(input) {
			lengths = append(lengths, int(b))
		}
		lengths = append(lengths, 17, 31, 73, 47, 23)
		list := makeList()
		current, skip := 0, 0
		for i := 0; i < 64; i++ {
			current, skip = doOneRound(list, lengths, current, skip)
		}
		result := ""
		for i := 0; i < 16; i++ {
			value := 0
			for j := 0; j < 16; j++ {
				value ^= list[16 * i + j]
			}
			result += fmt.Sprintf("%02x", value)
		}
		fmt.Println(result)
	}
}

func makeList() []int {
	list := make([]int, 256)
	for i, _ := range list {
		list[i] = i
	}
	return list
}

func doOneRound(list, lengths []int, current, skip int) (int, int) {
	for _, length := range lengths {
		for i := 0; i < length / 2; i++ {
			j1 := (current + i) & 0xff
			j2 := (current + length - 1 - i) & 0xff
			list[j1], list[j2] = list[j2], list[j1]
		}
		current = (current + length + skip) & 0xff
		skip++
	}
	return current, skip
}

func toInt(v string) int {
	i, e := strconv.Atoi(v)
	if e != nil { panic(e) }
	return i
}
