package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getPriority(c rune) int {
	if c > 96 {
		return int(c) - 96
	} else {
		return int(c) - 38
	}
}

func main() {

	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	sc := bufio.NewScanner(input)

	output := 0

	for sc.Scan() {
		input := sc.Text()
		l := len(input)
		firstCompartment := input[:l/2]
		secondCompartment := input[l/2:]
		
		SubstringLoop:
		for _, s := range secondCompartment {
			if strings.Contains(firstCompartment, string(s)) {
				output += getPriority(s)
				break SubstringLoop
			}
		}
	}

	fmt.Println(output)
}