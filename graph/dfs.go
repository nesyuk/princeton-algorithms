package graph


func Dfs(node *Node, target int) bool {
	if node == nil {
		return false
	}
	visited := map[*Node]bool{}
	stack := []*Node{node}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		visited[current] = true
		if current.Val == target {
			return true
		}
		for _, neighbour := range current.Neighbours {
			if !visited[neighbour] {
				stack = append(stack, neighbour)
			}
		}
	}
	
	return false
}

