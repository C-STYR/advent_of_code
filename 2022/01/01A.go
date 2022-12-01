package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"os"
)

type Elf struct {
	bag []int
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	input, err := os.Open("./input/A_input.txt")
	check(err)
	defer input.Close()

	elves := []int{}
	bag := []int{}
	
	fileScanner := bufio.NewScanner(input)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		// if line is not blank text 
		if fileScanner.Text() != "" {
			// convert to num and push into slice
			num, _ := strconv.Atoi(fileScanner.Text())
			bag = append(bag, num)
			// fmt.Println(bag)
		} else {
			// if it is blank text
			sum := 0
			// sum slice and push into averages
			for _, calorie := range bag {
				sum += calorie
			}
			elves = append(elves, sum)
			// empty slice
			bag = nil
		}
	}

	baddestElf := 0
	for _, calories := range elves {
		if calories > baddestElf {
			baddestElf = calories
		}
	}
	fmt.Println(baddestElf) //69883
}