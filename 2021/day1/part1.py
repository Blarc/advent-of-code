if __name__ == '__main__':
    with open('input.txt') as f:
        count = 0
        prev = float('inf')
        for line in f:
            current = int(line)
            if current > prev:
                count += 1

            prev = current

    print(count)
