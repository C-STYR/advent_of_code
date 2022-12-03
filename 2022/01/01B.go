package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sort"
)

func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	input, err := os.Open("./input/A_input.txt")
	Check(err)
	defer input.Close()

	elves := []int{}
	bag := []int{}

	fileScanner := bufio.NewScanner(input)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		if fileScanner.Text() != "" {
			num, _ := strconv.Atoi(fileScanner.Text())
			bag = append(bag, num)
		} else {
			sum := 0
			for _, calorie := range bag {
				sum += calorie
			}
			elves = append(elves, sum)
			bag = nil
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	baddestElf := elves[0]
	total := baddestElf + elves[1] + elves[2]
	fmt.Println(baddestElf, total)
}
