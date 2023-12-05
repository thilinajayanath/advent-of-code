package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
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

	// task1(lines)
	// task2(lines)
	t2alternative(lines)
}

type container struct {
	dest     int
	src      int
	srcRange int
}

func doStuff(lines []string, seeds []int) int {
	almanac := make(map[string][]container)

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
	var min int
	for _, seed := range seeds {
		tmp = seed
		for _, key := range keys {
			val := tmp
			for _, c := range almanac[key] {
				if tmp >= c.src && tmp < c.src+c.srcRange {
					val = tmp - c.src + c.dest
				}
			}

			tmp = val
			if key == "location" {
				if min == 0 {
					min = val
				} else if min > val {
					min = val
				}
			}
		}
	}

	return min
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
	fmt.Println(doStuff(lines[2:], seeds))
}

func t2Helper(start int, end int, lines []string, wg *sync.WaitGroup, ch chan int) {
	seeds := []int{}
	for i := start; i < start+end; i++ {
		seeds = append(seeds, i)
	}

	ch <- doStuff(lines, seeds)
	wg.Done()
}

// consumed about 43 GB of memory + paged memory when running this
func task2(lines []string) {
	numbers := strings.Split(strings.Split(lines[0], ": ")[1], " ")
	seedArr := []int{}
	for _, n := range numbers {
		val, _ := strconv.Atoi(n)
		seedArr = append(seedArr, val)
	}

	var wg sync.WaitGroup
	ch := make(chan int)
	defer close(ch)

	threadCount := 0
	for i := 0; i < len(seedArr)-1; i += 2 {
		go t2Helper(seedArr[i], seedArr[i+1], lines[2:], &wg, ch)
		threadCount++
		wg.Add(1)
	}

	var minValue int
	for v := range ch {
		threadCount--
		if minValue == 0 {
			minValue = v
		} else if minValue > v {
			minValue = v
		}

		if threadCount == 0 {
			break
		}
	}
	wg.Wait()

	fmt.Println(minValue)
}

// alternative solution for task 2 which uses much less memory

func doStuffAlt(lines *[]string, keys *[]string, almanac map[string][]container, seed int) int {
	var tmp int

	tmp = seed
	for _, key := range *keys {
		val := tmp
		for _, c := range almanac[key] {
			if tmp >= c.src && tmp < c.src+c.srcRange {
				val = tmp - c.src + c.dest
			}
		}

		tmp = val
		if key == "location" {
			return val
		}
	}

	return -1
}

func t2HelperAlt(start int, end int, lines *[]string, keys *[]string, almanac map[string][]container, wg *sync.WaitGroup, ch chan int) {
	min := 0
	for i := start; i < start+end; i++ {
		val := doStuffAlt(lines, keys, almanac, i)
		if min == 0 {
			min = val
		} else if min > val {
			min = val
		}
	}

	ch <- min
	wg.Done()
}

// This uses 1.7 MB of memory
func t2alternative(lines []string) {
	numbers := strings.Split(strings.Split(lines[0], ": ")[1], " ")
	seedArr := []int{}
	for _, n := range numbers {
		val, _ := strconv.Atoi(n)
		seedArr = append(seedArr, val)
	}

	almanac := make(map[string][]container)
	newLines := lines[2:]

	var key string
	keys := []string{}
	for _, line := range newLines {
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

	var wg sync.WaitGroup
	ch := make(chan int)
	defer close(ch)

	threadCount := 0
	for i := 0; i < len(seedArr)-1; i += 2 {
		go t2HelperAlt(seedArr[i], seedArr[i+1], &newLines, &keys, almanac, &wg, ch)
		threadCount++
		wg.Add(1)
	}

	var minValue int
	for v := range ch {
		threadCount--
		if minValue == 0 {
			minValue = v
		} else if minValue > v {
			minValue = v
		}

		if threadCount == 0 {
			break
		}
	}
	wg.Wait()

	fmt.Println(minValue)
}
