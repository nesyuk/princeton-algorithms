package graph

import (
	"testing"
)


func TestDfs(t *testing.T) {
	root := NewDiGraph(map[int][]int{1: []int{2, 3}, 2: []int{3}, 4: []int{4}}, 1)
	tests := []struct{data int; want bool}{
		{4, false},
		{3, true},
		{2, true},
	}
	for _, test := range tests {
		isConnected := Dfs(root, test.data)
		if isConnected != test.want {
			t.Fatalf("root: %v to node: %v, expected: %v, got: %v", root.Val, test.data, test.want, isConnected)
		}
	}
}

func TestDfsWithLoops(t *testing.T) {
	root := NewDiGraph(map[int][]int{1: []int{2, 3}, 2: []int{1, 3}, 4: []int{4}}, 1)
	tests := []struct{data int; want bool}{
		{4, false},
		{3, true},
		{2, true},
	}
	for _, test := range tests {
		isConnected := Dfs(root, test.data)
		if isConnected != test.want {
			t.Fatalf("root: %v to node: %v, expected: %v, got: %v", root.Val, test.data, test.want, isConnected)
		}
	}
}

func TestBfs(t *testing.T) {
	root := NewDiGraph(map[int][]int{1: []int{2, 3}, 2: []int{3}, 4: []int{4}}, 1)
	tests := []struct{data int; want bool}{
		{4, false},
		{3, true},
		{2, true},
	}
	for _, test := range tests {
		isConnected := Bfs(root, test.data)
		if isConnected != test.want {
			t.Fatalf("root: %v to node: %v, expected: %v, got: %v", root.Val, test.data, test.want, isConnected)
		}
	}
}

func TestBfsWithLoops(t *testing.T) {
	root := NewDiGraph(map[int][]int{1: []int{2, 3}, 2: []int{1, 3}, 4: []int{4}}, 1)
	tests := []struct{data int; want bool}{
		{4, false},
		{3, true},
		{2, true},
	}
	for _, test := range tests {
		isConnected := Bfs(root, test.data)
		if isConnected != test.want {
			t.Fatalf("root: %v to node: %v, expected: %v, got: %v", root.Val, test.data, test.want, isConnected)
		}
	}
}
