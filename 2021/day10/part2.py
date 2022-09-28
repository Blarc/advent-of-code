import math

if __name__ == '__main__':

    points_mapper = {
        ')': 1,
        ']': 2,
        '}': 3,
        '>': 4,
    }

    mapper = {
        ')': '(',
        ']': '[',
        '}': '{',
        '>': '<'
    }

    reverse_mapper = {
        '(': ')',
        '[': ']',
        '{': '}',
        '<': '>'
    }

    results = []
    with open('input.txt') as f:

        incomplete_lines = []

        for line in f:
            stack = []
            line = line.strip()

            broken = False
            for c in line:
                if c in mapper.values():
                    stack.append(c)
                else:
                    popped = stack.pop()
                    if popped != mapper[c]:
                        broken = True
                        break

            line_result = 0
            if not broken and len(stack) > 0:
                while len(stack) > 0:
                    popped = stack.pop()
                    line_result = line_result * 5
                    line_result += points_mapper[reverse_mapper[popped]]

                results.append(line_result)

    print(sorted(results)[math.floor(len(results) / 2)])
