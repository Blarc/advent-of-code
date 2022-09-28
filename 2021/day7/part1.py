from statistics import *

if __name__ == '__main__':
    with open('input.txt') as f:
        tmp = list(map(int, f.readline().strip().split(',')))

        med = median(tmp)
        tmp = list(map(lambda x: abs(x - med), tmp))
        print(sum(tmp))
