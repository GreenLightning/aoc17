package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const hextable = "0123456789abcdef"

func main() {
	file, err := os.Open("input.txt")
	if err != nil { panic(err) }

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	input := scanner.Text()

	grid := [128][128]int{}

	{
		regionId := 1
		for row := 0; row < 128; row++ {
			hash := calculateKnotHash(fmt.Sprintf("%s-%d", input, row))
			for baseCol, hexchar := range hash {
				val := strings.IndexRune(hextable, hexchar)
				for i := 0; i < 4; i++ {
					if (val >> uint(i)) & 1 == 1 {
						grid[row][4 * baseCol + 4 - i - 1] = regionId
						regionId++
					}
				}
			}
		}
	}

	{
		fmt.Println("--- Part One ---")

		count := 0

		for row := 0; row < 128; row++ {
			for col := 0; col < 128; col++ {
				if grid[row][col] > 0 {
					count++
				}
			}
		}

		fmt.Println(count)
	}

	{
		fmt.Println("--- Part Two ---")

		regions := make(map[int][]point)

		for row := 0; row < 128; row++ {
			for col := 0; col < 128; col++ {
				if grid[row][col] > 0 {
					regions[grid[row][col]] = []point{ point{ row, col } }
				}
			}
		}

		changed := true
		for changed {
			changed = false
			for row := 0; row + 1 < 128; row++ {
				for col := 0; col < 128; col++ {
					if grid[row][col] > 0 && grid[row+1][col] > 0 && grid[row][col] != grid[row+1][col] {
						merge(&grid, regions, row, col, row+1, col)
						changed = true
					}
				}
			}
			for row := 0; row < 128; row++ {
				for col := 0; col + 1 < 128; col++ {
					if grid[row][col] > 0 && grid[row][col+1] > 0 && grid[row][col] != grid[row][col+1] {
						merge(&grid, regions, row, col, row, col+1)
						changed = true
					}
				}
			}
		}

		fmt.Println(len(regions))
	}
}

type point struct {
	row, col int
}

func merge(grid *[128][128]int, regions map[int][]point, row1, col1, row2, col2 int) {
	regionId1, regionId2 := grid[row1][col1], grid[row2][col2]
	region1, region2 := regions[regionId1], regions[regionId2]
	for _, p := range region2 {
		region1 = append(region1, p)
		grid[p.row][p.col] = regionId1
	}
	regions[regionId1] = region1
	delete(regions, regionId2)
}

func calculateKnotHash(input string) string {
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
	return result
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
