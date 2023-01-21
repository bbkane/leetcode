from typing import List


def validPathHelper(
    frontier, n: int, edges: List[List[int]], source: int, destination: int
) -> bool:
    if source == destination:
        return True
    if not frontier:
        return False
    for e in frontier:
        if e[0] == source:
            if e[1] == destination:
                return True
            else:
                newfrontier = frozenset(i for i in frontier if i != e)
                return validPathHelper(newfrontier, n, edges, e[1], destination)
    return False


class Solution:
    def validPath(self, n: int, edges: List[List[int]], source: int, destination: int) -> bool:
        frontier = frozenset(tuple(e) for e in edges)
        frontier_back = frozenset((e[1], e[0]) for e in edges)
        frontier = frozenset([*frontier, *frontier_back])
        print(frontier)
        return validPathHelper(frontier, n, edges, source, destination)


s = Solution()

n = 3
edges = [[0, 1], [1, 2], [2, 0]]
source = 0
destination = 2

print(s.validPath(n, edges, source, destination))


n = 6
edges = [[0, 1], [0, 2], [3, 5], [5, 4], [4, 3]]
source = 0
destination = 5
print(s.validPath(n, edges, source, destination))


n = 1
edges = [[0, 4]]
source = 0
destination = 4
print(s.validPath(n, edges, source, destination))

n = 1
edges = [[0, 4]]
source = 0
destination = 4
print(s.validPath(n, edges, source, destination))
