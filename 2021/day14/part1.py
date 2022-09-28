from collections import Counter

if __name__ == '__main__':

    rules = {}
    with open('input.txt') as f:

        template = list(f.readline().strip())
        f.readline()

        for line in f:
            line = list(map(lambda x: x.strip(), line.strip().split('->')))
            rules[tuple(line[0])] = line[1]

    for _ in range(0, 10):
        for i in range(0, (len(template) - 1) * 2, 2):
            current = template[i], template[i + 1]
            template.insert(i + 1, rules[current])

    occurances = Counter(template)
    print(occurances)
    print(max(occurances.items(), key=lambda x: x[1])[1] - min(occurances.items(), key=lambda x: x[1])[1])
