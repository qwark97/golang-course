package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
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

func startQuiz(problems []problem, quizTimeout int) int {
	finalResult := 0
	timer := time.NewTimer(time.Duration(quizTimeout) * time.Second)
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		var answer string
		answerCh := make(chan string)
		go func() {
			fmt.Scanf("%s", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nTime has finished\n")
			return finalResult
		case answer := <-answerCh:
			if answer == p.a {
				finalResult++
			}
		}
	}
	return finalResult
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "A csv file in a format of 'problem,question'")
	quizTimeout := flag.Int("timeout", 30, "Time for solving problems in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	checkErr(err, fmt.Sprintf("Failed to open csv file: %s\n", *csvFilename))

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	checkErr(err, "Failed to parse provided csv file")

	problems := parseLines(lines)

	finalResult := startQuiz(problems, *quizTimeout)

	fmt.Printf("Final result: %d/%d\n", finalResult, len(problems))

}
