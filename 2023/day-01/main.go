package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	task1()
	task2()
}

func task1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("An error occured: %v", err)
	}
	defer f.Close()

	var sumOfValues int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		x := scanner.Text()
		cv := []rune{}
		for _, c := range x {
			v := c - 48
			if v < 10 {
				cv = append(cv, v)
			}
		}
		cvLen := len(cv)
		if cvLen == 1 {
			sumOfValues += int(cv[0]*10 + cv[0])
		} else {
			sumOfValues += int(cv[0]*10 + cv[cvLen-1])
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(sumOfValues)
}

func task2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("An error occured: %v", err)
	}
	defer f.Close()

	numberMap := make(map[string]int)

	numberMap["one"] = 1
	numberMap["two"] = 2
	numberMap["three"] = 3
	numberMap["four"] = 4
	numberMap["five"] = 5
	numberMap["six"] = 6
	numberMap["seven"] = 7
	numberMap["eight"] = 8
	numberMap["nine"] = 9

	var sumOfValues int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		str := scanner.Text()
		cv := []rune{}
		match := []int{}

		re := regexp.MustCompile(`^one|^two|^three|^four|^five|^six|^seven|^eight|^nine`)
		for k, c := range str {

			xLen := len(str)
			v := c - 48
			if v < 10 {
				cv = append(cv, v)
				continue
			}

			if xLen-k >= 5 {
				match = re.FindStringIndex(str[k : k+5])
			} else {
				match = re.FindStringIndex(str[k:])
			}
			if match == nil {
				continue
			} else {
				cv = append(cv, rune(numberMap[str[k+match[0]:k+match[1]]]))
			}
		}
		cvLen := len(cv)
		if cvLen == 1 {
			sumOfValues += int(cv[0]*10 + cv[0])
		} else {
			sumOfValues += int(cv[0]*10 + cv[cvLen-1])
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(sumOfValues)
}
