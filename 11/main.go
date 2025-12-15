package main

import (
	"aoc25/lib"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	ID          string
	Connections map[string]*Node
}

func main() {
	sourcePtr := flag.String("source", "11/sample.txt", "input file")
	flag.Parse()

	buff, err := os.ReadFile(*sourcePtr)
	if err != nil {
		lib.Fatalf("Error: %v", err)
	}
	input := string(buff)
	input = strings.TrimSpace(input)

	you, svr := parseInput(input)

	var paths int
	if you != nil {
		paths = findAllPaths(you)
		fmt.Printf("Part 1: total number of paths: %d\n", paths)
	} else {
		fmt.Println("Part 1: 'you' node not found")
	}

	if svr != nil {
		paths = solvePart2(svr)
		fmt.Printf("Part 2: total number of valid paths: %d\n", paths)
	} else {
		fmt.Println("Part 2: 'svr' node not found")
	}
}

func parseInput(input string) (*Node, *Node) {
	nodes := make(map[string]*Node)
	var head *Node
	for _, line := range strings.Split(input, "\n") {
		arr := strings.Split(line, ": ")
		id, arr := arr[0], strings.Split(arr[1], " ")
		if nodes[id] == nil {
			nodes[id] = &Node{ID: id}
		}
		if head == nil {
			head = nodes[id]
		}
		for _, conn := range arr {
			if nodes[conn] == nil {
				nodes[conn] = &Node{ID: conn}
			}
			if nodes[id].Connections == nil {
				nodes[id].Connections = make(map[string]*Node)
			}
			nodes[id].Connections[conn] = nodes[conn]
		}
	}
	return nodes["you"], nodes["svr"]
}

func findAllPaths(graph *Node) int {
	paths := 0
	visited := make(map[string]bool)
	visited[graph.ID] = true
	paths += findPaths(graph, visited)
	return paths
}

func findPaths(node *Node, visited map[string]bool) int {
	if node.ID == "out" {
		return 1
	}
	paths := 0
	for _, conn := range node.Connections {
		if visited[conn.ID] {
			continue
		}
		visited[conn.ID] = true
		paths += findPaths(conn, visited)
		visited[conn.ID] = false
	}
	return paths
}

type Note struct {
	Node     string
	dac, fft bool
}

func (n Note) String() string {
	return fmt.Sprintf("Note{Node: %s, dac: %v, fft: %v}", n.Node, n.dac, n.fft)
}

var memo map[string]int

func init() {
	memo = make(map[string]int)
}

func solvePart2(node *Node) int {
	return traverse(node, false, false)
}

func traverse(node *Node, dac, fft bool) int {
	if r, ok := memo[Note{Node: node.ID, dac: dac, fft: fft}.String()]; ok {
		return r
	}
	result := 0
	if node.ID == "out" {
		if dac && fft {
			result = 1
		}
	} else {
		for _, conn := range node.Connections {
			result += traverse(conn, dac || conn.ID == "dac", fft || conn.ID == "fft")
		}
	}
	memo[Note{Node: node.ID, dac: dac, fft: fft}.String()] = result
	return result
}
