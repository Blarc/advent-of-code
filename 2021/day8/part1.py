if __name__ == '__main__':

    # 1 - 2
    # 7 - 3
    # 4 - 4
    # 2 - 5
    # 3 - 5
    # 5 - 5
    # 0 - 6
    # 6 - 6
    # 9 - 6
    # 8 - 7

    result = 0
    with open('input.txt') as f:
        for line in f:
            line = line.strip()

            _, output = line.split(' | ')
            words = output.split()
            print(list(filter(lambda x: len(x) in {2, 3, 4, 7}, words)))
            result += len(list(filter(lambda x: len(x) in {2, 3, 4, 7}, words)))

    print(result)
