package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func task1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("An error occured: %v", err)
	}
	defer f.Close()

	count := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		x := scanner.Text()
		input := strings.Split(x, ",")

		elfOne := strings.Split(input[0], "-")
		elfTwo := strings.Split(input[1], "-")
		elfOneStart, _ := strconv.ParseInt(elfOne[0], 10, 64)
		elfOneEnd, _ := strconv.ParseInt(elfOne[1], 10, 64)
		elfTwoStart, _ := strconv.ParseInt(elfTwo[0], 10, 64)
		elfTwoEnd, _ := strconv.ParseInt(elfTwo[1], 10, 64)

		if (elfOneStart <= elfTwoStart && elfOneEnd >= elfTwoEnd) || (elfTwoStart <= elfOneStart && elfTwoEnd >= elfOneEnd) {
			count++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(count)
}

func task2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("An error occured: %v", err)
	}
	defer f.Close()

	count := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		x := scanner.Text()
		input := strings.Split(x, ",")

		elfOne := strings.Split(input[0], "-")
		elfTwo := strings.Split(input[1], "-")
		elfOneStart, _ := strconv.ParseInt(elfOne[0], 10, 64)
		elfOneEnd, _ := strconv.ParseInt(elfOne[1], 10, 64)
		elfTwoStart, _ := strconv.ParseInt(elfTwo[0], 10, 64)
		elfTwoEnd, _ := strconv.ParseInt(elfTwo[1], 10, 64)

		if (elfOneStart <= elfTwoStart && elfOneEnd >= elfTwoStart) || (elfTwoStart <= elfOneEnd && elfTwoEnd >= elfOneStart) {
			count++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(count)
}

func main() {
	task1()
	task2()
}
