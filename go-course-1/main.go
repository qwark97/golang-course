package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func checkErr(err error, msg string) {
	if err != nil {
		fmt.Printf(msg)
		os.Exit(1)
	}
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {

	res := make([]problem, len(lines))

	for i, line := range lines {
		res[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return res
}

func startQuestioning(problems []problem) int {
	finalResult := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s", &answer)
		if answer == p.a {
			finalResult++
		}
	}
	return finalResult
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "A csv file in a format of 'problem,question'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	checkErr(err, fmt.Sprintf("Failed to open csv file: %s\n", *csvFilename))

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	checkErr(err, "Failed to parse provided csv file")

	problems := parseLines(lines)

	finalResult := startQuestioning(problems)

	fmt.Printf("Final result: %d/%d\n", finalResult, len(problems))

}
