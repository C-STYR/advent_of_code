package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
		pair := strings.Split(sc.Text(), ",")

		if converter(pair) {
			output += 1
		}
	}

	fmt.Println(output)
}

func converter(slc []string) bool {
	a := strings.Split(slc[0], "-")
	b := strings.Split(slc[1], "-")

	x1, _ := strconv.Atoi(a[0])
	x2, _ := strconv.Atoi(b[0])
	y1, _ := strconv.Atoi(a[1])
	y2, _ := strconv.Atoi(b[1])

	// compare in both directions
	if x2 <= x1 && y2 >= y1 {
		return true
	}
	if x1 <= x2 && y1 >= y2 {
		return true
	}
	return false
}
