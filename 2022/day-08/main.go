package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func checkUp(arr [][]int, v, x, y int) bool {
	for i := x - 1; i >= 0; i-- {
		if arr[i][y] >= v {
			return true
		}
	}
	return false

}

func checkUpT2(arr [][]int, v, x, y int) int {
	count := 0
	for i := x - 1; i >= 0; i-- {
		count++
		if arr[i][y] >= v {
			break
		}
	}
	return count
}

func checkDown(arr [][]int, v, x, y int) bool {
	for i := x + 1; i < len(arr); i++ {
		if arr[i][y] >= v {
			return true
		}
	}
	return false
}

func checkDownT2(arr [][]int, v, x, y int) int {
	count := 0
	for i := x + 1; i < len(arr); i++ {
		count++
		if arr[i][y] >= v {
			break
		}
	}
	return count
}

func checkLeft(arr []int, v, y int) bool {
	for i := y - 1; i >= 0; i-- {
		if arr[i] >= v {
			return true
		}
	}
	return false
}

func checkLeftT2(arr []int, v, y int) int {
	count := 0
	for i := y - 1; i >= 0; i-- {
		count++
		if arr[i] >= v {
			break
		}
	}
	return count
}

func checkRight(arr []int, v, y int) bool {
	for i := y + 1; i < len(arr); i++ {
		if arr[i] >= v {
			return true
		}
	}
	return false
}

func checkRightT2(arr []int, v, y int) int {
	count := 0
	for i := y + 1; i < len(arr); i++ {
		count++
		if arr[i] >= v {
			break
		}
	}
	return count
}

func task1(arr [][]int) {
	count := 0
	arrLen := len(arr)

	for i := 0; i < arrLen; i++ {
		intArrLen := len(arr[0]) - 1
		for k, v := range arr[i] {
			if (i == 0) || (i == arrLen-1) {
				count++
				continue
			} else if (k == 0) || (k == intArrLen) {
				count++
				continue
			} else {
				if !checkUp(arr, v, i, k) || !checkDown(arr, v, i, k) || !checkLeft(arr[i], v, k) || !checkRight(arr[i], v, k) {
					count++
				}
			}
		}
	}

	fmt.Println(count)

}

func task2(arr [][]int) {
	count := 0
	arrLen := len(arr)

	for i := 0; i < arrLen; i++ {
		// intArrLen := len(arr[0]) - 1
		for k, v := range arr[i] {
			scenicScore := checkUpT2(arr, v, i, k) * checkDownT2(arr, v, i, k) * checkLeftT2(arr[i], v, k) * checkRightT2(arr[i], v, k)

			if scenicScore > count {
				count = scenicScore
			}
		}
	}
	fmt.Println(count)
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

	intArr := [][]int{}

	for _, line := range lines {
		arr := []int{}
		for _, v := range line {
			arr = append(arr, (int(v) - 48))
		}
		intArr = append(intArr, arr)
	}

	task1(intArr)
	task2(intArr)
}
