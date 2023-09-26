package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

func main() {
	task1()
	task2()
}

var valueMap = map[rune]int{
	'a': 1,
	'b': 2,
	'c': 3,
	'd': 4,
	'e': 5,
	'f': 6,
	'g': 7,
	'h': 8,
	'i': 9,
	'j': 10,
	'k': 11,
	'l': 12,
	'm': 13,
	'n': 14,
	'o': 15,
	'p': 16,
	'q': 17,
	'r': 18,
	's': 19,
	't': 20,
	'u': 21,
	'v': 22,
	'w': 23,
	'x': 24,
	'y': 25,
	'z': 26,
	'A': 27,
	'B': 28,
	'C': 29,
	'D': 30,
	'E': 31,
	'F': 32,
	'G': 33,
	'H': 34,
	'I': 35,
	'J': 36,
	'K': 37,
	'L': 38,
	'M': 39,
	'N': 40,
	'O': 41,
	'P': 42,
	'Q': 43,
	'R': 44,
	'S': 45,
	'T': 46,
	'U': 47,
	'V': 48,
	'W': 49,
	'X': 50,
	'Y': 51,
	'Z': 52,
}

func task1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var re_items []rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items := []rune(scanner.Text())
		length := len(items) / 2
		items1 := items[:length]
		items2 := items[length:]
		var tempItems []rune
		for _, val := range items1 {
			if slices.Contains(items2, val) {
				if !slices.Contains(tempItems, val) {
					tempItems = append(tempItems, val)
				}
			}
		}
		re_items = append(re_items, tempItems...)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var sum int = 0
	for _, v := range re_items {
		sum += valueMap[v]
	}
	fmt.Println(sum)
}

func task2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var groupItems []rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items1 := []rune(scanner.Text())
		scanner.Scan()
		items2 := []rune(scanner.Text())
		scanner.Scan()
		items3 := []rune(scanner.Text())

		for _, val := range items1 {
			if slices.Contains(items2, val) && slices.Contains(items3, val) {
				groupItems = append(groupItems, val)
				break
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var sum int = 0
	for _, v := range groupItems {
		sum += valueMap[v]
		fmt.Println(string(v))
	}
	fmt.Println(sum)
}
