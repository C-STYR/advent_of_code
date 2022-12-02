package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scoreMap := map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}

	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	output := 0

	for scanner.Scan() {
		output += scoreMap[scanner.Text()]
	}

	fmt.Println(output)
}
