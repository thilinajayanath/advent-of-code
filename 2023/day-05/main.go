package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
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

type container struct {
	dest     int
	src      int
	srcRange int
}

type req struct {
	name string
	val  int
}

func doStuff(lines []string, seeds []int) {
	almanac := make(map[string][]container)
	seedReq := make(map[int][]req)

	var key string
	keys := []string{}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		if unicode.IsLetter(rune(line[0])) {
			splittedLine := strings.Split(line, " ")
			key = strings.Split(splittedLine[0], "-")[2]
			keys = append(keys, key)
		} else {
			re := regexp.MustCompile(`\d+`)
			vals := re.FindAllString(line, -1)

			c := container{}
			c.dest, _ = strconv.Atoi(vals[0])
			c.src, _ = strconv.Atoi(vals[1])
			c.srcRange, _ = strconv.Atoi(vals[2])
			almanac[key] = append(almanac[key], c)
		}
	}

	var tmp int
	for _, seed := range seeds {
		tmp = seed
		for _, key := range keys {
			val := tmp
			for _, c := range almanac[key] {
				if tmp >= c.src && tmp < c.src+c.srcRange {
					val = tmp - c.src + c.dest
				}
			}

			seedReq[seed] = append(seedReq[seed], req{name: key, val: val})
			tmp = val
		}
	}

	var min int
	for _, sq := range seedReq {
		for _, r := range sq {
			if r.name == "location" {
				if min == 0 {
					min = r.val
				} else if min > r.val {
					min = r.val
				}
			}
		}
	}

	fmt.Println(min)
}

func task1(lines []string) {
	seeds := []int{}

	splittedLine := strings.Split(lines[0], " ")
	if splittedLine[0] == "seeds:" {
		for i := 1; i < len(splittedLine); i++ {
			val, _ := strconv.Atoi(splittedLine[i])
			seeds = append(seeds, val)
		}
	}
	doStuff(lines[2:], seeds)
}

func task2(lines []string) {
	numbers := strings.Split(strings.Split(lines[0], ": ")[1], " ")
	seedArr := []int{}
	for _, n := range numbers {
		val, _ := strconv.Atoi(n)
		seedArr = append(seedArr, val)
	}
	fmt.Println(len(seedArr))
	seeds := []int{}
	for i := 0; i < len(seedArr)-1; i += 2 {
		for j := seedArr[i]; j < seedArr[i]+seedArr[i+1]; j++ {
			seeds = append(seeds, j)
		}
	}
	fmt.Println(len(seeds))
	// doStuff(lines[2:], seeds)
	// revisit later
}
