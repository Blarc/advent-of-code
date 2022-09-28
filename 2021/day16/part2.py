from functools import reduce
from operator import mul


def read_literal_value(index_, encoded_string_):
    number_ = ''
    while True:
        group = encoded_string_[index_:index_ + 5]
        index_ += 5
        number_ += group[1:]

        if group[0] == '0':
            # print(f'Number: {int(number_, 2)}')
            break

    return index_, int(number_, 2)


def read_packet_version(index_, encoded_string_):
    packet_version_ = encoded_string_[index_:index_ + 3]
    index_ += 3
    packet_version_ = int(packet_version_, 2)
    # print(f'Packet version: {packet_version_}')
    return index_, packet_version_


def read_packet_type(index_, encoded_string_):
    packet_type_ = encoded_string_[index_:index_ + 3]
    index_ += 3
    packet_type_ = int(packet_type_, 2)
    # print(f'Packet type: {packet_type_}')
    return index_, packet_type_


def read_length_type(index_, encoded_string_):
    length_type_ = encoded_string_[index_]
    # print(f'Length type: {length_type_}')
    return index_ + 1, int(length_type_, 2)


def read_total_length(index_, encoded_string_):
    total_length_ = encoded_string_[index_:index_ + 15]
    index_ += 15
    total_length_ = int(total_length_, 2)
    # print(f'Total length: {total_length_}')
    return index_, total_length_


def read_number_subpackets(index_, encoded_string_):
    number_sub_packets_ = encoded_string_[index_:index_ + 11]
    index_ += 11
    number_sub_packets_ = int(number_sub_packets_, 2)
    # print(f'Number of sub-packets: {number_sub_packets_}')
    return index_, number_sub_packets_


def do_operation(packet_type, numbers_):
    if packet_type == 0:
        return sum(numbers_)
    elif packet_type == 1:
        return reduce(mul, numbers_, 1)
    elif packet_type == 2:
        return min(numbers_)
    elif packet_type == 3:
        return max(numbers_)
    elif packet_type == 5:
        return int(numbers_[0] > numbers_[1])
    elif packet_type == 6:
        return int(numbers_[0] < numbers_[1])
    elif packet_type == 7:
        return int(numbers_[0] == numbers_[1])


def decode(encoded_string, i):
    if i == len(encoded_string):
        return i, None

    i, packet_version = read_packet_version(i, encoded_string)
    i, packet_type = read_packet_type(i, encoded_string)

    if packet_type == 4:
        i, number = read_literal_value(i, encoded_string)
        return i, number

    i, length_type_id = read_length_type(i, encoded_string)
    if length_type_id == 0:
        i, total_length = read_total_length(i, encoded_string)
        saved_i = i
        numbers = []
        while i - saved_i != total_length:
            i, number = decode(encoded_string, i)
            numbers.append(number)

        return i, do_operation(packet_type, numbers)

    elif length_type_id == 1:
        i, number_sub_packets = read_number_subpackets(i, encoded_string)
        numbers = []
        for _ in range(number_sub_packets):
            i, number = decode(encoded_string, i)
            numbers.append(number)

        return i, do_operation(packet_type, numbers)

    else:
        print('Somethings wrong!')

    return i, None


if __name__ == '__main__':
    with open('input.txt') as f:
        line = f.readline().strip()

    line = format(int(line, 16), f'0{len(line) * 4}b')
    size = len(line)

    index = 0
    print(decode(line, 0)[1])
