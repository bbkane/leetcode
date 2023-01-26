# https://leetcode.com/problems/find-the-town-judge/

from collections import defaultdict
from typing import List

# solved by 1522 (slower than 94% of other solutions lol)


class Solution:
    def findJudge(self, n: int, trust: List[List[int]]) -> int:
        if n == 1 and not trust:
            return -1
        trusted_by = defaultdict(list)
        has_trusts = set()
        for pair in trust:
            trusted_by[pair[1]].append(pair[0])
            has_trusts.add(pair[0])

        for trust_recipient, trust_givers in trusted_by.items():
            if len(trust_givers) == n - 1 and not trust_recipient in has_trusts:
                return trust_recipient
        return -1


s = Solution()

n = 2
trust = [[1, 2]]
assert s.findJudge(n, trust) == 2

n = 3
trust = [[1, 3], [2, 3]]
assert s.findJudge(n, trust) == 3

n = 3
trust = [[1, 3], [2, 3], [3, 1]]
assert s.findJudge(n, trust) == -1
