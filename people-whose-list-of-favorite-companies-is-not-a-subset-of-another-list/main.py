# https://leetcode.com/problems/people-whose-list-of-favorite-companies-is-not-a-subset-of-another-list/

from typing import List

class Solution:
    def peopleIndexesold(self, favoriteCompanies: List[List[str]]) -> List[int]:
        is_subset = []
        for p1 in favoriteCompanies:
            sp1 = set(p1)
            for p2 in favoriteCompanies:
                sp2 = set(p2)
                if sp1 != sp2:
                    if sp1.issubset(p2):
                        print(f"is subset: {sp1}, {sp2}")
                        is_subset.append(True)
                        continue
            is_subset.append(False)
        return [i for (i, e) in enumerate(is_subset) if e == True]

    def peopleIndexes(self, favoriteCompanies: List[List[str]]) -> List[int]:
        # map is_not_subset: groupid -> group id
        is_subset = dict()
        for i, p1 in enumerate(favoriteCompanies):
            for j, p2 in enumerate(favoriteCompanies):
                if i == j:
                    continue
                if set(p1).issubset(set(p2)):
                    is_subset[i] = j
        return sorted(list(set(range(len(favoriteCompanies))) - is_subset.keys()))






s = Solution()

print("-- t1")
t1 = [["leetcode","google","facebook"],["google","microsoft"],["google","facebook"],["google"],["amazon"]]
print(s.peopleIndexes(t1))  # [0,1,4]

print("-- t2")
t2 = [["leetcode","google","facebook"],["leetcode","amazon"],["facebook","google"]]
print(s.peopleIndexes(t2)) # 0, 1

print("-- t3")
t3 = [["leetcode"],["google"],["facebook"],["amazon"]]
print(s.peopleIndexes(t3)) # 0, 1, 2, 3