from typing import List

class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children

def traverse(acc, node):
    acc.append(node.val)
    for c in acc.children:
        traverse(acc, c)

class Solution:
    def preorder(self, root: 'Node') -> List[int]:
        acc = []
        traverse(acc, root)
        return acc



# def build_tree(li) -> Node:
#     cur = Node(li[0], [])
#     for e in li[1:]:
#         node = Node(e, [])




s = Solution()

i1 = [1,None,3,2,4,None,5,6]
oE = [1,3,5,6,2,4]
oA = s.preorder(i1)
assert oE == oA

