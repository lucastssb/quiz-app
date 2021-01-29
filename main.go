package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	problems := readCsv("problems.csv")
	score := 0
	for k, v := range problems {
		var answer string
		fmt.Printf("Problem #%d: %s = ", k, v[0])
		fmt.Scan(&answer)

		if answer == v[1] {
			score++
		}
	}

	fmt.Println("Score: ", score)

}

func readCsv(filePath string) [][]string {
	//Gets file from its path
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//Parse file
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for: "+filePath, err)
	}

	return records
}
