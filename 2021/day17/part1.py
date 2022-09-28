if __name__ == '__main__':
    with open('input.txt') as f:
        line = f.readline()

    line = line.strip()
    line = line[len('target area: '):].split(', ')

    target_x = list(map(int, line[0].split('=')[1].split('..')))
    target_y = list(map(int, line[1].split('=')[1].split('..')))
    print(target_x, target_y)

    x_vel = 0
    y_vel = 0

    while not (x_vel * (x_vel + 1)) / 2 >= target_x[0]:
        x_vel += 1

    y_vel = abs(target_y[0]) - 1
    print((y_vel * (y_vel + 1)) / 2)