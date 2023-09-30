package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func task1(arr [][]string) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("An error occured: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var re = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	operations := [][]int{}

	for scanner.Scan() {
		x := scanner.Text()
		match := re.FindStringSubmatch(x)
		amount, _ := strconv.Atoi(match[1])
		startIndex, _ := strconv.Atoi(match[2])
		endIndex, _ := strconv.Atoi(match[3])
		operations = append(operations, []int{
			amount,
			startIndex - 1,
			endIndex - 1,
		})
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	for _, v := range operations {
		for i := 0; i < v[0]; i++ {
			x := arr[v[1]][len(arr[v[1]])-1]
			arr[v[1]] = arr[v[1]][:len(arr[v[1]])-1]
			arr[v[2]] = append(arr[v[2]], x)
		}
	}

	fmt.Println(arr)
}

func task2(arr [][]string) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("An error occured: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var re = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	operations := [][]int{}

	for scanner.Scan() {
		x := scanner.Text()
		match := re.FindStringSubmatch(x)
		amount, _ := strconv.Atoi(match[1])
		startIndex, _ := strconv.Atoi(match[2])
		endIndex, _ := strconv.Atoi(match[3])
		operations = append(operations, []int{
			amount,
			startIndex - 1,
			endIndex - 1,
		})
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	for _, v := range operations {
		amount := v[0]
		startIndex := v[1]
		endIndex := v[2]

		x := arr[startIndex][len(arr[startIndex])-amount:]
		arr[startIndex] = arr[startIndex][:len(arr[startIndex])-amount]
		arr[endIndex] = append(arr[endIndex], x...)
	}

	fmt.Println(arr)
}

func main() {
	x := [][]string{
		{"F", "C", "J", "P", "H", "T", "W"},
		{"G", "R", "V", "F", "Z", "J", "B", "H"},
		{"H", "P", "T", "R"},
		{"Z", "S", "N", "P", "H", "T"},
		{"N", "V", "F", "Z", "H", "J", "C", "D"},
		{"P", "M", "G", "F", "W", "D", "Z"},
		{"M", "V", "Z", "W", "S", "J", "D", "P"},
		{"N", "D", "S"},
		{"D", "Z", "S", "F", "M"},
	}
	y := [][]string{
		{"F", "C", "J", "P", "H", "T", "W"},
		{"G", "R", "V", "F", "Z", "J", "B", "H"},
		{"H", "P", "T", "R"},
		{"Z", "S", "N", "P", "H", "T"},
		{"N", "V", "F", "Z", "H", "J", "C", "D"},
		{"P", "M", "G", "F", "W", "D", "Z"},
		{"M", "V", "Z", "W", "S", "J", "D", "P"},
		{"N", "D", "S"},
		{"D", "Z", "S", "F", "M"},
	}

	task1(x)
	task2(y)
}
