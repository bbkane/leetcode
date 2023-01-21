package findifpathexistsingraph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// -- CappedStack

type CappedStack struct {
	stack []int
	index int
}

func NewCappedStack(cap int) CappedStack {
	return CappedStack{
		stack: make([]int, cap),
		index: 0,
	}
}

func (cs *CappedStack) Push(val int) {
	cs.stack[cs.index] = val
	cs.index++
}

func (cs *CappedStack) Pop() int {
	cs.index--
	val := cs.stack[cs.index]
	return val
}

func (cs *CappedStack) Empty() bool {
	return cs.index == 0
}

// -- Set

type Set struct {
	set map[int]struct{}
}

func NewSet(size int) Set {
	return Set{
		set: make(map[int]struct{}, size),
	}
}

func (s *Set) Add(v int) {
	s.set[v] = struct{}{}
}

func (s *Set) Contains(v int) bool {
	_, exists := s.set[v]
	return exists
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}
	// a one-node graph and source doesn't equal dest
	if n == 1 {
		return false
	}
	// convert edgelist to graph (map)
	graph := make(map[int][]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	// make a set of "explored" nodes
	explored := NewSet(n)
	// make a stack/queue of "toExplore" nodes
	// upper bound is number of edges (because len(edges) >= len(nodes))
	// and it's a bidirectional graph so *2
	toExplore := NewCappedStack(len(edges) * 2)
	toExplore.Push(source)
	// for everything in toExplore
	for !toExplore.Empty() {
		current := toExplore.Pop()
		if explored.Contains(current) {
			continue
		}
		explored.Add(current)
		for _, child := range graph[current] {
			if child == destination {
				return true
			}
			toExplore.Push(child)
		}

	}
	return false
}

func Test_validPath(t *testing.T) {
	type args struct {
		n           int
		edges       [][]int
		source      int
		destination int
	}
	tests := []struct {
		name     string
		args     args
		expected bool
	}{
		{
			name: "1",
			args: args{
				n: 3,
				edges: [][]int{
					{0, 1},
					{1, 2},
					{2, 0},
				},
				source:      0,
				destination: 2,
			},
			expected: true,
		},
		{
			name: "2",
			args: args{
				n: 6,
				edges: [][]int{
					{0, 1},
					{0, 2},
					{3, 5},
					{5, 4},
					{4, 3},
				},
				source:      0,
				destination: 5,
			},
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := validPath(tt.args.n, tt.args.edges, tt.args.source, tt.args.destination)
			require.Equal(t, tt.expected, actual)
		})
	}
}
