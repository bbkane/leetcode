package findifpathexistsingraph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// -- CappedStack

type CappedStack[T any] struct {
	stack []T
	index int
}

func NewCappedStack[T any](cap int) CappedStack[T] {
	return CappedStack[T]{
		stack: make([]T, cap),
		index: 0,
	}
}

func (cs *CappedStack[T]) Push(val T) {
	cs.stack[cs.index] = val
	cs.index++
}

func (cs *CappedStack[T]) Pop() T {
	cs.index--
	val := cs.stack[cs.index]
	return val
}

func (cs *CappedStack[T]) Empty() bool {
	return cs.index == 0
}

// -- Set

type Set[T comparable] struct {
	set map[T]struct{}
}

func NewSet[T comparable](size int) Set[T] {
	return Set[T]{
		set: make(map[T]struct{}, size),
	}
}

func (s *Set[T]) Add(v T) {
	s.set[v] = struct{}{}
}

func (s *Set[T]) Contains(v T) bool {
	_, exists := s.set[v]
	return exists
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	// convert edgelist to graph (map)
	graph := make(map[int][]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	// make a set of "explored" nodes
	explored := NewSet[int](n)
	// make a stack/queue of "toExplore" nodes
	// upper bound is number of edges (because len(edges) >= len(nodes))
	// and it's a bidirectional graph so *2
	toExplore := NewCappedStack[int](len(edges) * 2)
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
