if __name__ == '__main__':

    rules = {}

    pairs_one = {}
    pairs_two = {}

    edges = {}
    occurrences = {}

    with open('input.txt') as f:

        template = list(f.readline().strip())
        f.readline()

        for line in f:
            line = list(map(lambda x: x.strip(), line.strip().split('->')))
            edges[line[0]] = [line[0][0] + line[1], line[1] + line[0][1]]
            rules[line[0]] = line[1]
            pairs_one[line[0]] = 0
            pairs_two[line[0]] = 0
            if line[1] not in occurrences:
                occurrences[line[1]] = 0

    for x in template:
        occurrences[x] += 1

    for i in range(len(template) - 1):
        pair = ''.join(template[i:i + 2])
        pairs_one[pair] += 1

    for step in range(40):
        if step % 2 == 0:
            for pair in pairs_one:
                if pairs_one[pair] != 0:
                    occurrences[rules[pair]] += pairs_one[pair]

                    two = edges[pair]
                    pairs_two[two[0]] += pairs_one[pair]
                    pairs_two[two[1]] += pairs_one[pair]

                    pairs_one[pair] = 0

        else:
            for pair in pairs_two:
                if pairs_two[pair] != 0:
                    occurrences[rules[pair]] += pairs_two[pair]

                    two = edges[pair]
                    pairs_one[two[0]] += pairs_two[pair]
                    pairs_one[two[1]] += pairs_two[pair]

                    pairs_two[pair] = 0

    print(occurrences)
    print(max(occurrences.items(), key=lambda x: x[1])[1] - min(occurrences.items(), key=lambda x: x[1])[1])

