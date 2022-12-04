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

		if checkOverlap(pair) {
			output += 1
		}
	}

	fmt.Println(output)
}

func checkOverlap(slc []string) bool {
	a := strings.Split(slc[0], "-")
	b := strings.Split(slc[1], "-")

	a1, _ := strconv.Atoi(a[0])
	a2, _ := strconv.Atoi(b[0])
	z1, _ := strconv.Atoi(a[1])
	z2, _ := strconv.Atoi(b[1])

	if a1 <= a2 && z1 >= a2 {
		return true
	}
	if a2 <= a1 && z2 >= a1 {
		return true
	}
	return false
}
