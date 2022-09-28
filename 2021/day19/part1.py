import itertools
from collections import Counter
from typing import Tuple


def mult_t(a_: Tuple, b_: Tuple) -> Tuple:
    return tuple(map(lambda x: x[0] * x[1], zip(a_, b_)))


def abs_diff_t(a_: Tuple, b_: Tuple) -> Tuple:
    return tuple(map(lambda x: abs(x[0] - x[1]), zip(a_, b_)))


def diff_t(a_: Tuple, b_: Tuple) -> Tuple:
    return tuple(map(lambda x: x[0] - x[1], zip(a_, b_)))


def roll(v):
    return v[0], v[2], -v[1]


def turn(v):
    return -v[1], v[0], v[2]


def sequence(v):
    for cycle in range(2):
        for step in range(3):  # Yield RTTT 3 times
            v = roll(v)
            yield v  # Yield R
            for _ in range(3):  # Yield TTT
                v = turn(v)
                yield v
        v = roll(turn(roll(v)))  # Do RTR


def rotations(v):
    for _ in range(2):
        for _ in range(3):
            v = list(map(lambda x: roll(x), v))
            yield v
            for _ in range(3):
                v = list(map(lambda x: turn(x), v))
                yield v
        v = list(map(lambda x: roll(turn(roll(x))), v))


if __name__ == '__main__':

    scanners = []

    with open('test.txt') as f:
        f.readline()

        beacons = set()
        for line in f:
            line = line.strip()

            if line.startswith('---'):
                scanners.append(beacons)
                beacons = set()
            elif line == '':
                continue
            else:
                beacons.add(tuple(map(int, line.split(','))))

        scanners.append(beacons)

    # a = set(map(lambda x: abs_diff_t(*x), itertools.combinations(scanners[0], 2)))
    # b = set(map(lambda x: abs_diff_t(*x), itertools.combinations(scanners[1], 2)))
    # 
    # rotations = list(itertools.product([-1, 1], repeat=3))
    # for rotation in rotations:
    #     rotated_b = set(map(lambda x, r=rotation: mult_t(x, r), scanners[1]))
    #     distances_b = set(map(lambda x: abs_diff_t(*x), itertools.combinations(rotated_b, 2)))
    # 
    #     intersection = a.intersection(distances_b)
    #     if len(intersection) > 0:
    #         print(rotation)
    #         print(intersection)
    #         print(len(intersection), len(intersection) / 12)

    # base_dist = set(map(lambda x: abs_diff_t(*x), itertools.combinations(scanners[0], 2)))
    for i_a, scanner_a in enumerate(scanners):
        for i_b, scanner_b in enumerate(scanners):
            # dist_a = Counter(list(map(lambda x: diff_t(*x), filter(lambda x: x[0] != x[1], itertools.product(scanner_a, repeat=2)))))
            dist_a = Counter(list(map(lambda x: diff_t(*x), itertools.product(scanner_a, repeat=2))))
            if scanner_a != scanner_b:
                rotations_ = list(rotations(scanner_b))
                for rotation in rotations_:

                    # dist_b = Counter(list(map(lambda x: diff_t(*x), filter(lambda x: x[0] != x[1], itertools.product(rotation, repeat=2)))))
                    dist_b = Counter(list(map(lambda x: diff_t(*x),itertools.product(rotation, repeat=2))))
                    tmp = dist_a & dist_b
                    print(len(tmp), len(dist_a))
                    # print(f'list: {len(dist_b)}')
                    # dist_b = set(dist_b)
                    # print(f'set: {len(dist_b)}')
                    
                    intersection = [x for x in dist_a if x in dist_b]
                    
                    # if len(dist_a.intersection(dist_b)) > 100:
                    if len(intersection) > 100:
                        print(i_a, i_b)
                        # print(len(dist_a.intersection(dist_b)))
                        print(len(intersection))
                        # print(sorted(scanner_a))
                        # print(sorted(rotation))
                        # print(sorted(dist_a))
                        # print(sorted(dist_b))
                        print('------')
                # for rotation in rotations_:
                #     dist_b = set(map(lambda x: diff_t(*x), itertools.combinations(rotation, 2)))
                # 
                # 
                # 
                #     intersection = dist_a.intersection(dist_b)
                #     if len(intersection) > 0:
                #         print(i_a, i_b)
                #         print(len(intersection))
                #         print(sorted(dist_a))
                #         print(sorted(dist_b))
                
                print('---- new ----')
        break
