from pprint import pprint


def neighbors(row_number, column_number, matrix_):
    tmp = []
    for i in range(row_number - 1, row_number + 2):
        for j in range(column_number - 1, column_number + 2):
            if -1 < i < len(matrix_) and -1 < j < len(matrix_[0]) and (i != row_number or j != column_number):
                tmp.append((i, j))
    return tmp


def flash_bright(j, i):
    result = 0
    for n in neighbors(j, i, matrix):
        n_j = n[0]
        n_i = n[1]
        if matrix[n_j][n_i] != 0:
            matrix[n_j][n_i] += 1
            if matrix[n_j][n_i] > 9:
                matrix[n_j][n_i] = 0
                result += flash_bright(n_j, n_i) + 1

    return result


if __name__ == '__main__':

    matrix = []
    with open('input.txt') as f:
        for line in f:
            line = line.strip()
            matrix.append(list(map(int, list(line))))

    for step in range(1, 500):
        print(f'Step: {step}')
        flashes = []
        for y in range(len(matrix)):
            for x in range(len(matrix[0])):
                matrix[y][x] += 1
                if matrix[y][x] > 9:
                    matrix[y][x] = 0
                    flashes.append((y, x))

        number_of_flashes = 0
        for flash in flashes:
            number_of_flashes += flash_bright(flash[0], flash[1]) + 1

        if number_of_flashes == 100:
            break

        # pprint(matrix)
