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
	input := toInt(scanner.Text())

	{
		fmt.Println("--- Part One ---")

		buffer := make([]int, 1, 2018)
		buffer[0] = 0

		pos := 0
		for i := 1; i <= 2017; i++ {
			pos = ((pos + input) % len(buffer)) + 1
			buffer = buffer[:len(buffer) + 1]
			for j := len(buffer) - 1; j > pos; j-- {
				buffer[j] = buffer[j-1]
			}
			buffer[pos] = i
		}

		fmt.Println(buffer[(pos + 1) % len(buffer)])
	}

	{
		fmt.Println("--- Part Two ---")

		data := make([]int, 1, bucketSize)
		data[0] = 0
		buckets := make([]bucket, 1)
		buckets[0] = bucket{ data }
		list := bucketList{ buckets, 1 }

		it := bucketIterator{ &list, 0, 0 }
		for i := 1; i <= 50000000; i++ {
			it.advance(input)
			it = list.insert(i, it)
		}

		it = bucketIterator{ &list, 0, 0 }
		for it.get() != 0 { it.advance(1) }
		it.advance(1)

		fmt.Println(it.get())
	}
}

const bucketSize = 4096

type bucketList struct {
	buckets []bucket
	size int
}

type bucket struct {
	data []int
}

type bucketIterator struct {
	list *bucketList
	bucketIndex, dataIndex int
}

func (list *bucketList) insert(value int, it bucketIterator) bucketIterator {
	list.size++
	if len(list.buckets[it.bucketIndex].data) == bucketSize {
		list.buckets = append(list.buckets, bucket{})
		for i := len(list.buckets) - 1; i > it.bucketIndex + 1; i-- {
			list.buckets[i] = list.buckets[i-1]
		}
		data := make([]int, bucketSize / 2, bucketSize)
		for i := 0; i < bucketSize / 2; i++ {
			data[i] = list.buckets[it.bucketIndex].data[bucketSize - bucketSize / 2 + i]
		}
		list.buckets[it.bucketIndex].data = list.buckets[it.bucketIndex].data[:bucketSize - bucketSize / 2]
		list.buckets[it.bucketIndex + 1] = bucket{ data }
		if it.dataIndex >= len(list.buckets[it.bucketIndex].data) {
			it.dataIndex -= len(list.buckets[it.bucketIndex].data)
			it.bucketIndex++
		}
	}
	bucket := list.buckets[it.bucketIndex]
	bucket.data = bucket.data[:len(bucket.data) + 1]
	it.dataIndex++
	for i := len(bucket.data) - 1; i > it.dataIndex; i-- {
		bucket.data[i] = bucket.data[i-1]
	}
	bucket.data[it.dataIndex] = value
	list.buckets[it.bucketIndex] = bucket
	return it
}

func (it *bucketIterator) advance(offset int) {
	offset = offset % it.list.size
	for it.dataIndex + offset >= len(it.list.buckets[it.bucketIndex].data) {
		offset -= len(it.list.buckets[it.bucketIndex].data)
		it.bucketIndex = (it.bucketIndex + 1) % len(it.list.buckets)
	}
	it.dataIndex += offset
}

func (it *bucketIterator) get() int {
	return it.list.buckets[it.bucketIndex].data[it.dataIndex]
}

func toInt(v string) int {
	i, e := strconv.Atoi(v)
	if e != nil { panic(e) }
	return i
}
