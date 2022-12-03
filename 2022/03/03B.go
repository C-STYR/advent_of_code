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
	cache := []string{}

	for sc.Scan() {
		input := sc.Text()
		cache = append(cache, input)

		if len(cache) == 3 {
			CommonLoop:
			for _, c := range cache[0] {
				if strings.Contains(cache[1], string(c)) && strings.Contains(cache[2], string(c)) {
					output += getPriority(c)
					break CommonLoop
				} 
			}
			cache = nil
		}
		
	}

	fmt.Println(output)
}