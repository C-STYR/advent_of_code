package main

import (
	"bufio"
	"fmt"

	"log"
	"os"
)

func logic(elf, me byte) int {
	score := 0
	if elf == 'A' {
		if me == 'X' {
			score += 4
		}
		if me == 'Y' {
			score += 8
		}
		if me == 'Z' {
			score += 3
		}
	}
	if elf == 'B' {
		if me == 'X' {
			score += 1
		}
		if me == 'Y' {
			score += 5
		}
		if me == 'Z' {
			score += 9
		}
	}
	if elf == 'C' {
		if me == 'X' {
			score += 7
		}
		if me == 'Y' {
			score += 2
		}
		if me == 'Z' {
			score += 6
		}
	}
	return score
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	fileScanner := bufio.NewScanner(input)

	fileScanner.Split(bufio.ScanLines)

	output := 0

	for fileScanner.Scan() {
		elf := fileScanner.Text()[0]
		me := fileScanner.Text()[2]
		output += logic(elf, me)
	}

	fmt.Println(output)
}
