if __name__ == '__main__':
    with open('input.txt') as f:
        count = 0
        prev_window_sum = float('inf')

        window = []
        for index, line in enumerate(f):
            current = int(line)

            if index < 3:
                window.append(current)
            else:
                window.pop(0)
                window.append(current)

                window_sum = sum(window)
                if window_sum > prev_window_sum:
                    count += 1

                prev_window_sum = window_sum

    print(count)
