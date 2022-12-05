package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Stack []string

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(str string) {
	*s = append(*s, str)
} 

func (s *Stack) pop() (string) {
	if s.isEmpty() {
		log.Fatalf("Stack %v Empty", s)
	}
	i := len(*s) -1
	str := (*s)[i]
	*s = (*s)[:i]
	return str
}

func main() {

	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	sc := bufio.NewScanner(input)

	stack1 := Stack{"T", "P", "Z", "C", "S", "L", "Q", "N"}
	stack2 := Stack{"L", "P", "T", "V", "H", "C", "G"}
	stack3 := Stack{"D", "C", "Z", "F"}
	stack4 := Stack{"G", "W", "T", "D", "L", "M", "V", "C"}
	stack5 := Stack{"P", "W", "C"}
	stack6 := Stack{"P", "F", "J", "D", "C", "T", "S", "Z"}
	stack7 := Stack{"V", "W", "G", "B", "D"}
	stack8 := Stack{"N", "J", "S", "Q", "H", "W"}
	stack9 := Stack{"R", "C", "Q", "F", "S", "L", "V"}

	stackMap := map[int]*Stack{1: &stack1, 2: &stack2, 3: &stack3, 4: &stack4, 5: &stack5, 6: &stack6, 7: &stack7, 8: &stack8, 9: &stack9}

	var n, origin, destination int

	for sc.Scan() {
		fmt.Sscanf(sc.Text(), "move %d from %d to %d", &n, &origin, &destination)

		tempStack := Stack{}

		for i := n; i > 0; i-- {
			tempStack.push(stackMap[origin].pop())
		}
		for i := n; i > 0; i-- {
			stackMap[destination].push(tempStack.pop())
		}
	}

	fmt.Println("output", getOutput([]Stack{stack1, stack2, stack3, stack4, stack5, stack6, stack7, stack8, stack9,}))
}

func getOutput(s []Stack) string {
	var output string
	for i := range s {
		output += s[i][len(s[i]) -1]
	}
	return output
}