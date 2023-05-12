package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// var score int

	totals := []int{}
	alttotals := []int{}
	scores := make(map[string]int)
	altscores := make(map[string]int)

	scores["AX"] = 4
	scores["AY"] = 8
	scores["AZ"] = 3
	scores["BX"] = 1
	scores["BY"] = 5
	scores["BZ"] = 9
	scores["CX"] = 7
	scores["CY"] = 2
	scores["CZ"] = 6

	altscores["AX"] = 3
	altscores["AY"] = 4
	altscores["AZ"] = 8
	altscores["BX"] = 1
	altscores["BY"] = 5
	altscores["BZ"] = 9
	altscores["CX"] = 2
	altscores["CY"] = 6
	altscores["CZ"] = 7

	for scanner.Scan() {

		moves := strings.ReplaceAll(scanner.Text(), " ", "")
		totals = append(totals, scores[moves])
		alttotals = append(alttotals, altscores[moves])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	t := 0
	for _, num := range totals {
		t += num
	}

	at := 0
	for _, num := range alttotals {
		at += num
	}

	fmt.Println(t)
	fmt.Println(at)
}
