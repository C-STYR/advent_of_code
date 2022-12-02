package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	
	scoreMap := map[string]int{"A X": 3, "A Y": 4, "A Z": 8, "B X": 1, "B Y": 5, "B Z": 9, "C X": 2, "C Y": 6, "C Z": 7}


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
