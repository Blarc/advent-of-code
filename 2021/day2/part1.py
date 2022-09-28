if __name__ == '__main__':
    with open('input.txt') as f:

        horizontal = 0
        depth = 0

        for line in f:
            command, size = line.strip().split()
            size = int(size)
            if command == 'forward':
                horizontal += size
            elif command == 'down':
                depth += size
            else:
                depth -= size

    print(f'Horizontal: {horizontal}')
    print(f'Depth: {depth}')
    print(f'Horizontal * Depth: {horizontal * depth}')
