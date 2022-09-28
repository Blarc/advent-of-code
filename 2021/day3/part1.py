from operator import add

if __name__ == '__main__':
    with open('input.txt') as f:

        result = [0] * 12
        lines = 0
        for line in f:
            result = list(map(add, result, list(map(int, list(line.strip())))))
            lines += 1

    gamma = int(''.join(map(lambda x: '1' if x > lines / 2 else '0', result)), 2)
    epsilon = int(''.join(map(lambda x: '1' if x < lines / 2 else '0', result)), 2)
    print(f'Gamma: {gamma}')
    print(f'Epsilon: {epsilon}')
    print(f'Gamma * Epsilon: {gamma * epsilon}')
