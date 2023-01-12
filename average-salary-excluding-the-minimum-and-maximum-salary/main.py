# https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/

import statistics
from typing import List

# started at 15:14 - ended at 15:19


class Solution:
    def average(self, salary: List[int]) -> float:
        return (sum(salary) - min(salary) - max(salary)) / (len(salary) - 2)
