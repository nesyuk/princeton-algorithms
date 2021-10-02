package graph

func Bfs(node *Node, target int) bool {
	if node == nil {
		return false
	}
	visited := map[*Node]bool{}
	queue := []*Node{node}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		visited[current] = true
		if current.Val == target {
			return true
		}
		for _, neighbour := range current.Neighbours {
			if !visited[neighbour] {
				queue = append(queue, neighbour)
			}
		}
	}

	return false
}
