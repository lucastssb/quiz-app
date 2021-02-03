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
	flag.Parse()

	problems := parseRecords(readCsv(csvFileName))

	showQuestionAndScore(problems)

}

func showQuestionAndScore(problems []problem) {
	score := 0
	for i, p := range problems {
		var answer string
		fmt.Printf("Problem #%d: %s \n", i, p.question)
		fmt.Println("A. ", p.firstAlternative)
		fmt.Println("B. ", p.secondAlternative)
		fmt.Println("C. ", p.thirdAlternative)
		fmt.Println("D. ", p.fourthAlternative)
		fmt.Scanf("%s\n", &answer)

		if strings.ToUpper(answer) == p.answer {
			score++
		}
	}

	fmt.Printf("You scored %d out of %d.\n ", score, len(problems))
}

func timeCounter(done chan int, t int) {
	for i := 1; i <= t; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
	done <- 1
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
	//Gets file from its path
	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	//Parse file
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for: "+*filePath, err)
	}

	return records
}
