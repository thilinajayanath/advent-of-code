package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var cardRanking map[string]int
var handRanking map[string]int

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatalf("An error occured: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	cardRanking = map[string]int{
		"2": 0,
		"3": 1,
		"4": 2,
		"5": 3,
		"6": 4,
		"7": 5,
		"8": 6,
		"9": 7,
		"T": 8,
		"J": 9,
		"Q": 10,
		"K": 11,
		"A": 12,
	}

	handRanking = map[string]int{
		"high":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"full":  4,
		"four":  5,
		"five":  6,
	}

	task1(&lines)
	// task2(&lines)

}

type node struct {
	hand string
	bet  int
	next *node
	prev *node
}

var start *node
var itemCount int

func classify(hand []int) string {
	handType := ""
	if hand[0] == 5 {
		handType = "five"
	} else if hand[0] == 4 {
		handType = "four"
	} else if hand[0] == 3 {
		if hand[1] == 2 {
			handType = "full"
		} else {
			handType = "three"
		}
	} else if hand[0] == 2 {
		if hand[1] == 2 {
			handType = "two"
		} else {
			handType = "one"
		}
	} else {
		handType = "high"
	}
	return handType
}

func t1CompareHands(hand1 string, hand2 string) bool {
	hand1Char := make(map[rune]int)
	hand2Char := make(map[rune]int)

	for _, c := range hand1 {
		hand1Char[c]++
	}
	for _, c := range hand2 {
		hand2Char[c]++
	}

	hand1Nums := []int{}
	hand2Nums := []int{}

	for _, v := range hand1Char {
		hand1Nums = append(hand1Nums, v)
	}

	sort.Slice(hand1Nums, func(i, j int) bool {
		return hand1Nums[i] > hand1Nums[j]
	})

	for _, v := range hand2Char {
		hand2Nums = append(hand2Nums, v)
	}

	sort.Slice(hand2Nums, func(i, j int) bool {
		return hand2Nums[i] > hand2Nums[j]
	})
	h1Classification := classify(hand1Nums)
	h2Classification := classify(hand2Nums)

	if h1Classification != h2Classification {
		if handRanking[h1Classification] > handRanking[h2Classification] {
			fmt.Println("hand1 better than hand 1", hand1, hand2)
			return true
		} else {
			fmt.Println("hand2 better than hand 1", hand1, hand2)
			return false
		}
	} else {
		fmt.Println(hand1, hand2)
		for i := 0; i < 5; i++ {
			if hand1[i] != hand2[i] {
				check := cardRanking[string(hand1[i])] > cardRanking[string(hand2[i])]
				fmt.Println(i, check, hand1, hand2)

				return check
			}
		}
	}
	return false
}

func t1Insert(h string, b int) {
	n := node{
		hand: h,
		bet:  b,
	}

	if start == nil {
		start = &n
	} else {
		if t1CompareHands(start.hand, h) {
			start.prev = &n
			n.next = start
			start = &n
		} else {
			tmpNode := start
			for {
				fmt.Println(h)

				check := t1CompareHands(h, tmpNode.hand)

				if !check {
					n.next = tmpNode
					n.prev = tmpNode.prev
					tmpNode.prev.next = &n
					tmpNode.prev = &n

					break
				} else if tmpNode.next == nil {
					tmpNode.next = &n
					break
				}
				tmpNode = tmpNode.next
			}
		}
	}
	itemCount++
}

func t1Print() {
	if start == nil {
		return
	}
	tmpNode := start
	for {
		fmt.Println(*tmpNode)
		if tmpNode.next == nil {
			break
		}
		tmpNode = tmpNode.next
	}
}

func task1(lines *[]string) {
	for _, line := range *lines {
		splittedLine := strings.Split(line, " ")
		bet, _ := strconv.Atoi(splittedLine[1])
		t1Insert(splittedLine[0], bet)
	}
	t1Print()
}
