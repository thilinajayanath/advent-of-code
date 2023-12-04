package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
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

	task1(lines)
	task2(lines)
}

func pow(n int) int {
	if n == 1 {
		return 1
	} else if n < 0 {
		return 0
	}

	return 2 * pow(n-1)
}

func task1(lines []string) {
	var sum int
	for _, line := range lines {
		winningNumbers := make(map[int]bool)
		splittedStr := strings.Split(line, ": ")
		numbers := strings.Split(splittedStr[1], " | ")
		re := regexp.MustCompile(`\d+`)
		winningNums := re.FindAllString(numbers[0], -1)

		for _, n := range winningNums {
			v, err := strconv.Atoi(n)
			if err != nil {
				fmt.Println(err.Error())
			}
			winningNumbers[v] = true
		}

		lotteryNums := re.FindAllString(numbers[1], -1)

		var count int
		for _, n := range lotteryNums {
			v, err := strconv.Atoi(n)
			if err != nil {
				fmt.Println(err.Error())
			}
			_, ok := winningNumbers[v]

			if ok {
				count++
			}
		}
		x := pow(count)
		sum += x

	}

	fmt.Println(sum)
}

func task2(lines []string) {
	processCount := make(map[int]int)

	for _, line := range lines {
		winningNumbers := make(map[int]bool)
		splittedStr := strings.Split(line, ": ")
		numbers := strings.Split(splittedStr[1], " | ")
		re := regexp.MustCompile(`\d+`)

		ticketNum, err := strconv.Atoi(re.FindString(splittedStr[0]))
		if err != nil {
			fmt.Println(err.Error())
		}
		processCount[ticketNum]++

		winningNums := re.FindAllString(numbers[0], -1)

		for _, n := range winningNums {
			v, err := strconv.Atoi(n)
			if err != nil {
				fmt.Println(err.Error())
			}
			winningNumbers[v] = true
		}

		lotteryNums := re.FindAllString(numbers[1], -1)

		var count int

		for _, n := range lotteryNums {
			v, err := strconv.Atoi(n)
			if err != nil {
				fmt.Println(err.Error())
			}
			_, ok := winningNumbers[v]

			if ok {
				count++
				processCount[ticketNum+count] += processCount[ticketNum]
			}
		}
	}

	sum := 0
	for _, v := range processCount {
		sum += v
	}

	fmt.Println(sum)
}
