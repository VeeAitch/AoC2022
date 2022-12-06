package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var prio (map[string]int)

func main() {

	setPrio()
	fmt.Println(prio)

	fmt.Printf("\nDay 3 Part One: %v", partOne())
	fmt.Printf("\nDay 3 Part Two: %v", partTwo())
}

func partTwo() int {
	readFile, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	prioSum := 0
	rowcount := 1
	elvecount := 0
	group := make([]string, 3)

	for fileScanner.Scan() {
		rucksack := fileScanner.Text()
		group[elvecount] = rucksack
		if (rowcount % 3) == 0 {
			for _, char := range group[0] {
				if strings.Contains(group[1], string(char)) {
					if strings.Contains(group[2], string(char)) {
						prioSum += prio[string(char)]
						// fmt.Printf("share item type %v with prio %v \n", string(char), prio[string(char)])
						break
					}
				}
			}
			elvecount = -1
		}
		rowcount++
		elvecount++
	}
	readFile.Close()
	return prioSum
}

func partOne() int {
	readFile, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var prioSum int = 0
	for fileScanner.Scan() {
		rucksack := fileScanner.Text()
		// fmt.Println(rucksack)
		firstCompartment := rucksack[:(len(rucksack) / 2)]
		secondCompartment := rucksack[(len(rucksack) / 2):len(rucksack)]
		// fmt.Println(firstCompartment)
		// fmt.Println(secondCompartment)

		for _, char := range firstCompartment {
			if strings.Contains(secondCompartment, string(char)) {
				prioSum += prio[string(char)]
				// fmt.Printf("share item type %v with prio %v \n", string(char), prio[string(char)])
				break
			}
		}

	}
	readFile.Close()
	return prioSum
}

func setPrio() {
	prio = make(map[string]int)
	for a := 1; a < 27; a++ {
		prio[toCharLowercase(a)] = a
	}
	for A := 1; A < 27; A++ {
		prio[toCharUppercase(A)] = A + 26
	}

}

func toCharLowercase(i int) string {
	return string('a' - 1 + i)
}

func toCharUppercase(i int) string {
	return string('A' - 1 + i)
}
