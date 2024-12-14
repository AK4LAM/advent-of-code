package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//Open text file
	file, err := os.Open("day1-input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	var left []int
	var right []int

	//Scanner to read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.Fields(scanner.Text())
		leftValue, _ := strconv.Atoi((text[0]))
		left = append(left, leftValue)
		rightValue, _ := strconv.Atoi((text[1]))
		right = append(right, rightValue)
	}
	rcopy := right
	sort.Ints(left)
	sort.Ints(right)

	difference := 0

	for i := 0; i < len(left) && i < len(right); i++ {
		if left[i] > right[i] {
			difference += left[i] - right[i]
		} else {
			difference += right[i] - left[i]
		}
	}
	fmt.Println("Difference: ", difference)

	count := make(map[int]int)
	for _, v := range rcopy {
		count[v] += 1
	}

	similarity := 0

	for _, v := range left {
		similarity += (v * count[v])
	}

	fmt.Println("Similarity Score: ", similarity)
}
