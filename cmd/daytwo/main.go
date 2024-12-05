package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/CezaryKierzyk/AdventOfCode/pkg/daytwo"
)

func main() {
	path := ""
	flag.StringVar(&path, "i", "", "input file to parse")
	flag.Parse()

	input, err := daytwo.ParseInput(path)
	if err != nil {
		log.Fatalln("Cannot parse input:", err)
	}

	safeCount, err := daytwo.CountSafeReports(input)
	if err != nil {
		log.Fatalln("Failed counting safe reports:", err)
	}
	fmt.Println("Safe reports count without dampener:", safeCount)
	safeCount, err := daytwo.CountSafeReportsWithDampener(input)
}
