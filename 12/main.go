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

type Point struct {
	x int
	y int
}

type Area struct {
	areamap map[int]map[int]byte
	start   Point
	end     Point
	length  int
	height  int
}

type MazeNode struct {
	point    Point
	parent   Point
	visited  bool
	distance int
}

func main() {
	fmt.Printf("\nDay 12: %v\n", partOne())
}

func partOne() int {

	readFile, err := os.Open("input")
	// readFile, err := os.Open("input-example")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	area := createArea(fileScanner)
	result := computeShortest(area)

	readFile.Close()
	return result
}

// start explorer for a random way, returns 0 if deadend
func computeShortest(area Area) int {

	// Setup maze
	maze := make(map[int]map[int]MazeNode)

	for i := 0; i < area.height; i++ {
		row := make(map[int]MazeNode)
		for j := 0; j < area.length; j++ {
			var mn MazeNode
			mn.point.x = j
			mn.point.y = i
			mn.distance = -1
			mn.visited = false
			row[j] = mn
		}
		maze[i] = row
	}

	// Setup Queue
	q := make([]MazeNode, 0)
	mn := maze[area.start.y][area.start.x]
	mn.distance = 0
	mn.visited = true
	maze[area.start.y][area.start.y] = mn

	q = append(q, mn)

	// Explore
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		// Change S to a for the second part of the puzzle
		if string(area.areamap[node.point.y][node.point.x]) == "S" {
			return node.distance
		}
		// get neighbors
		moves := getBorderavailableMoves(node.point, area)
		moves = checkHeight(moves, area.areamap, node.point)

		for _, m := range moves {
			neighbor := maze[m.y][m.x]
			if !neighbor.visited {
				neighbor.visited = true
				neighbor.parent = node.point
				neighbor.distance = node.distance + 1
				q = append(q, neighbor)
				maze[m.y][m.x] = neighbor
			}

		}

	}

	return 0
}

func createArea(fileScanner *bufio.Scanner) Area {
	linenumber := 0
	// Read in map
	var area Area
	areamap := make(map[int]map[int]byte)

	for fileScanner.Scan() {
		row := make(map[int]byte)
		areamap[linenumber] = row
		line := fileScanner.Text()
		area.length = len(line)
		for i := 0; i < len(line); i++ {
			areamap[linenumber][i] = line[i]
			if string(line[i]) == "E" {
				area.start.y = linenumber
				area.start.x = i
				// fmt.Printf("start: %v\n", area.start)
			} else if string(line[i]) == "S" {
				area.end.y = linenumber
				area.end.x = i
				// fmt.Printf("end: %v\n", area.end)
			}
		}
		linenumber++
	}
	area.areamap = areamap
	area.height = linenumber

	return area
}

func checkHeight(moves []Point, area map[int]map[int]byte, from Point) []Point {
	var result []Point
	origin := area[from.y][from.x]
	if string(origin) == "E" {
		t := []byte("z")
		origin = t[0]
	}
	for _, p := range moves {
		target := area[p.y][p.x]
		if string(target) == "S" {
			t := []byte("a")
			target = t[0]
		}
		if target >= origin-1 {
			result = append(result, p)
		} else {
		}
	}
	return result
}

func getBorderavailableMoves(p Point, area Area) []Point {
	var result []Point
	if checkBorder(p.x, p.y+1, area.length, area.height) {
		move := Point{p.x, p.y + 1}
		result = append(result, move)
	}
	if checkBorder(p.x, p.y-1, area.length, area.height) {
		move := Point{p.x, p.y - 1}
		result = append(result, move)
	}
	if checkBorder(p.x+1, p.y, area.length, area.height) {
		move := Point{p.x + 1, p.y}
		result = append(result, move)
	}
	if checkBorder(p.x-1, p.y, area.length, area.height) {
		move := Point{p.x - 1, p.y}
		result = append(result, move)
	}
	return result
}

func checkBorder(x, y int, length, height int) bool {
	if x+1 > length || x < 0 {
		return false
	} else if y+1 > height || y < 0 {
		return false
	}
	return true
}
