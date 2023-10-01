package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func task1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("An error occured: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		charArr := []rune(scanner.Text())

		for i := 0; i < len(charArr)-4; i++ {
			if (charArr[i] != charArr[i+1]) && (charArr[i] != charArr[i+2]) && (charArr[i] != charArr[i+3]) && (charArr[i+1] != charArr[i+2]) && (charArr[i+1] != charArr[i+3]) && (charArr[i+2] != charArr[i+3]) {
				fmt.Println(i + 4)
				break
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func task2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("An error occured: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		charArr := scanner.Text()
		for i := 0; i < len(charArr)-14; i++ {
			success := true
			tepCharArr := []rune(charArr[i : i+14])
			sort.Slice(tepCharArr, func(i, j int) bool {
				return tepCharArr[i] < tepCharArr[j]
			})

			for j := 0; j < len(tepCharArr)-1; j++ {
				if tepCharArr[j] == tepCharArr[j+1] {
					success = false
				}
			}

			if success {
				fmt.Println(i + 14)
				break
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	task1()
	task2()
}
