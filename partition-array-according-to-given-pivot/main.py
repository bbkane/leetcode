# https://leetcode.com/problems/partition-array-according-to-given-pivot/

from typing import List


# got by 15:28
class Solution:
    def pivotArray(self, nums: List[int], pivot: int) -> List[int]:
        before_pivot = []
        at_pivot = []
        after_pivot = []
        for num in nums:
            if num < pivot:
                before_pivot.append(num)
            elif num == pivot:
                at_pivot.append(num)
            else:
                after_pivot.append(num)

        return before_pivot + at_pivot + after_pivot


s = Solution()

nums = [9, 12, 5, 10, 14, 3, 10]
pivot = 10
assert s.pivotArray(nums, pivot) == [9, 5, 3, 10, 10, 12, 14]

nums = [-3, 4, 3, 2]
pivot = 2
assert s.pivotArray(nums, pivot) == [-3, 2, 4, 3]
