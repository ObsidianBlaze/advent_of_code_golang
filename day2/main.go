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
	//valued := [][]int{{1, 3, 6, 7, 9}, {8, 6, 4, 4, 1}, {1, 3, 2, 4, 5}, {9, 7, 6, 2, 1}, {1, 2, 7, 8, 9}, {7, 6, 4, 2, 1}}
	for _, v := range valued {
		//Check if all values are being incremented or decremented
		if allIncreasing(v) || allDecreasing(v) {
			safe += 1
		}
	}
	fmt.Println(safe)

}

func allIncreasing(values []int) bool {
	checker := false
	for i := 0; i < len(values)-1; i++ {
		//Check if it is all increasing
		if values[i] > values[i+1] {
			checker = true
			if values[i]-values[i+1] > 3 {
				return false
			}
		} else {
			return false
		}
	}
	return checker
}

func allDecreasing(values []int) bool {
	checker := false
	for i := 0; i < len(values)-1; i++ {
		//Check if it is all decreasing
		if values[i] < values[i+1] {
			checker = true
			if values[i+1]-values[i] > 3 {
				return false
			}
		} else {
			return false
		}
	}
	return checker
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
