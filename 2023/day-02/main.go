package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// task1()
	task2()
}

func task1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("An error occured: %v", err)
	}
	defer f.Close()

	maxColour := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	var count int
	var skip bool

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		skip = false
		l := strings.Split(line, ": ")

		sets := strings.Split(l[1], "; ")

		for _, set := range sets {
			cubes := strings.Split(set, ", ")
			for _, cube := range cubes {
				c := strings.Split(cube, " ")
				val, err := strconv.Atoi(c[0])
				if err != nil {
					fmt.Println(err)
				}

				if val > maxColour[c[1]] {
					skip = true
					break
				}
			}
			if skip {
				break
			}
		}
		if !skip {
			v, err := strconv.Atoi(strings.Split(l[0], " ")[1])
			if err != nil {
				fmt.Println(err)
			}
			count += v
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

	colour := make(map[string]int)
	var count int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		colour["blue"] = 0
		colour["green"] = 0
		colour["red"] = 0

		l := strings.Split(line, ": ")
		sets := strings.Split(l[1], "; ")

		for _, set := range sets {
			cubes := strings.Split(set, ", ")

			for _, cube := range cubes {
				c := strings.Split(cube, " ")
				val, err := strconv.Atoi(c[0])
				if err != nil {
					fmt.Println(err)
				}
				if val > colour[c[1]] {
					colour[c[1]] = val
				}
			}
		}
		count += colour["blue"] * colour["green"] * colour["red"]
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(count)
}
