from typing import List


# got by 15:22
class Solution:
    def isCovered(self, ranges: List[List[int]], left: int, right: int) -> bool:
        tocover = set(range(left, right + 1))
        numsubrangestocheck = len(ranges)
        while tocover and numsubrangestocheck:
            subrangetocheck = ranges[numsubrangestocheck - 1]
            subrangetocheck_set = set(range(subrangetocheck[0], subrangetocheck[1] + 1))
            # print(f"{subrangetocheck = }")
            tocover = tocover - subrangetocheck_set
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
