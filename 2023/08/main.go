package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alperkose/adventofcode/2023/input"
)

func main() {
	network := NewNetwork()
	input.Parse(os.Stdin, network)
	fmt.Println("Pt1: ", network.SolvePt1())
	fmt.Println("Pt2: ", network.SolvePt2_alt())
	// 8245452805243
}

type Network struct {
	instructions string
	nodes        map[string]Node
}

func NewNetwork() *Network {
	return &Network{"", map[string]Node{}}
}
func (gr *Network) Accept(inp string) {
	trimmedInp := strings.TrimSpace(inp)
	if len(trimmedInp) == 0 {
		return
	}
	if strings.Index(trimmedInp, "=") > 0 {
		gr.nodes[trimmedInp[0:3]] = Node{trimmedInp[0:3], trimmedInp[7:10], trimmedInp[12:15]}
	} else {
		gr.instructions = trimmedInp
	}
}

func (gr *Network) SolvePt1() int {
	runningSum := 0
	node := gr.nodes["AAA"]
	for {
		for _, cur := range gr.instructions {
			runningSum++
			if cur == 'R' {
				node = gr.nodes[node.right]
			} else if cur == 'L' {
				node = gr.nodes[node.left]
			}
			// fmt.Printf("go %v -> %v\n", string(cur), node)
			if node.name == "ZZZ" {
				return runningSum
			}
		}
	}
	// return 0
}

func (gr *Network) SolvePt2() int {
	runningSum := 0

	nodes := []Node{}
	for _, v := range gr.nodes {
		if v.name[2] == 'A' {
			nodes = append(nodes, v)
		}
	}
	// for x := 0; x < 4; x++ {
	for {
		for _, cur := range gr.instructions {
			runningSum++
			allEndWithZ := true
			for i := 0; i < len(nodes); i++ {
				if cur == 'R' {
					nodes[i] = gr.nodes[nodes[i].right]
				} else if cur == 'L' {
					nodes[i] = gr.nodes[nodes[i].left]
				}
				allEndWithZ = allEndWithZ && nodes[i].name[2] == 'Z'

			}
			// fmt.Printf("go %v -> %v\n", string(cur), node)
			if allEndWithZ {
				return runningSum
			}
		}
	}
	return runningSum
}

func (gr *Network) SolvePt2_alt() int {
	runningSum := 0

	nodes := []Node{}
	for _, v := range gr.nodes {
		if v.name[2] == 'A' {
			nodes = append(nodes, v)
		}
	}
	steps := make([]int, len(nodes))
	// for x := 0; x < 4; x++ {
	for i := 0; i < len(nodes); i++ {
		runningSum = 0
		for runningSum == 0 {
			for _, cur := range gr.instructions {
				runningSum++
				if cur == 'R' {
					nodes[i] = gr.nodes[nodes[i].right]
				} else if cur == 'L' {
					nodes[i] = gr.nodes[nodes[i].left]
				}
				if nodes[i].name[2] == 'Z' {
					steps[i] = runningSum / len(gr.instructions)
					fmt.Println("sdfsadf", steps)
					break
				}
			}
		}
	}
	runningSum = 1
	for _, v := range steps {
		runningSum *= v
	}
	return runningSum * len(gr.instructions)
}

type Node struct {
	name        string
	left, right string
}
