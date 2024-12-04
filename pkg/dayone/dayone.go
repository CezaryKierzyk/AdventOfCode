package dayone

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseInput(path string) (data [2][]int) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a 2D array with 2 rows and an appropriate number of columns
	scanner := bufio.NewScanner(file)

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()

		// Split each line into parts (assuming space-separated values)
		parts := strings.Fields(line)

		// For each part, convert it to an integer and append to the respective row in the 2D array
		for i, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting to integer:", err)
				return
			}
			if i%2 == 0 {
				// First number in each pair goes to the first row
				data[0] = append(data[0], num)
			} else {
				// Second number in each pair goes to the second row
				data[1] = append(data[1], num)
			}
		}
	}
	return
}

func FindSmallest(input []int) (index int) {
	smallest := input[0]
	for i := range input {
		if smallest > input[i] {
			smallest = input[i]
			index = i
		}
	}
	return
}

func RemoveAtIndex(input []int, index int) (out []int) {
	out = append(input[:index], input[index+1:]...)
	return
}

func HasIndex(indices []int, index int) bool {
	for i := range indices {
		if index == indices[i] {
			return true
		}
	}
	return false
}

func RemoveManyIndices(input []int, indices []int) (out []int) {

	for i := range input {
		if !HasIndex(indices, i) {
			out = append(out, input[i])
		}
	}
	return
}

func SortInput(input [2][]int) (sorted [2][]int) {
	// pair up left - right values in ascending order
	var reducedInput [2][]int
	for i := range input {
		reducedInput[i] = make([]int, len(input[i]))
		copy(reducedInput[i], input[i])
	}
	for range reducedInput[0] {
		for j := 0; j < 2; j++ {
			smallestIndex := FindSmallest(reducedInput[j])
			sorted[j] = append(sorted[j], reducedInput[j][smallestIndex])
			reducedInput[j] = RemoveAtIndex(reducedInput[j], smallestIndex)
		}
	}
	return
}

func CalculateDistance(sorted [2][]int) (distance int) {
	// calculate sum of distances between all pairs
	for i := range sorted[0] {
		if sorted[0][i]-sorted[1][i] < 0 {
			distance += (sorted[0][i] - sorted[1][i]) * -1
		} else {
			distance += sorted[0][i] - sorted[1][i]
		}
	}
	return
}

func FindSimilar(value int, input [2][]int) (indices [2][]int) {
	for i := 0; i < 2; i++ {
		for j := range input[i] {
			if input[i][j] == value {
				indices[i] = append(indices[i], j)
			}
		}
	}
	return
}

func CountWithRemove(input [2][]int) (score int) {
	reducedInput := input
	i := 0
	for i < len(reducedInput[0]) {
		value := reducedInput[0][i]
		indices := FindSimilar(value, reducedInput)
		leftCount := len(indices[0])
		rightCount := len(indices[1])
		if leftCount > 0 {
			reducedInput[0] = RemoveManyIndices(reducedInput[0], indices[0])
		}
		if rightCount > 0 {
			reducedInput[1] = RemoveManyIndices(reducedInput[1], indices[1])
		}
		score += leftCount * rightCount * value
	}
	return
}
