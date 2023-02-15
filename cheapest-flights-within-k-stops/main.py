# https://leetcode.com/problems/cheapest-flights-within-k-stops/

import dataclasses
import itertools
from collections import defaultdict
from heapq import *
from typing import List, NamedTuple


class Flight(NamedTuple):
    to: int
    price: int


class Node(NamedTuple):
    flight: Flight
    total_hops: int
    total_price: int


# https://docs.python.org/3/library/heapq.html


class PriorityQueue:
    def __init__(self):
        self.pq = []  # list of entries arranged in a heap
        self.entry_finder = {}  # mapping of tasks to entries
        self.REMOVED = "<removed-task>"  # placeholder for a removed task
        self.counter = itertools.count()  # unique sequence count

    def empty(self):
        return len(self.pq) == 0

    def add_task(self, task, priority=0):
        "Add a new task or update the priority of an existing task"
        if task in self.entry_finder:
            self.remove_task(task)
        count = next(self.counter)
        entry = [priority, count, task]
        self.entry_finder[task] = entry
        heappush(self.pq, entry)

    def remove_task(self, task):
        "Mark an existing task as REMOVED.  Raise KeyError if not found."
        entry = self.entry_finder.pop(task)
        entry[-1] = self.REMOVED

    def pop_task(self):
        "Remove and return the lowest priority task. Raise KeyError if empty."
        while self.pq:
            priority, count, task = heappop(self.pq)
            if task is not self.REMOVED:
                del self.entry_finder[task]
                return task
        raise KeyError("pop from an empty priority queue")


class Solution:
    def findCheapestPrice(
        self, n: int, flights: List[List[int]], src: int, dst: int, k: int
    ) -> int:
        # create graph
        graph = defaultdict(set)
        for flight in flights:
            from_, to, price = flight
            graph[from_].add(Flight(to=to, price=price))

        # priority queue!
        pq = PriorityQueue()
        # Add first one
        pq.add_task(
            Node(
                flight=Flight(to=src, price=0),
                total_hops=0,
                total_price=0,
            ),
            priority=0,
        )
        seen = set()
        while not pq.empty():
            current = pq.pop_task()
            seen.add(current.flight.to)
            if current.flight.to == dst:
                return current.total_price

            if current.total_hops > k:
                continue
            for child in graph[current.flight.to]:
                if child in seen:
                    continue
                pq.add_task(
                    Node(
                        flight=child,
                        total_hops=current.total_hops + 1,
                        total_price=current.total_price + child.price,
                    ),
                    priority=current.total_price + child.price,
                )
        return -1


s = Solution()

n = 4
flights = [[0, 1, 100], [1, 2, 100], [2, 0, 100], [1, 3, 600], [2, 3, 200]]
src = 0
dst = 3
k = 1

assert s.findCheapestPrice(n, flights, src, dst, k) == 700


n = 3
flights = [[0, 1, 100], [1, 2, 100], [0, 2, 500]]
src = 0
dst = 2
k = 1
solution = s.findCheapestPrice(n, flights, src, dst, k)
print(solution)
assert solution == 200

n = 3
flights = [[0, 1, 100], [1, 2, 100], [0, 2, 500]]
src = 0
dst = 2
k = 0
assert s.findCheapestPrice(n, flights, src, dst, k) == 500
