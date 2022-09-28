if __name__ == '__main__':
    with open('input.txt') as f:
        tmp = list(map(int, f.readline().strip().split(',')))

        # could probably be improved with binary search
        prev = float('inf')
        for i in range(1000):
            cur = sum(map(lambda x, idx=i: (abs(x - idx) * (abs(x - idx) + 1)) / 2, tmp))
            if cur < prev:
                prev = cur
            if cur > prev:
                print(i - 1)
                print(prev)
                break
