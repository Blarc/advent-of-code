from pprint import pprint

if __name__ == '__main__':

    chars = ['a', 'b', 'c', 'd', 'e', 'f']

    # 1 - 2
    # 7 - 3
    # 4 - 4
    # 2 - 5
    # 3 - 5
    # 5 - 5
    # 0 - 6
    # 6 - 6
    # 9 - 6
    # 8 - 7

    result = 0
    with open('input.txt') as f:
        for line in f:
            line = line.strip()

            input_, output = line.split(' | ')
            input_ = list(map(lambda x: ''.join(sorted(x)), input_.split()))
            output = list(map(lambda x: ''.join(sorted(x)), output.split()))

            one = next(x for x in input_ if len(x) == 2)
            seven = next(x for x in input_ if len(x) == 3)
            four = next(x for x in input_ if len(x) == 4)
            eight = next(x for x in input_ if len(x) == 7)

            two_three_five = list(x for x in input_ if len(x) == 5)
            zero_six_nine = list(x for x in input_ if len(x) == 6)

            six = next(filter(lambda x, one_=one: not set(one_).issubset(x), zero_six_nine))
            zero_six_nine.remove(six)
            zero_nine = zero_six_nine

            nine = next(filter(lambda x, four_=four: set(four_).issubset(x), zero_nine))
            zero = next(filter(lambda x, four_=four: not set(four_).issubset(x), zero_nine))

            three = next(filter(lambda x, one_=one: set(one_).issubset(x), two_three_five))
            two_three_five.remove(three)
            two_five = two_three_five

            five = next(filter(lambda x, nine_=nine: len(set(nine_).difference(x)) == 1, two_five))
            two = next(filter(lambda x, nine_=nine: len(set(nine_).difference(x)) == 2, two_five))

            mapper = {
                one: 1,
                seven: 7,
                four: 4,
                eight: 8,
                six: 6,
                nine: 9,
                zero: 0,
                three: 3,
                five: 5,
                two: 2
            }

            result += int(''.join(map(lambda x, m=mapper: str(m[x]) if x in m else 'x', output)))

    print(result)
