package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read file content
	data, err := os.ReadFile("day2-input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Process file content into slices of integers
	valued := parseData(string(data))
	safe := 0
	//valued := [][]int{
	//	{1, 3, 6, 7, 9},
	//	{8, 6, 4, 4, 1},
	//	{1, 3, 2, 4, 5},
	//	{9, 7, 6, 2, 1},
	//	{1, 2, 7, 8, 9},
	//	{7, 6, 4, 2, 1},
	//}

	for _, v := range valued {
		// Check if the report is already safe
		if isSafe(v) {
			safe++
			continue
		}

		// Try removing one element at a time and check if it becomes safe
		for i := 0; i < len(v); i++ {
			temp := append([]int{}, v[:i]...) // Copy elements before index i
			temp = append(temp, v[i+1:]...)   // Copy elements after index i

			if isSafe(temp) {
				safe++
				break
			}
		}
	}

	fmt.Println(safe)
}

// Check if a sequence is safe (either all increasing or all decreasing)
func isSafe(values []int) bool {
	return allIncreasing(values) || allDecreasing(values)
}

func allIncreasing(values []int) bool {
	if len(values) < 2 {
		return false
	}

	for i := 0; i < len(values)-1; i++ {
		diff := values[i+1] - values[i]
		if diff < 1 || diff > 3 { // Difference must be between 1 and 3
			return false
		}
		if values[i] >= values[i+1] { // Must always increase
			return false
		}
	}
	return true
}

func allDecreasing(values []int) bool {
	if len(values) < 2 {
		return false
	}

	for i := 0; i < len(values)-1; i++ {
		diff := values[i] - values[i+1]
		if diff < 1 || diff > 3 { // Difference must be between 1 and 3
			return false
		}
		if values[i] <= values[i+1] { // Must always decrease
			return false
		}
	}
	return true
}

// Parse the file content into [][]int
func parseData(data string) [][]int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	var valued [][]int

	for _, line := range lines {
		numStrs := strings.Fields(line) // Split line into numbers
		var nums []int

		for _, numStr := range numStrs {
			num, err := strconv.Atoi(numStr) // Convert string to int
			if err != nil {
				fmt.Println("Error converting number:", err)
				continue
			}
			nums = append(nums, num)
		}

		valued = append(valued, nums)
	}

	return valued
}
