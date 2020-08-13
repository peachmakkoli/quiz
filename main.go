package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// creates a help file
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	// open the csv file
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	// read the csv file
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(lines)

	// TODO: break this section out into a helper function
	// keep track of score
	correct := 0
	// print out each problem
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		// prompt user for answer
		var answer string
		fmt.Scanf("%s\n", &answer)
		// check answer, increment score if correct
		if answer == p.a {
			correct++
		}
	}

	// print final score
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

// helper function that converts csv data into problem structs
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]), // remove spaces from csv data
		}
	}
	return ret
}

// define struct for problem data
type problem struct {
	q string
	a string
}

// helper function, prints an error message and exits the program
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
