package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	task2()
}

func task() {
	f, err := os.Open("file.txt")

	if err != nil {
		log.Fatalf("An error occured: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var max int64
	var cur int64
	for scanner.Scan() {
		x := scanner.Text()
		if x == "" {
			if max < cur {
				max = cur
			}
			cur = 0
		} else {
			x_int, _ := strconv.ParseInt(x, 10, 64)
			cur += x_int
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(max)
}

func task2() {
	f, err := os.Open("file.txt")

	if err != nil {
		log.Fatalf("An error occured: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var max []int
	var cur int64
	for scanner.Scan() {
		x := scanner.Text()
		if x == "" {
			max = append(max, int(cur))
			cur = 0
		} else {
			x_int, _ := strconv.ParseInt(x, 10, 64)
			cur += x_int
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	sort.Slice(max, func(i, j int) bool {
		return max[i] > max[j]
	})
	fmt.Println(max[0] + max[1] + max[2])
}
