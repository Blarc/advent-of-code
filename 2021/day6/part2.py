if __name__ == '__main__':

    queue = [0] * 9

    with open('input.txt') as f:
        line = f.readline()
        for x in map(int, line.strip().split(',')):
            queue[x] += 1

    for i in range(256):
        tmp = queue.pop(0)
        queue[6] += tmp
        queue.append(tmp)

    print(sum(queue))
