package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	sc := bufio.NewScanner(input)

	sc.Scan()
	line := sc.Text()
	i := 13
	for {
		cache := map[string]int{}
		for j := i; j >= i-13; j-- {
			cache[string(line[j])] = j
		}
		if len(cache) == 14 {
			fmt.Println(i + 1)
			break
		}
		i++
	}
}
