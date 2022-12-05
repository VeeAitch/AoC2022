package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"unicode"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func main() {

	// Read File
	// readFile, err := os.Open("input-example")
	readFile, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	crates := make(map[int][]rune)
	// Process File line by line
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 && line[0:1] != "m" && line[0:2] != " 1" {
			// Read stacks stating order
			columnCount := ((len(line) - 3) / 4) + 1
			for i := 1; i <= columnCount; i++ {
				position := (i * 4) - 3
				rune := []rune(line[position : position+1])
				if !unicode.IsSpace(rune[0]) {
					crates[i] = append(crates[i], rune[0])
				}
			}

		} else if len(line) == 0 {
			// Reverse order for easier LIFO processing
			for _, crate := range crates {
				ReverseSlice(crate)
			}
			fmt.Printf("Start:\n")
			printCrates(crates)
		} else if len(line) > 0 && line[0:2] != " 1" {
			// Process Moves
			// Get  number of crates
			left := "move "
			right := " from"
			rx := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(left) + `(.*?)` + regexp.QuoteMeta(right))
			matches := rx.FindAllStringSubmatch(line, -1)
			amount, _ := strconv.Atoi(matches[0][1])
			// Get from
			left = " from "
			right = " to "
			rx = regexp.MustCompile(`(?s)` + regexp.QuoteMeta(left) + `(.*?)` + regexp.QuoteMeta(right))
			matches = rx.FindAllStringSubmatch(line, -1)
			from, _ := strconv.Atoi(matches[0][1])
			// Get to
			left = " to "
			rx = regexp.MustCompile(`(?s)` + regexp.QuoteMeta(left) + `(.*?)\z`)
			matches = rx.FindAllStringSubmatch(line, -1)
			to, _ := strconv.Atoi(matches[0][1])
			fmt.Printf("Move %v from %v to %v\n", amount, from, to)

			// Rearange

			indexFrom := len(crates[from]) - amount
			// sub := make([]rune, len(crates[to])+amount)
			sub := crates[from][indexFrom:len(crates[from])]
			crates[to] = append(crates[to], sub...)
			crates[from] = crates[from][:indexFrom]

			printCrates(crates)
		}

	}
	// GetResult
	var result string
	keys := make([]int, 0, len(crates))
	for k := range crates {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		result += string(crates[k][len(crates[k])-1])
	}
	fmt.Printf("result: %v", result)

}

func printCrates(crates map[int][]rune) {
	keys := make([]int, 0, len(crates))
	for k := range crates {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("crate %v: %v \n", k, string(crates[k]))
	}
}

func ReverseSlice[T comparable](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return i > j
	})
}
