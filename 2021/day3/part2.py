if __name__ == '__main__':
    oxygen = []
    co2 = []

    with open('input.txt') as f:

        for line in f:
            tmp = list(map(int, list(line.strip())))
            oxygen.append(tmp)
            co2.append(tmp)

        for i in range(len(oxygen[0])):
            ones = list(filter(lambda x, idx=i: x[idx] == 1, oxygen))
            if len(ones) >= len(oxygen) / 2:
                oxygen = ones
            else:
                oxygen = list(filter(lambda x, idx=i: x[idx] == 0, oxygen))

            if len(oxygen) == 1:
                break

        for i in range(len(co2[0])):
            ones = list(filter(lambda x, idx=i: x[idx] == 1, co2))
            if len(ones) < len(co2) / 2:
                co2 = ones
            else:
                co2 = list(filter(lambda x, idx=i: x[idx] == 0, co2))

            if len(co2) == 1:
                break


    oxygen = int(''.join(list(map(str, oxygen[0]))), 2)
    co2 = int(''.join(list(map(str, co2[0]))), 2)
    print(f'Gamma: {oxygen}')
    print(f'Epsilon: {co2}')
    print(f'Gamma * Epsilon: {oxygen * co2}')
