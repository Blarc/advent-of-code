if __name__ == '__main__':
    with open('input.txt') as f:
        sequence = list(map(int, f.readline().split(',')))

        boards = []
        for i, line in enumerate(f):

            if i % 6 == 0:
                board = [set() for _ in range(6)]
                boards.append(board)
            else:
                row = list(map(int, line.strip().split()))
                for idx, x in enumerate(row):
                    board[idx + 1].add(x)
                board.append(set(row))
                board[0].update(row)

        i = 4
        num_boards = 0
        bingo = set(sequence[0:5])
        while len(boards) > 0:
            for board in boards:
                for combination in board[1:]:
                    if combination.issubset(bingo):
                        print(combination)
                        print(bingo)
                        print(f'Number called last: {sequence[i]}')
                        print(f'Sum of unmarked numbers: {sum(board[0] - bingo)}')
                        print(f'Result: {sum(board[0] - bingo) * sequence[i]}')
                        num_boards += 1

                        boards.remove(board)
                        break

            i += 1
            bingo.add(sequence[i])
