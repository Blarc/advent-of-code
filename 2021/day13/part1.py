import matplotlib.pyplot as plt

if __name__ == '__main__':
    folds = []
    m = set()
    with open('input.txt') as f:
        for line in f:
            line = line.strip()
            if line.startswith('fold'):
                sf = line.split(' ')[2]
                f = sf.split('=')
                folds.append((f[0], int(f[1])))
            elif line == '':
                continue
            else:
                tmp = tuple(map(int, line.split(',')))
                m.add(tmp)

    for fold in folds:
        if fold[0] == 'y':
            y = fold[1]
            tmp = set()
            removes = set()
            for pointx, pointy in m:
                if pointy > y:
                    removes.add((pointx, pointy))
                    diff = pointy - y
                    tmp.add((pointx, y - diff))

            m.update(tmp)
            m.difference_update(removes)

        if fold[0] == 'x':
            x = fold[1]
            tmp = set()
            removes = set()
            for pointx, pointy in m:
                if pointx > x:
                    removes.add((pointx, pointy))
                    diff = pointx - x
                    tmp.add((x - diff, pointy))

            m.update(tmp)
            m.difference_update(removes)

    xa = []
    ya = []
    for point in m:
        xa.append(point[0])
        ya.append(point[1])

    plt.scatter(xa, ya)
    plt.axis('equal')
    plt.gca().invert_yaxis()
    plt.show()
