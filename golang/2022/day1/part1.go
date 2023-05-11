package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var largest int = 0
	var total int

	for scanner.Scan() {

		if scanner.Text() == "" {
			if total > largest {
				largest = total
			}
			total = 0
		} else {
			num, err := strconv.Atoi(scanner.Text())

			if err != nil {
				log.Fatal(err)
			}

			total += num
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(largest)

}
