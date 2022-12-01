#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from typing import List

# https://leetcode.com/problems/intersection-of-multiple-arrays/submissions/853013436/


class Solution:
    def intersection(self, nums: List[List[int]]) -> List[int]:
        start = set(nums[0])
        for li in nums:
            start = start.intersection(set(li))
        ret = list(start)
        ret.sort()
        return ret


s = Solution()
t1 = [[3, 1, 2, 4, 5], [1, 2, 3, 4], [3, 4, 5, 6]]
print(s.intersection(t1))

t2 = [[1, 2, 3], [4, 5, 6]]
print(s.intersection(t2))
