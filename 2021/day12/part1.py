def find_paths(node, path):
    if node == 'end':
        paths.append(path)
        return

    for n in mapper[node]:
        if n.isupper() or n not in path:
            path.append(node)
            find_paths(n, path)
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

    find_paths('start', [])
    print(len(paths))
