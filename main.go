package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 3, "the time limit for the quiz in seconds")
	flag.Parse()

	problems := parseRecords(readCsv(csvFileName))

	showQuestionAndScore(problems, timeLimit)

}

func showQuestionAndScore(problems []problem, timeLimit *int) {
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	score := 0

questionsLoop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s \n", i, p.question)
		fmt.Println("A. ", p.firstAlternative)
		fmt.Println("B. ", p.secondAlternative)
		fmt.Println("C. ", p.thirdAlternative)
		fmt.Println("D. ", p.fourthAlternative)
		answerChan := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break questionsLoop
		case answer := <-answerChan:
			if strings.ToUpper(answer) == p.answer {
				score++
			}
		}
	}

	fmt.Printf("You scored %d out of %d. ", score, len(problems))
}

type problem struct {
	question          string
	firstAlternative  string
	secondAlternative string
	thirdAlternative  string
	fourthAlternative string
	answer            string
}

func parseRecords(records [][]string) []problem {
	ret := make([]problem, len(records))
	for i, record := range records {
		ret[i] = problem{
			question:          record[0],
			firstAlternative:  record[1],
			secondAlternative: record[2],
			thirdAlternative:  record[3],
			fourthAlternative: record[4],
			answer:            record[5],
		}
	}

	return ret
}

func readCsv(filePath *string) [][]string {
	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for: "+*filePath, err)
	}

	return records
}
