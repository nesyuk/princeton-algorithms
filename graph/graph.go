package graph

import "fmt"

type Node struct {
	Val int
	Neighbours []*Node
}


func (n *Node) String() string {
	if n == nil {
		return ""
	}
	queue := []*Node{n}
	visited := map[*Node]bool{}
	allNodesStr := ""
	for len(queue) > 0 {
		current := queue[0]
		visited[current] = true
		queue = queue[1:]
		neighbours := []int{}
		for _, neighbour := range current.Neighbours {
			neighbours = append(neighbours, neighbour.Val)
			if !visited[neighbour] {
				queue = append(queue, neighbour)
			}
		}	
		nodeStr := fmt.Sprintf("node: %v, neighbours: %v\t", n.Val, neighbours)
		allNodesStr += nodeStr
	}
	return allNodesStr
}

func NewGraph(edges map[int][]int, root int) *Node {
	vertices := map[int]*Node{}
	var rootNode *Node
	for v, neighbours := range edges {
		if _, exist := vertices[v]; !exist {
			vertices[v] = &Node{Val: v}
		}
		if root == v {
			rootNode = vertices[v]
		}
		for _, n := range neighbours {
			if _, exist := vertices[n]; !exist {
				vertices[n] = &Node{Val: n, Neighbours: []*Node{vertices[v]}}
			}
			vertices[v].Neighbours = append(vertices[v].Neighbours, vertices[n])
		}
	}
	return rootNode
}

func NewDiGraph(edges map[int][]int, root int) *Node {
        vertices := map[int]*Node{}
        var rootNode *Node
        for v, neighbours := range edges {
		if _, exist := vertices[v]; !exist {
                        vertices[v] = &Node{Val: v}
                }
                if root == v {
                        rootNode = vertices[v]
                }
                for _, n := range neighbours {
                        if _, exist := vertices[n]; !exist {
                                vertices[n] = &Node{Val: n, Neighbours: []*Node{}}
                        }
                        vertices[v].Neighbours = append(vertices[v].Neighbours, vertices[n])
                }
        }
        return rootNode
}
