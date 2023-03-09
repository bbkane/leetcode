import math
from typing import List

# NOT WORKING
# https://leetcode.com/problems/koko-eating-bananas/


def calc_h(piles, k) -> int:
    return sum(math.ceil(p / k) for p in piles)


def binary_search(min_, max_, ans, f):
    guess = min_ + (min_ + max_) // 2
    guessed_ans = f(guess)
    if guessed_ans > max_:
        return binary_search(min_, guess, ans, f)
    elif guessed_ans < min_:
        return binary_search(guess, max_, ans, f)
    return guess


class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        min_k = len(piles)
        max_k = max(piles)

        def f(guess):
            return calc_h(piles, guess)

        return binary_search(min_k, max_k, h, f)


s = Solution()


piles = [3, 6, 7, 11]
h = 8
k = s.minEatingSpeed(piles, h)
assert k == 4

piles = [30, 11, 23, 4, 20]
h = 5
k = s.minEatingSpeed(piles, h)
assert k == 30

piles = [30, 11, 23, 4, 20]
h = 6
k = s.minEatingSpeed(piles, h)
assert k == 23
