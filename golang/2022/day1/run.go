package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totals := []int{}
	var cals int

	for scanner.Scan() {

		if scanner.Text() == "" {
			totals = append(totals, cals)
			cals = 0
		} else {
			num, err := strconv.Atoi(scanner.Text())

			if err != nil {
				log.Fatal(err)
			}

			cals += num
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	sort.Ints(totals[:])

	length := len(totals)
	total_top := totals[length-1] + totals[length-2] + totals[length-3]

	fmt.Println("Part One:", totals[length-1])
	fmt.Println("Part Two:", total_top)

}
