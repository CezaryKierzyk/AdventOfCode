package daytwo

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ParseInput(path string) (input [][]int, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	l := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		input = append(input, make([]int, len(parts)))
		for p := range parts {
			value, e := strconv.Atoi(parts[p])
			if err != nil {
				return [][]int{}, e
			}
			input[l][p] = value
		}
		l++
	}
	return
}

// Tendency type is true for ascending false for descending
func CheckTendency(input []int, dampen bool) (hasTendency bool) {
	hasTendency = false
	isAscending := false
	for i := 1; i < len(input); i++ {
		if i == 1 {
			if input[i] > input[i-1] {
				hasTendency = true
				isAscending = true
				continue
			}
			if input[i] < input[i-1] {
				if dampen {
					dampen = !dampen
				}
				hasTendency = true
				isAscending = false
				continue
			}
			if input[i] == input[i-1] {
				if dampen {
					dampen = !dampen
				}
				hasTendency = false
				isAscending = false
				return
			}
		}
		if i != 1 && hasTendency {
			if input[i] > input[i-1] && isAscending {
				continue
			}
			if input[i] < input[i-1] && !isAscending {
				continue
			}
			if input[i] > input[i-1] && !isAscending {
				if dampen {
					dampen = !dampen
				}
				hasTendency = false
				return
			}
			if input[i] < input[i-1] && isAscending {
				if dampen {
					dampen = !dampen
				}
				hasTendency = false
				return
			}
			if input[i] == input[i-1] {
				if dampen {
					dampen = !dampen
				}
				hasTendency = false
				return
			}
		}
	}
	return
}

func IsSafe(input []int) (safe bool) {
	safe = false
	for i := 1; i < len(input); i++ {
		diff := input[i] - input[i-1]
		if diff < -3 || diff > 3 {
			return
		}
	}
	return true
}

func CountSafeReports(input [][]int) (count int, err error) {
	for i := range input {
		hasTendency := CheckTendency(input[i], false)
		if !hasTendency {
			continue
		}
		if IsSafe(input[i]) {
			count++
		}
	}
	return
}

func CountSafeReportsWithDampener(input [][]int) (count int, err error) {
	for i := range input {
		hasTendency := CheckTendency(input[i], true)
		if !hasTendency {
			continue
		}
		if IsSafe(input[i]) {
			count++
		}
	}
	return
}
