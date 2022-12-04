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
		var a1, a2, z1, z2 int
		fmt.Sscanf(sc.Text(), "%d-%d,%d-%d", &a1, &z1, &a2, &z2)

		if a2 <= a1 && z2 >= z1 || a1 <= a2 && z1 >= z2 {
			output += 1
		}
	}

	fmt.Println(output)
}
