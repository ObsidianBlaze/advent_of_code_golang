package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	data, err := os.ReadFile("day1/day1-part1-input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the content into lines
	lines := strings.Split(string(data), "\n")

	var leftSlice, rightSlice []int

	// Parse each line
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue // Skip empty lines
		}

		// Split each line into two parts
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Invalid line format:", line)
			continue
		}

		// Convert the parts to integers and append to slices
		left, err1 := strconv.Atoi(parts[0])
		right, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil {
			fmt.Println("Error parsing integers in line:", line)
			continue
		}

		leftSlice = append(leftSlice, left)
		rightSlice = append(rightSlice, right)
	}

	leftSorted := sort(leftSlice)
	rightSorted := sort(rightSlice)
	//Get distance by subtracting both inputs
	distance := getDistance(leftSorted, rightSorted)
	//get total distance
	totalDistance := totalDistance(distance)
	fmt.Println("Total distance: ", totalDistance)

	duplicate := countDuplicate(leftSorted, rightSorted)

	repeatSum := sumOfRepeat(duplicate)

	fmt.Println("Sum of repeated numbers: ", repeatSum)

}

// sort the list
func sort(number []int) []int {
	for i := 0; i < len(number)-1; i++ {
		for j := 0; j < len(number)-1-i; j++ {
			if number[j] > number[j+1] {
				number[j], number[j+1] = number[j+1], number[j]
			}
		}
	}
	return number
}

// Get the absolute value
func abs(number int) int {
	if number < 0 {
		return -number
	} else {
		return number
	}
}

// Get the difference of both slices
func getDistance(slice1 []int, slice2 []int) []int {
	var difference []int
	for i := 0; i < len(slice1); i++ {
		difference = append(difference, abs(slice1[i]-slice2[i]))
	}
	return difference
}

// Get the total distance
func totalDistance(distance []int) int {
	total := 0
	for _, v := range distance {
		total += v
	}
	return total
}

// Get duplicate
func countDuplicate(leftSlice []int, rightSlice []int) []int {

	var myRepeatList []int
	for i := 0; i < len(leftSlice); i++ {
		counter := 0
		for j := 0; j < len(rightSlice); j++ {
			if leftSlice[i] == rightSlice[j] {
				counter++
			}
		}
		myRepeatList = append(myRepeatList, leftSlice[i]*counter)
	}

	return myRepeatList
}

func sumOfRepeat(slice []int) int {
	sum := 0
	for i := 0; i < len(slice); i++ {
		sum += slice[i]
	}
	return sum
}
