def find_paths(node, path, small_twice):
    if node == 'end':
        paths.append(path)
        return

    for n in mapper[node]:
        if n != 'start':
            if n.isupper() or n not in path or not small_twice:
                path.append(node)
                find_paths(n, path, True if n.islower() and n in path else small_twice)
                path.pop()


if __name__ == '__main__':

    mapper = {}
    paths = []

    with open('input.txt') as f:
        for line in f:
            line = line.strip()
            [a, b] = line.split('-')

            if a not in mapper:
                mapper[a] = []

            if b not in mapper:
                mapper[b] = []

            mapper[a].append(b)
            mapper[b].append(a)

    find_paths('start', [], False)
    print(len(paths))
