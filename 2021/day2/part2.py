if __name__ == '__main__':
    with open('input.txt') as f:

        horizontal = 0
        depth = 0
        aim = 0

        for line in f:
            command, size = line.strip().split()
            size = int(size)
            if command == 'forward':
                horizontal += size
                depth += aim * size
            elif command == 'down':
                aim += size
            elif command == 'up':
                aim -= size
            else:
                print('Unknown command!')

    print(f'Horizontal: {horizontal}')
    print(f'Depth: {depth}')
    print(f'Aim: {aim}')
    print(f'Horizontal * Depth: {horizontal * depth}')
