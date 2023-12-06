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

	task1(&lines)
	task2(&lines)

}

func task1(lines *[]string) {
	time := []int{}
	distance := []int{}

	for _, line := range *lines {
		splittedLine := strings.Split(line, ":")
		if splittedLine[0] == "Time" {
			re := regexp.MustCompile(`\d+`)
			vals := re.FindAllString(splittedLine[1], -1)
			for _, val := range vals {
				v, _ := strconv.Atoi(val)
				time = append(time, v)
			}
		} else if splittedLine[0] == "Distance" {
			re := regexp.MustCompile(`\d+`)
			vals := re.FindAllString(splittedLine[1], -1)
			for _, val := range vals {
				v, _ := strconv.Atoi(val)
				distance = append(distance, v)
			}
		}
	}
	sum := 1
	count := 0
	for i := 0; i < len(time); i++ {
		for j := 0; j <= time[i]; j++ {
			velocity := 1 * j
			dis := velocity * (time[i] - j)
			if dis > distance[i] {
				count++
			}
		}
		sum *= count
		count = 0
	}
	fmt.Println(sum)
}

func task2(lines *[]string) {
	var time int
	var distance int

	for _, line := range *lines {
		splittedLine := strings.Split(line, ":")
		if splittedLine[0] == "Time" {
			re := regexp.MustCompile(`\d+`)
			vals := re.FindAllString(splittedLine[1], -1)
			val := ""
			for _, v := range vals {
				val += v
			}
			v, _ := strconv.Atoi(val)
			time = v

		} else if splittedLine[0] == "Distance" {
			re := regexp.MustCompile(`\d+`)
			vals := re.FindAllString(splittedLine[1], -1)
			val := ""
			for _, v := range vals {
				val += v
			}
			v, _ := strconv.Atoi(val)
			distance = v
		}
	}

	count := 0
	for i := 0; i <= time; i++ {

		dis := i * (time - i)
		if dis > distance {
			count++
		}
	}

	fmt.Println(count)
}
