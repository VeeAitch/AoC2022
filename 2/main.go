package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

var scoreMap map[string]int

func main() {

	// Read File
	// readFile, err := os.Open("input-example")
	readFile, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	result := 0
	createScoreMap()

	for fileScanner.Scan() {
		line := fileScanner.Text()
		result += scoreMap[line]
	}
	fmt.Printf("result: %v", result)
}

func createScoreMap() {

	scoreMap = map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}
}
