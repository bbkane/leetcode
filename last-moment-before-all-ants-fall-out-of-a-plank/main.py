from typing import List


class Solution:
    def getLastMoment(self, n: int, left: List[int], right: List[int]) -> int:
        t = 0
        while left or right:
            t += 1
            # step
            for i, ant in enumerate(left):
                left[i] = ant - 1
            for i, ant in enumerate(right):
                right[i] = ant + 1
            # fall off
            left = list(a for a in left if a >= 0)
            right = list(a for a in right if a < n)
            # collision
            for i, ant1 in enumerate(left):
                for j, ant2 in enumerate(right):
                    if ant1 == ant2:
                        left[i], right[j] = right[j], left[i]
        return t

s = Solution()

a = s.getLastMoment(n = 4, left = [4,3], right = [0,1])
assert a == 4