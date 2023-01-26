# https://leetcode.com/problems/cheapest-flights-within-k-stops/

import dataclasses
import itertools
from collections import defaultdict
from heapq import *
from typing import List, NamedTuple


class Flight(NamedTuple):
    to: int
    price: int


@dataclasses.dataclass
class Node:
    flight: Flight
    hops: int
    total_price: int


# https://docs.python.org/3/library/heapq.html


class PriorityQueue:
    def __init__(self):
        self.pq = []  # list of entries arranged in a heap
        self.entry_finder = {}  # mapping of tasks to entries
        self.REMOVED = "<removed-task>"  # placeholder for a removed task
        self.counter = itertools.count()  # unique sequence count

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
        pq.add_task(
            Node(
                flight=Flight(to=src, price=0),
                hops=0,
                total_price=0,
            ),
        )
