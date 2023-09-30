package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	task1()
	task2()
}

func task1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var sum int = 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		x := strings.Split(scanner.Text(), " ")
		if x[0] == "A" {
			if x[1] == "X" {
				sum += 3 + 1
			} else if x[1] == "Y" {
				sum += 6 + 2
			} else if x[1] == "Z" {
				sum += 3
			}
		} else if x[0] == "B" {
			if x[1] == "X" {
				sum += 1
			} else if x[1] == "Y" {
				sum += 3 + 2
			} else if x[1] == "Z" {
				sum += 6 + 3
			}
		} else if x[0] == "C" {
			if x[1] == "X" {
				sum += 6 + 1
			} else if x[1] == "Y" {
				sum += 2
			} else if x[1] == "Z" {
				sum += 3 + 3
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}

func task2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var sum int = 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		x := strings.Split(scanner.Text(), " ")
		if x[0] == "A" {
			if x[1] == "X" {
				sum += 3
			} else if x[1] == "Y" {
				sum += 3 + 1
			} else if x[1] == "Z" {
				sum += 6 + 2
			}
		} else if x[0] == "B" {
			if x[1] == "X" {
				sum += 1
			} else if x[1] == "Y" {
				sum += 3 + 2
			} else if x[1] == "Z" {
				sum += 6 + 3
			}
		} else if x[0] == "C" {
			if x[1] == "X" {
				sum += 2
			} else if x[1] == "Y" {
				sum += 3 + 3
			} else if x[1] == "Z" {
				sum += 6 + 1
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}
