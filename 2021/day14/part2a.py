import functools
from collections import Counter
from collections import deque
from pprint import pprint
from time import time

if __name__ == '__main__':

    memo = {}
    rules = {}
    occurrences = {}
    with open('test.txt') as f:

        template = list(f.readline().strip())
        f.readline()

        for line in f:
            line = list(map(lambda x: x.strip(), line.strip().split('->')))
            rules[line[0]] = line[1]
            if line[1] not in occurrences:
                occurrences[line[1]] = 0

    parts = deque()
    for i in range(len(template) - 1):
        parts.append(''.join(template[i:i + 2]))

    start_time = time()
    # for step in range(2):
    # print(step)
    # print(parts)
    step = 0
    counter = 0
    starter = len(parts)
    
    while len(parts) != 0 and step < 20:
        if counter >= starter:
            starter += 2 * starter
            step += 1
            # print(f'\nstep {step}')

        part = parts.popleft()
        # print(part, end=' ')

        subpart_left = part[0] + rules[part]
        subpart_right = rules[part] + part[1]

        # parts.append(subpart_left)
        # parts.append(subpart_right)

        if subpart_left == part:
            counter += 1
            # print(f'{step} : {subpart_left} : yay left!')
        else:
            parts.append(subpart_left)
            counter += 1

        if subpart_right == part:
            counter += 1
            # print(f'{step} : {subpart_right} : yay right!')
        else:
            parts.append(subpart_right)
            counter += 1

    # parts = new_parts
    # print()

    print(time() - start_time)
