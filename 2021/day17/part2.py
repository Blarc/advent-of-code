import itertools

if __name__ == '__main__':
    with open('input.txt') as f:
        line = f.readline()

    line = line.strip()
    line = line[len('target area: '):].split(', ')

    target_x = list(map(int, line[0].split('=')[1].split('..')))
    target_y = list(map(int, line[1].split('=')[1].split('..')))

    max_steps = abs(target_y[0]) * 2

    map_y = {}
    for i in range(target_y[0], abs(target_y[0])):
        for step in range(max_steps + 1):
            start = sum(range(i, i - step, -1))

            if target_y[0] <= start <= target_y[1]:
                if step not in map_y:
                    map_y[step] = []
                map_y[step].append(i)

    min_x_vel = 0
    while not (min_x_vel * (min_x_vel + 1)) / 2 >= target_x[0]:
        min_x_vel += 1

    map_x = {x: [] for x in map_y}
    for i in range(min_x_vel, target_x[1] + 1):
        if i * (i + 1) / 2 <= target_x[1]:
            for key in map_x:
                if key > i:
                    map_x[key].append(i)

        for j in range(i):
            if target_x[0] <= i * (i + 1) / 2 - j * (j + 1) / 2 <= target_x[1]:
                map_x[i - j].append(i)

    result = set()
    for x in map_y:
        result.update(set(itertools.product(map_x[x], map_y[x])))

    print(len(result))
