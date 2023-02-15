# https://leetcode.com/problems/minimum-number-of-vertices-to-reach-all-nodes/
from typing import List

# The nodes with no incoming edges MUST BE part of the solution set.
# All other nodes are reachable from one of these source nodes because they have incoming edges


class Solution:
    def findSmallestSetOfVertices(self, n: int, edges: List[List[int]]) -> List[int]:
        remaining = {i for i in range(n)}
        for edge in edges:
            incoming = edge[1]
            if incoming in remaining:
                remaining.remove(edge[1])
        ret = sorted(remaining)
        return ret


s = Solution()

n = 6
edges = [[0, 1], [0, 2], [2, 5], [3, 4], [4, 2]]
expected = [0, 3]
actual = s.findSmallestSetOfVertices(n, edges)
assert expected == actual

n = 5
edges = [[0, 1], [2, 1], [3, 1], [1, 4], [2, 4]]
expected = [0, 2, 3]
actual = s.findSmallestSetOfVertices(n, edges)
assert expected == actual
