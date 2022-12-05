package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func main() {

	readFile, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var elves []int
	var sum int
	for fileScanner.Scan() {

		line := fileScanner.Text()
		if line == "" {
			elves = append(elves, sum)
			sum = 0
		} else {
			value, err := strconv.Atoi(line)
			check(err)
			sum += value
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))

	threeFirst := elves[0] + elves[1] + elves[2]

	readFile.Close()
	fmt.Printf("threeFirst: %v", threeFirst)

}
