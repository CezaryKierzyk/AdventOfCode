package main

import (
	"flag"
	"fmt"

	"github.com/CezaryKierzyk/AdventOfCode/pkg/dayone"
)

func main() {
	// parse flags
	var path string

	flag.StringVar(&path, "i", "", "path to input")
	flag.Parse()

	// parse input file
	input := dayone.ParseInput(path)

	//! PART 1 calculate sum of distances between lowest values of two lists

	sorted := dayone.SortInput(input)
	distance := dayone.CalculateDistance(sorted)

	fmt.Println("Distance: " + fmt.Sprint(distance))

	//! PART 2 calculate lists similiarity
	similiarityScore := dayone.CountWithRemove(input)
	fmt.Println("Similiarity Score: " + fmt.Sprintf("%v", similiarityScore))
}
