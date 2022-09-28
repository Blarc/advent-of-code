from bresenham import bresenham


def intersect(l1, l2):
    x1, y1 = l1[0]
    x2, y2 = l1[1]
    x3, y3 = l2[0]
    x4, y4 = l2[1]
    denom = (y4 - y3) * (x2 - x1) - (x4 - x3) * (y2 - y1)

    # Check if parallel
    if denom == 0:
        if x2 - x1 == 0:
            slope = 0
        else:
            slope = (y2 - y1) / (x2 - x1)

        b = y1 - slope * x1

        # Check if on same line
        if y3 == slope * x3 + b and y4 == slope * x4 + b:
            start_x = max(min(x1, x2), min(x3, x4))
            end_x = min(max(x1, x2), max(x3, x4))
            start_y = max(min(y1, y2), min(y3, y4))
            end_y = min(max(y1, y2), max(y3, y4))

            return bresenham(start_x, start_y, end_x, end_y)
        return None

    ua = ((x4 - x3) * (y1 - y3) - (y4 - y3) * (x1 - x3)) / denom
    if ua < 0 or ua > 1:  # out of range
        return None
    ub = ((x2 - x1) * (y1 - y3) - (y2 - y1) * (x1 - x3)) / denom
    if ub < 0 or ub > 1:  # out of range
        return None
    x = x1 + ua * (x2 - x1)
    y = y1 + ua * (y2 - y1)
    return [(int(x), int(y))]


# def bresenham(x1, y1, x2, y2):
#     print((x1, y1), (x2, y2))
#     m_new = 2 * (y2 - y1)
#     slope_error_new = m_new - (x2 - x1)
# 
#     y = y1
#     result = []
#     for x in range(x1, x2 + 1):
# 
#         # print("(", x, ",", y, ")")
#         result.append((x, y))
# 
#         # Add slope to increment angle formed
#         slope_error_new = slope_error_new + m_new
# 
#         # Slope error reached limit, time to
#         # increment y and update slope error.
#         if slope_error_new >= 0:
#             y = y + 1
#             slope_error_new = slope_error_new - 2 * (x2 - x1)
# 
#     return result


if __name__ == '__main__':
    with open('input.txt') as f:

        points = {}
        for line in f:
            line = list(map(lambda x: tuple(map(int, x.split(','))), line.strip().split(' -> ')))
            if line[0][0] == line[1][0] or line[0][1] == line[1][1]:
                x1, y1, x2, y2 = line[0][0], line[0][1], line[1][0], line[1][1]
                bresenham_points = bresenham(x1, y1, x2, y2)
                for point in bresenham_points:
                    if point not in points:
                        points[point] = 1
                    else:
                        points[point] += 1

    print(len(list(filter(lambda x: x > 1, points.values()))))
