from bresenham import bresenham

if __name__ == '__main__':
    with open('input.txt') as f:

        points = {}
        for line in f:
            line = list(map(lambda x: tuple(map(int, x.split(','))), line.strip().split(' -> ')))
            x1, y1, x2, y2 = line[0][0], line[0][1], line[1][0], line[1][1]
            bresenham_points = bresenham(x1, y1, x2, y2)
            for point in bresenham_points:
                if point not in points:
                    points[point] = 1
                else:
                    points[point] += 1

    print(len(list(filter(lambda x: x > 1, points.values()))))
