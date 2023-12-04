package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

type pos struct {
	y int // line number
	x int // character position in the line
}

var specialChars map[pos]bool

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

	specialChars = make(map[pos]bool)

	for i, l := range lines {
		for j, c := range l {
			if string(c) == "." {
				continue
			} else if unicode.IsDigit(c) {
				continue
			} else if string(c) == "*" {
				specialChars[pos{y: i, x: j}] = true
			} else {
				specialChars[pos{y: i, x: j}] = false
			}
		}
	}

	task1(lines)
	task2(lines)
}

func checkSpecialCharProyimity(start, end, lineNum, numOfLines, lineLength int) bool {
	if start < 0 {
		start = 0
	}
	if end > lineLength {
		end = lineLength
	}

	// line above
	if lineNum != 0 {
		for i := start; i <= end; i++ {
			_, ok := specialChars[pos{lineNum - 1, i}]
			if ok {
				return ok
			}
		}
	}

	// in the same line
	_, sOk := specialChars[pos{lineNum, start}]

	if sOk {
		return sOk
	}

	_, eOk := specialChars[pos{lineNum, end}]
	if eOk {
		return eOk
	}

	// below the line
	if lineNum != numOfLines-1 {
		for i := start; i <= end; i++ {
			_, ok := specialChars[pos{lineNum + 1, i}]
			if ok {
				return ok
			}
		}
	}

	return false
}

func task1(lines []string) {
	arrLength := len(lines)
	var sum int

	for k, line := range lines {
		lineLen := len(line)
		re := regexp.MustCompile(`\d+`)
		nums := re.FindAllStringIndex(line, -1)
		for _, num := range nums {
			if checkSpecialCharProyimity(num[0]-1, num[1], k, arrLength, lineLen) {
				v, err := strconv.Atoi(line[num[0]:num[1]])
				if err != nil {
					fmt.Println(err.Error())
				}
				sum += v
			}
		}
	}

	fmt.Println(sum)
}

func task2(lines []string) {
	var sum int

	for p, v := range specialChars {
		vals := []int{}

		if v {
			re := regexp.MustCompile(`\d+`)
			if p.y >= 1 {
				numsAbove := re.FindAllStringIndex(lines[p.y-1], -1)
				for _, num := range numsAbove {
					if (p.x-num[0] >= -1 && p.x-num[0] <= 1) || (p.x-num[1] >= -2 && p.x-num[1] <= 0) {
						val, err := strconv.Atoi(lines[p.y-1][num[0]:num[1]])
						if err != nil {
							fmt.Println(err.Error())
						}
						vals = append(vals, val)
					}
				}
			}

			numsLine := re.FindAllStringIndex(lines[p.y], -1)
			for _, num := range numsLine {
				if p.x-num[1] == 0 || p.x-num[0] == -1 {
					val, err := strconv.Atoi(lines[p.y][num[0]:num[1]])
					if err != nil {
						fmt.Println(err.Error())
					}
					vals = append(vals, val)
				}
			}

			if len(lines) > p.y+1 {
				numsBelow := re.FindAllStringIndex(lines[p.y+1], -1)

				for _, num := range numsBelow {
					if (p.x-num[0] >= -1 && p.x-num[0] <= 1) || (p.x-num[1] >= -2 && p.x-num[1] <= 0) {
						val, err := strconv.Atoi(lines[p.y+1][num[0]:num[1]])
						if err != nil {
							fmt.Println(err.Error())
						}
						vals = append(vals, val)
					}
				}
			}
		}

		if len(vals) == 2 {
			sum += vals[0] * vals[1]
		}
	}

	fmt.Println(sum)
}
