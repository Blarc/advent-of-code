def neighbors(matrix_, row_number, column_number):
    tmp = []
    for i in range(row_number - 1, row_number + 2):
        for j in range(column_number - 1, column_number + 2):
            if -1 < i < len(matrix_) and -1 < j < len(matrix_[0]) and \
                    ((i != row_number and j == column_number) or (i == row_number and j != column_number)):
                tmp.append((i, j))
    return tmp


if __name__ == '__main__':
    matrix = []
    with open('input.txt') as f:
        for line in f:
            line = list(map(int, list(line.strip())))
            matrix.append(line)

    result = 0
    for y in range(len(matrix)):
        for x in range(len(matrix[0])):
            n = neighbors(matrix, y, x)
            l = list(filter(lambda a, x_=x, y_=y: matrix[a[0]][a[1]] > matrix[y_][x_], n))
            if len(l) == len(n):
                risk = (matrix[y][x] + 1)
                result += risk

    print(result)
