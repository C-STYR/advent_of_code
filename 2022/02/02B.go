package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func winLoseDraw(elf byte, WLD byte) byte {
	// X = lose, Y = draw, Z = win
	switch WLD {
	case 'Z':
		switch elf {
		case 'A':
			return 'Y'
		case 'B':
			return 'Z'
		case 'C':
			return 'X'
		}
	case 'X':
		switch elf {
		case 'A':
			return 'Z'
		case 'B':
			return 'X'
		case 'C':
			return 'Y'
		}
	case 'Y':
		switch elf {
		case 'A':
			return 'X'
		case 'B':
			return 'Y'
		case 'C':
			return 'Z'
		}
	}
	return 'P'
}

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
		desiredOutcome := fileScanner.Text()[2]
		myPlay := winLoseDraw(elf, desiredOutcome)

		output += logic(elf, myPlay)
	}

	fmt.Println(output)

}
