package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type node struct {
	name   string
	parent *node
	subDir []*node
	files  map[string]int
}

const (
	diskSize         = 70_000_000
	requiredFreeSize = 30_000_000
)

var root *node
var size int
var t2 map[string]int

func buildTree(lines []string, arrPos int, p *node, nm string) (int, *node) {
	i := arrPos
	dir := &node{
		parent: p,
		name:   nm,
	}

	if root == nil && nm == "/" {
		root = dir
	}
	if dir.files == nil {
		dir.files = make(map[string]int)
	}

	for i < len(lines) {
		line := strings.Split(lines[i], " ")
		i++
		if line[1] == "ls" || line[0] == "dir" {
			continue
		} else if line[0] == "$" && line[1] == "cd" && line[2] != ".." {
			x, n := buildTree(lines, i, dir, line[2])
			i = x
			dir.subDir = append(dir.subDir, n)
		} else if line[0] == "$" && line[1] == "cd" && line[2] == ".." {
			return i, dir
		} else {
			s, _ := strconv.Atoi(line[0])
			dir.files[line[1]] = s
		}
	}

	return i, dir
}

func task1(n node) int {
	dirSize := 0

	for _, v := range n.files {
		dirSize += v
	}

	if len(n.subDir) != 0 {
		for _, subDir := range n.subDir {
			dirSize += task1(*subDir)
		}
	}
	if dirSize <= 100_000 {
		// fmt.Printf("%v => %d\n", n.name, dirSize)
		size += dirSize
	}

	return dirSize
}

func usedDiskSize(n node) int {
	dirSize := 0

	for _, v := range n.files {
		dirSize += v
	}

	if len(n.subDir) != 0 {
		for _, subDir := range n.subDir {
			dirSize += usedDiskSize(*subDir)
		}
	}

	return dirSize
}

func task2(n node, needsToBeFreed int) int {
	dirSize := 0

	for _, v := range n.files {
		dirSize += v
	}

	if len(n.subDir) != 0 {
		for _, subDir := range n.subDir {
			dirSize += task2(*subDir, needsToBeFreed)
		}
	}
	if dirSize >= needsToBeFreed {
		t2[n.name] = dirSize
	}

	return dirSize
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

	buildTree(lines, 0, nil, "")
	task1(*root)
	fmt.Println(size)

	size = 0
	diskUsed := usedDiskSize(*root)
	needsToBeFreed := requiredFreeSize - (diskSize - diskUsed)

	t2 = make(map[string]int)
	task2(*root, needsToBeFreed)

	keys := make([]string, 0, len(t2))

	for key := range t2 {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return t2[keys[i]] < t2[keys[j]]
	})

	fmt.Println(keys)
}
