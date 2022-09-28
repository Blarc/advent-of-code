import functools
from collections import Counter
from pprint import pprint
from time import time


def do_rule(rule, iterations):
    for _ in range(0, iterations):
        new_rule = []
        for i in range(0, len(rule) - 1):
            current = rule[i], rule[i + 1]
            new_rule.append(rule[i])
            new_rule.append(rules[current])

        new_rule.append(rule[len(rule) - 1])
        rule = new_rule

    return rule


def do_rec(rule, depth, iterations):
    print(depth, rule)
    if rule in memo:
        saved_rule, saved_depth = memo[rule]
        depth_diff = saved_depth - depth
        if depth < saved_depth:
            return do_rec(saved_rule, iterations - depth_diff, iterations)

    if depth == iterations:
        return rule

    tmp = ''
    for i in range(0, len(rule) - 1):
        current = (rule[i], rule[i + 1])
        next_rule = rule[i] + rules[current] + rule[i + 1]

        res = do_rec(next_rule, depth + 1, iterations)

        if i == len(rule) - 2:
            tmp += res
        else:
            tmp += res[:-1]

    memo[rule] = (tmp, depth)
    return tmp


def do_rec_two(poly, depth, iterations):
    print(f'{depth}:{"----" * depth} {poly}')
    # print(occurrences)
    # if poly in memo:
    #     saved_poly, saved_depth = memo[poly]
    #     depth_diff = saved_depth - depth
    #     if depth < saved_depth:
    #         return do_rec_two(saved_poly, iterations - depth_diff, iterations)

    if depth == iterations:
        new_counter = {}
        for c in empty_counter:
            new_counter[c] = 0

        new_counter[poly[1]] = 1
        print(new_counter)
        return new_counter

    subpoly_left = poly[0] + poly[1]
    subpoly_left = poly[0] + rules[subpoly_left] + poly[1]

    subpoly_right = poly[1] + poly[2]
    subpoly_right = poly[1] + rules[subpoly_right] + poly[2]

    # print(subpoly_left, subpoly_right)

    if poly == subpoly_left and poly == subpoly_right:
        new_counter = {}
        for c in empty_counter:
            new_counter[c] = 0

        new_counter[subpoly_left[1]] += 2 * (iterations - depth)
        print(new_counter)
        return new_counter

    elif poly == subpoly_right:
        left_result = do_rec_two(subpoly_left, depth + 1, iterations)
        left_result[subpoly_right[1]] += (iterations - depth)

        left_result[poly[1]] += 1
        print(left_result)
        return left_result

    elif poly == subpoly_left:
        right_result = do_rec_two(subpoly_right, depth + 1, iterations)
        right_result[subpoly_left[1]] += (iterations - depth)

        right_result[poly[1]] += 1
        print(right_result)
        return right_result

    else:
        left_result = do_rec_two(subpoly_left, depth + 1, iterations)
        right_result = do_rec_two(subpoly_right, depth + 1, iterations)

        for c in left_result:
            right_result[c] += left_result[c]

        right_result[poly[1]] += 1
        print(right_result)
        return right_result


if __name__ == '__main__':

    memo = {}
    rules = {}
    occurrences = {}
    empty_counter = {}

    with open('test.txt') as f:

        template = list(f.readline().strip())
        f.readline()

        for line in f:
            line = list(map(lambda x: x.strip(), line.strip().split('->')))
            rules[line[0]] = line[1]
            if line[1] not in occurrences:
                occurrences[line[1]] = 0
                empty_counter[line[1]] = 0

    # pprint(rules)
    # pprint(occurrences)

    start_time = time()

    for x in template:
        occurrences[x] += 1

    # pprint(occurrences)
    for i in range(len(template) - 1):
        current = ''.join(template[i: i + 2])
        # occurrences[rules[current]] += 1
        current = current[0] + rules[current] + current[1]
        res = do_rec_two(current, 0, 3)
        print(res)
        for xa in res:
            occurrences[xa] += res[xa]

    # do_rec_two('NNC', 0, 2)
    # do_rec_two('')

    # for i in range(len(template) - 1):
    #     part = ''.join(template[i:i + 2])
    #     print(part)
    #     print("--")
    #     do_rec_two(part, 0, 2)
    #     print(occurrences)
    #     print("--")

    pprint(occurrences)
    pprint(Counter(list('NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB')))
    print(time() - start_time)
