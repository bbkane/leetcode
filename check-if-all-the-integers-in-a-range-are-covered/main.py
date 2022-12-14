from typing import List


def subrange_contains(subrange: List[int], n: int) -> bool:
    return subrange[0] <= n <= subrange[1]


# got by 15:22
class Solution:
    def isCovered(self, ranges: List[List[int]], left: int, right: int) -> bool:
        tocover = set(range(left, right + 1))
        numsubrangestocheck = len(ranges)
        while tocover and numsubrangestocheck:
            subrangetocheck = ranges[numsubrangestocheck - 1]
            tocover = set(i for i in tocover if not subrange_contains(subrangetocheck, i))
            numsubrangestocheck -= 1
        # print(tocover)
        if not tocover:
            return True
        return False


s = Solution()

ranges = [[1, 2], [3, 4], [5, 6]]
left = 2
right = 5

print(s.isCovered(ranges, left, right))

ranges = [[1, 10], [10, 20]]
left = 21
right = 21
print(s.isCovered(ranges, left, right))

ranges = [[1, 50]]
left = 1
right = 50
print(s.isCovered(ranges, left, right))
