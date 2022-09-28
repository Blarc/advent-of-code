from functools import reduce
from operator import mul


def neighbors(height, width, row_number, column_number):
    tmp = []
    for i in range(row_number - 1, row_number + 2):
        for j in range(column_number - 1, column_number + 2):
            if -1 < i < height and -1 < j < width and \
                    ((i != row_number and j == column_number) or (i == row_number and j != column_number)):
                tmp.append((i, j))
    return tmp


if __name__ == '__main__':
    mapper = {}
    basins = {}
    basin_index = 0

    with open('input.txt') as f:
        for line_index, line in enumerate(f):
            line = list(map(int, list(line.strip())))

            for x_index, x in enumerate(line):

                current = (line_index, x_index)

                if x != 9:
                    if current not in mapper:
                        mapper[current] = basin_index
                        basins[basin_index] = {current}
                        basin_index += 1

                    n = neighbors(line_index + 1, len(line), line_index, x_index)
                    for y in n:
                        if (y[0] <= line_index and y in mapper) or (y[0] == line_index and line[y[1]] != 9):
                            if y in mapper and mapper[y] != mapper[current]:
                                saved_key = mapper[y]
                                for b in list(basins[mapper[y]]):
                                    mapper[b] = mapper[current]
                                    basins[mapper[current]].add(b)

                                basins.pop(saved_key)

                            else:
                                mapper[y] = mapper[current]
                                basins[mapper[current]].add(y)

    print(reduce(mul, (sorted(map(len, basins.values()), reverse=True)[0:3]), 1))
