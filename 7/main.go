package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

type Dir struct {
	name   string
	parent *Dir
	files  map[string]int
	size   int
}

var treemap []*Dir

func main() {
	fmt.Printf("\nDay 7 Part One: %v\n", partOne())
	fmt.Printf("\nDay 7 Part Two: %v", partTwo())
}

func partTwo() int {
	linenumber := 0
	// readFile, err := os.Open("input-example")
	readFile, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// create directory-tree
	master := new(Dir)
	master.name = "/"
	master.files = make(map[string]int)
	treemap = append(treemap, master)
	activeDir := master
	for fileScanner.Scan() {
		linenumber++
		line := fileScanner.Text()
		// fmt.Printf("%v> line %v: %v\n", activeDir.name, linenumber, line)
		switch string(line[0]) {
		case "$":
			switch string(line[2]) {
			case "c":
				switch line {
				case "$ cd /":
					// fmt.Println("Start at main dir.")
				case "$ cd ..":
					if activeDir.parent != nil {
						// fmt.Printf("go up to %v\n", activeDir.parent.name)
						activeDir = activeDir.parent
					} else {
						// fmt.Println("-------------------------------------------Step out of bounce up.")
					}
				default:
					targetDir := line[5:]
					// fmt.Printf("go to %v\n", targetDir)
					activeDir = getDirectory(targetDir, activeDir)
				}
			case "l":

			}
		case "d":
			// fmt.Printf("create dir %v\n", line[4:])
			createDir(line[4:], activeDir)
		default:
			// fmt.Printf("add file %v\n", line)
			addFile(line, activeDir)
		}
	}

	for _, dir := range treemap {
		dir.size = getDirSize(dir)
	}

	unusedSpace := 70000000 - getDirectory("/", nil).size
	spaceToFree := 30000000 - unusedSpace

	var bigEnoughFolders []*Dir
	for _, dir := range treemap {
		if dir.size >= spaceToFree {
			bigEnoughFolders = append(bigEnoughFolders, dir)
		}
	}
	min := 70000000
	for _, dir := range bigEnoughFolders {
		if dir.size <= min {
			min = dir.size
		}
	}

	readFile.Close()
	return min
}

func partOne() int {
	linenumber := 0
	// readFile, err := os.Open("input-example")
	readFile, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// create directory-tree
	master := new(Dir)
	master.name = "/"
	master.files = make(map[string]int)
	treemap = append(treemap, master)
	activeDir := master
	for fileScanner.Scan() {
		linenumber++
		line := fileScanner.Text()
		// fmt.Printf("%v> line %v: %v\n", activeDir.name, linenumber, line)
		switch string(line[0]) {
		case "$":
			switch string(line[2]) {
			case "c":
				switch line {
				case "$ cd /":
					// fmt.Println("Start at main dir.")
				case "$ cd ..":
					if activeDir.parent != nil {
						// fmt.Printf("go up to %v\n", activeDir.parent.name)
						activeDir = activeDir.parent
					} else {
						fmt.Println("-------------------------------------------Step out of bounce up.")
					}
				default:
					targetDir := line[5:]
					// fmt.Printf("go to %v\n", targetDir)
					activeDir = getDirectory(targetDir, activeDir)
				}
			case "l":

			}
		case "d":
			// fmt.Printf("create dir %v\n", line[4:])
			createDir(line[4:], activeDir)
		default:
			// fmt.Printf("add file %v\n", line)
			addFile(line, activeDir)
		}
	}

	for _, dir := range treemap {
		dir.size = getDirSize(dir)
	}

	var sumofatmostDirs int
	for _, dir := range treemap {
		if dir.size < 100000 {
			sumofatmostDirs += dir.size
		}
	}

	readFile.Close()
	return sumofatmostDirs
}

func getDirSize(dir *Dir) int {
	// size of containing files
	var fileSize int
	for i := range dir.files {
		fileSize += dir.files[i]
	}
	// size of containing dirs
	var dirSize int
	for _, node := range treemap {
		if node.parent == dir {
			dirSize += getDirSize(node)
		}
	}
	return fileSize + dirSize
}

func createDir(name string, parentDir *Dir) {
	dir := new(Dir)
	dir.name = name
	dir.parent = parentDir
	dir.files = make(map[string]int)
	treemap = append(treemap, dir)
}

func addFile(line string, Dir *Dir) {
	fileinfo := strings.Split(line, " ")
	size, _ := strconv.Atoi(fileinfo[0])
	name := fileinfo[1]
	Dir.files[name] = size
}

func getDirectory(name string, parent *Dir) *Dir {

	for _, dir := range treemap {
		if dir.name == name && dir.parent == parent {
			return dir
		}
	}
	return nil
}
