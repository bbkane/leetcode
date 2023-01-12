# https://leetcode.com/problems/word-break-ii/

from typing import List, NamedTuple


class View(NamedTuple):
    start: int
    end: int


class Node:
    data: View
    children: list[View]


def find_all(haystack: str, needle: str) -> List[View]:
    res = []
    start = 0
    first = haystack.find(needle, start)
    while first != -1:
        res.append(first)
        start += len(needle)
    return res


class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> List[str]:
        # build list of views for words
        # turn that list of views into trees
        # iterate the trees!!
        pass


print(find_all("aba", "a"))
