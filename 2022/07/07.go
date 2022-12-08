package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/google/uuid"
)

type idn = uuid.UUID

type Dir struct {
	name      string
	id        idn
	parent    idn
	children  map[string]idn
	fileSize  int
	totalSize int
}

var Registry = make(map[idn]Dir)
var root = Dir{name: "root", id: uuid.New(), parent: uuid.New(), children: map[string]idn{}, fileSize: 0, totalSize: 0}
var currentDir idn = root.id
var total int = 0

func main() {
	Registry[root.id] = root

	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	sc := bufio.NewScanner(input)
	for sc.Scan() {
		sortEntry(sc.Text())
	}
	getFilesizes(root.id)
	fmt.Println("total:", total) // Answer to part A

	findSmallestDir() // Answer to part B
}

func sortEntry(line string) {
	if line[0] == '$' {
		sortCommand(line)
	}
	if line[0] >= 48 && line[0] <= 57 {
		appendFileSize(line)
	}
	if strings.Contains(line, "dir") {
		addDir(line)
	}
}

func appendFileSize(line string) {
	var number int
	fmt.Sscanf(line, "%d", &number)

	cd := Registry[currentDir]
	cd.fileSize += number
	Registry[currentDir] = cd
}

func addDir(line string) {
	var dir string
	fmt.Sscanf(line, "dir %s", &dir)
	newDir := createDir(dir)
	Registry[currentDir].children[newDir.name] = newDir.id
}

func createDir(dir string) Dir {
	children := map[string]idn{}
	id := uuid.New()
	parent := currentDir

	newDir := Dir{dir, id, parent, children, 0, 0}

	Registry[id] = newDir
	return newDir
}

func sortCommand(line string) {
	var command, dir string
	fmt.Sscanf(line, "$ %s %s", &command, &dir)
	switch command {
	case "cd":
		switch dir {
		case "..":
			currentDir = Registry[currentDir].parent
		case "/":
			currentDir = root.id
		default:
			currentDir = Registry[currentDir].children[dir]
		}
	case "ls":
	}
}

// traverse the tree in postorder (depth first search)
func getFilesizes(root idn) {
	for _, i := range Registry[root].children {
		getFilesizes(i)
	}

	parent := Registry[Registry[root].parent]
	parent.fileSize += Registry[root].fileSize
	Registry[Registry[root].parent] = parent

	if Registry[root].fileSize < 100000 {
		total += Registry[root].fileSize
	}
}

func findSmallestDir() {
	need := Registry[root.id].fileSize - 40000000
	sizes := []int{}

	for _, dir := range Registry {
		if dir.fileSize > need {
			sizes = append(sizes, dir.fileSize)
		}
	}
	sort.Ints(sizes)
	fmt.Println("min directory size:", sizes[0])
}
