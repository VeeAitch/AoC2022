package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"os"
	"regexp"
	"strconv"
)

type valve struct {
	id       int
	name     string
	flowrate int
	connect  []string
}

type frame struct {
	graph map[int]valve
	start int
}

var (
	input frame
)

func main() {

	parseInput("input")
	p1, p2 := compute()
	fmt.Printf("\nDay 16 Part I: %v\n", p1)
	fmt.Printf("\nDay 16 Part Two: %v\n", p2)

}

func compute() (int, int) {
	var graph []Vertex
	for i := range input.graph {
		graph = append(graph, Vertex(i))
	}
	g := ig{graph, make(map[Vertex]map[Vertex]int)}
	for _, node := range input.graph {
		for _, neighbor := range node.connect {
			g.edge(Vertex(node.id), Vertex(getNodeByName(neighbor)), 1)
		}
	}

	// simplify with floyd
	dist, _ := FloydWarshall(g)

	// remove 0er
	var usefullValves []Vertex
	for _, v := range input.graph {
		if v.flowrate > 0 {
			usefullValves = append(usefullValves, Vertex(v.id))
		}
	}

	// do an dfs
	p1 := getDFSPartOne(30, usefullValves, dist, input.start)
	p2 := getDFSPartTwo([]int{26, 26}, usefullValves, dist, []int{input.start, input.start})

	return p1, p2
}

func getDFSPartTwo(timeToGo []int, useful []Vertex, dist map[Vertex]map[Vertex]int, start []int) int {
	best := 0
	var actor int
	if timeToGo[0] > timeToGo[1] {
		actor = 0
	} else {
		actor = 1
	}

	for _, u := range useful {
		newTimeToGo := timeToGo[actor] - dist[Vertex(start[actor])][u] - 1
		if newTimeToGo > 0 {
			timer := []int{newTimeToGo, timeToGo[1-actor]}
			pos := []int{int(u), start[1-actor]}
			gain := newTimeToGo*input.graph[int(u)].flowrate + getDFSPartTwo(timer, removeVertex(useful, u), dist, pos)
			if best < gain {
				best = gain
			}
		}

	}
	return best
}

func getDFSPartOne(timeToGo int, useful []Vertex, dist map[Vertex]map[Vertex]int, start int) int {
	best := 0
	for _, u := range useful {
		newTimeToGo := timeToGo - dist[Vertex(start)][u] - 1
		if newTimeToGo > 0 {
			gain := newTimeToGo*input.graph[int(u)].flowrate + getDFSPartOne(newTimeToGo, removeVertex(useful, u), dist, int(u))
			if best < gain {
				best = gain
			}
		}

	}
	return best
}

func removeVertex(array []Vertex, v Vertex) []Vertex {
	var result []Vertex
	for _, r := range array {
		if r != v {
			result = append(result, r)
		}
	}
	return result
}

func parseInput(filename string) {

	linenumber := 0
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	graph := make(map[int]valve)
	for fileScanner.Scan() {
		linenumber++
		line := fileScanner.Text()
		var v valve
		re := regexp.MustCompile("[0-9]+")
		v.flowrate, _ = strconv.Atoi(re.FindString(line))
		re = regexp.MustCompile("[A-Z][A-Z]")
		caps := re.FindAllString(line, -1)
		v.name = caps[0]
		if len(caps) > 1 {
			for i := 1; i <= len(caps)-1; i++ {
				v.connect = append(v.connect, caps[i])
			}
		}
		v.id = linenumber
		graph[v.id] = v
		if v.name == "AA" {
			input.start = v.id
		}
	}

	readFile.Close()
	input.graph = graph

}

func hash(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32())
}

func getNodeByName(name string) int {
	for index, i := range input.graph {
		if i.name == name {
			return index
		}
	}
	return -1
}
