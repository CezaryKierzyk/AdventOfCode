package main

import (
	"flag"
	"fmt"

	"github.com/CezaryKierzyk/AdventOfCode2024/pkg/dayone"
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

/*{17633   15737}
79440   47531
44767   73309
86871   26386
66575   90774
31637   38259
89855   50198
81829   20253
95071   98569
34163   43719
28696   84802
10975   89855
48232   15737
14078   99983
40882   73431*/
