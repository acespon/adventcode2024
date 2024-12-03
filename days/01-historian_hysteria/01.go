package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var first, second []int
	var differences int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Fields(line)
		if len(split) != 2 {
			fmt.Println("Invalid line format:", line)
			return
		}

		f, err := strconv.Atoi(split[0])
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}
		s, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}

		first = append(first, f)
		second = append(second, s)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	slices.Sort(first)
	slices.Sort(second)

	for i := 0; i < len(first); i++ {
		differences += int(math.Abs(float64(first[i] - second[i])))
	}

	fmt.Println("Part One:", differences)

	// Part Two
	// Create an array that is length largest number of the first list
	// For each number in the first list, increment the value at that index
	var second2 = make([]int, slices.Max(first)+1)
	var similarityScore int

	for i := 0; i < len(first); i++ {
		var number = second[i]
		second2[number]++
	}

	// For each number in the first list, multiply the value at that index by the value at the same index in the second list
	for i := 0; i < len(first); i++ {
		var number = first[i]
		similarityScore += first[i] * second2[number]
	}

	println("Part Two:", similarityScore)
}
