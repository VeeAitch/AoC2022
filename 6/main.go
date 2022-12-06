package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func main() {

	// Read File
	data, err := os.ReadFile("input")
	check(err)
	var result int

	for i := 0; i < len(data)-3; i++ {
		var one, two, three, four string
		one = string(data[i : i+1])
		two = string(data[i+1 : i+2])
		three = string(data[i+2 : i+3])
		four = string(data[i+3 : i+4])
		fmt.Printf("%v %v %v %v\n", one, two, three, four)

		checkerMap := make(map[string]bool)

		checkerMap[one] = true
		if _, ok := checkerMap[two]; ok {
			// do nothing :=)
		} else {
			checkerMap[two] = true
			if _, ok := checkerMap[three]; ok {
				// do nothing :=)
			} else {
				checkerMap[three] = true
				if _, ok := checkerMap[four]; ok {
					// do nothing :=)
				} else {
					checkerMap[four] = true
					result = i + 4
					break
				}
			}
		}

	}
	fmt.Printf("result: %v\n", result)
}
