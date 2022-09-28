from pprint import pprint
from heapq import *
from typing import List


def neighbors(row_number, column_number):
    tmp = []
    for i in range(row_number - 1, row_number + 2):
        for j in range(column_number - 1, column_number + 2):
            if -1 < i < len(matrix) and -1 < j < len(matrix[0]) and \
                    ((i != row_number and j == column_number) or (i == row_number and j != column_number)):
                tmp.append((i, j))
    return tmp


class Node:
    def __init__(self, pos, risk):
        self.pos = pos
        self.g = 0
        self.h = 0
        self.f = 0
        self.risk = risk
        self.parent: Node = None
        self.closed = False

    def calc_h(self):
        return abs(self.pos[0] - height) + abs(self.pos[1] - width)

    def __lt__(self, other):
        return self.f < other.f

    def __str__(self):
        return f'{self.pos[0], self.pos[1]}: {self.risk}'

    def get_path(self):
        path = []
        current = self

        while current.parent is not None:
            path.insert(0, current)
            current = current.parent

        path.insert(0, current)
        return path


if __name__ == '__main__':

    matrix = []

    with open('input.txt') as f:

        for j, line in enumerate(f):
            line = line.strip()
            line = list(map(int, list(line)))
            matrix.append(list(map(lambda x, j_=j: Node((j_, x[0]), x[1]), enumerate(line))))

    height = len(matrix)
    width = len(matrix[0])

    queue = []
    matrix[0][0].risk = 0
    heappush(queue, matrix[0][0])

    while len(queue) > 0:
        popped: Node = heappop(queue)
        popped.closed = True

        if popped.pos == (height - 1, width - 1):
            print('Path found')
            # print('\n'.join(map(lambda x: x.__str__(), popped.get_path())))
            print(sum(map(lambda x: x.risk, popped.get_path())))
            exit(0)
        else:
            n = list(map(lambda x: matrix[x[0]][x[1]], neighbors(*popped.pos)))
            # print(list(map(lambda x: x.__str__(), n)))
            for s in n:
                if not s.closed and (s.parent is None or s.parent.g > s.g):
                    s.parent = popped
                    s.g = popped.g + s.risk
                    s.h = s.calc_h()
                    s.f = s.g + s.h
                    heappush(queue, s)
