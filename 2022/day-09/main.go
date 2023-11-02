package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	x int
	y int
}

func newTail(h, t pos) pos {
	nt := t
	diff := pos{h.x - t.x, h.y - t.y}

	if diff.x > 1 {
		nt.x = h.x - 1
		nt.y = h.y
	} else if diff.x < -1 {
		nt.x = h.x + 1
		nt.y = h.y
	}
	if diff.y > 1 {
		nt.x = h.x
		nt.y = h.y - 1
	} else if diff.y < -1 {
		nt.x = h.x
		nt.y = h.y + 1
	}

	return nt
}

func task1(strArr []string) {
	went := make(map[pos]bool)
	went[pos{x: 0, y: 0}] = true
	head := pos{x: 0, y: 0}
	tail := pos{x: 0, y: 0}

	for _, v := range strArr {
		val := strings.Split(v, " ")

		direction := val[0]
		steps, _ := strconv.Atoi(val[1])

		switch direction {
		case "U":
			for i := 0; i < steps; i++ {
				head.y += 1
				tail = newTail(head, tail)
				went[tail] = true
			}
		case "R":
			for i := 0; i < steps; i++ {
				head.x += 1
				tail = newTail(head, tail)
				went[tail] = true
			}
		case "D":
			for i := 0; i < steps; i++ {
				head.y -= 1
				tail = newTail(head, tail)
				went[tail] = true
			}
		case "L":
			for i := 0; i < steps; i++ {
				head.x -= 1
				tail = newTail(head, tail)
				went[tail] = true
			}
		}
	}
	fmt.Println(len(went))
}

func task2(strArr []string) {
	went := make(map[pos]bool)
	went[pos{x: 0, y: 0}] = true

	r := struct {
		tail [10]pos
	}{
		tail: [10]pos{},
	}

	for _, v := range strArr {
		val := strings.Split(v, " ")

		direction := val[0]
		steps, _ := strconv.Atoi(val[1])

		switch direction {
		case "U":
			for i := 0; i < steps; i++ {
				r.tail[0].y += 1
				for j := 0; j <= 8; j++ {
					r.tail[j+1] = newTail(r.tail[j], r.tail[j+1])
				}
				went[r.tail[9]] = true
			}
		case "R":
			for i := 0; i < steps; i++ {
				r.tail[0].x += 1
				for j := 0; j <= 8; j++ {
					r.tail[j+1] = newTail(r.tail[j], r.tail[j+1])
				}
				went[r.tail[9]] = true
			}
		case "D":
			for i := 0; i < steps; i++ {
				r.tail[0].y -= 1
				for j := 0; j <= 8; j++ {
					r.tail[j+1] = newTail(r.tail[j], r.tail[j+1])
				}
				went[r.tail[9]] = true
			}
		case "L":
			for i := 0; i < steps; i++ {
				r.tail[0].x -= 1
				for j := 0; j <= 8; j++ {
					r.tail[j+1] = newTail(r.tail[j], r.tail[j+1])
				}
				went[r.tail[9]] = true
			}
		}
	}
	fmt.Println(len(went))

}

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

	// task2(lines)
}
