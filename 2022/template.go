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

	output := 0

	for sc.Scan() {
		
	}

	fmt.Println(output)
}