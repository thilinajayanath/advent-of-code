package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
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

	task1(lines)
	task2(lines)
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

func t2helper(start int, end int, lines []string, wg *sync.WaitGroup, ch chan int) {
	seeds := []int{}
	for i := start; i < start+end; i++ {
		seeds = append(seeds, i)
	}

	ch <- doStuff(lines, seeds)
	wg.Done()
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

// consumed about 40 GB of memory + paged memory when running this
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
		go t2helper(seedArr[i], seedArr[i+1], lines[2:], &wg, ch)
		threadCount++
		wg.Add(1)
	}

	minValues := []int{}
	for v := range ch {
		threadCount--
		minValues = append(minValues, v)
		if threadCount == 0 {
			break
		}
	}
	wg.Wait()
	sort.Slice(minValues, func(i, j int) bool {
		return minValues[i] < minValues[j]
	})
	fmt.Println(minValues[0])
}
