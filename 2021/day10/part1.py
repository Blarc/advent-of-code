if __name__ == '__main__':

    points_mapper = {
        ')': 3,
        ']': 57,
        '}': 1197,
        '>': 25137,
    }

    mapper = {
        ')': '(',
        ']': '[',
        '}': '{',
        '>': '<'
    }

    result = 0
    with open('input.txt') as f:

        stack = []
        for line in f:
            line = line.strip()
            for c in line:
                if c in mapper.values():
                    stack.append(c)
                else:
                    popped = stack.pop()
                    if popped != mapper[c]:
                        result += points_mapper[c]

    print(result)
