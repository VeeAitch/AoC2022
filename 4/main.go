package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Printf("\nDay 4 Part One: %v\n", partOne())
	fmt.Printf("\nDay 4 Part Two: %v", partTwo())
}

func partOne() int {
	readFile, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}
	counter := 0
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		pair := fileScanner.Text()
		re := regexp.MustCompile(`\d+`)
		matches := re.FindAllString(pair, -1)
		elveOneStart, _ := strconv.Atoi(matches[0])
		elveOneEnd, _ := strconv.Atoi(matches[1])
		elveTwoStart, _ := strconv.Atoi(matches[2])
		elveTwoEnd, _ := strconv.Atoi(matches[3])

		if (elveOneStart >= elveTwoStart && elveOneEnd <= elveTwoEnd) || (elveOneStart <= elveTwoStart && elveOneEnd >= elveTwoEnd) {
			counter++
		}
	}
	readFile.Close()
	return counter
}

func partTwo() int {
	readFile, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}
	counter := 0
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		pair := fileScanner.Text()
		re := regexp.MustCompile(`\d+`)
		matches := re.FindAllString(pair, -1)
		elveOneStart, _ := strconv.Atoi(matches[0])
		elveOneEnd, _ := strconv.Atoi(matches[1])
		elveTwoStart, _ := strconv.Atoi(matches[2])
		elveTwoEnd, _ := strconv.Atoi(matches[3])

		if ((elveOneStart >= elveTwoStart && elveOneStart <= elveTwoEnd) || (elveOneEnd <= elveTwoEnd && elveOneEnd >= elveTwoStart)) || ((elveOneStart >= elveTwoStart && elveOneEnd <= elveTwoEnd) || (elveOneStart <= elveTwoStart && elveOneEnd >= elveTwoEnd)) {
			counter++
		}
	}
	readFile.Close()
	return counter
}
