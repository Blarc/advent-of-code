from heapq import *


def neighbors_resized(row_number, column_number):
    tmp = []
    for i in range(row_number - 1, row_number + 2):
        for j in range(column_number - 1, column_number + 2):
            if -1 < i < height * RESIZE and -1 < j < width * RESIZE and \
                    ((i != row_number and j == column_number) or (i == row_number and j != column_number)):
                tmp.append((i, j))
    return tmp


class Node:
    def __init__(self, pos, risk):
        self.pos = pos
        self.g = float('inf')
        self.h = 0
        self.risk = risk
        self.closed = False
        self.parent: Node = None

    def get_neighbours(self):
        neighbour_nodes = []

        for pos in neighbors_resized(*self.pos):
            if pos not in nodes:
                y, x = pos

                y_index = y // height
                y = y % height

                x_index = x // width
                x = x % width

                risk = matrix[y][x] + y_index + x_index

                if risk > 9:
                    risk = risk % 9

                nodes[pos] = Node(pos, risk)

            neighbour_nodes.append(nodes[pos])

        return neighbour_nodes

    def calc_h(self):
        return abs(self.pos[0] - ((RESIZE * height) - 1)) + abs(self.pos[1] - ((RESIZE * width) - 1))

    def __lt__(self, other):
        return self.g + self.h < other.g + other.h

    def __str__(self):
        # return f'{self.pos[0], self.pos[1]}: {self.risk}, {self.h}'
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

    RESIZE = 5
    matrix = []
    nodes = {}

    with open('input.txt') as f:

        for line in f:
            line = line.strip()
            line = list(map(int, list(line)))
            matrix.append(line)

    height = len(matrix)
    width = len(matrix[0])

    start_node = Node((0, 0), 0)
    start_node.g = 0
    nodes[start_node.pos] = start_node

    queue = []
    heappush(queue, start_node)

    while len(queue) > 0:
        popped: Node = heappop(queue)
        popped.closed = True

        if popped.pos == (RESIZE * height - 1, RESIZE * width - 1):
            print('Path found')
            # print('\n'.join(map(lambda x: x.__str__(), popped.get_path())))
            print(sum(map(lambda x: x.risk, popped.get_path())))
            break
        else:
            neighbours = popped.get_neighbours()
            for s in neighbours:
                if not s.closed and (popped.g + s.risk < s.g):
                    s.parent = popped
                    s.g = popped.g + s.risk
                    s.h = s.calc_h()
                    # If s.h == 0 => Dijkstra
                    heappush(queue, s)
