def read_literal_value(index_):
    number_ = ''
    while True:
        group = line[index_:index_ + 5]
        index_ += 5
        number_ += group[1:]

        if group[0] == '0':
            # print(f'{number_} : {len(number_)}')
            print(f'Number: {int(number_, 2)}')
            break

    return index_, number_


def read_packet_version(index_):
    packet_version_ = line[index_:index_ + 3]
    index_ += 3
    packet_version_ = int(packet_version_, 2)
    print(f'Packet version: {packet_version_}')
    return index_, packet_version_


def read_packet_type(index_):
    packet_type_ = line[index_:index_ + 3]
    index_ += 3
    packet_type_ = int(packet_type_, 2)
    print(f'Packet type: {packet_type_}')
    return index_, packet_type_


def read_length_type(index_):
    length_type_ = line[index_]
    print(f'Length type: {length_type_}')
    return index_ + 1, int(length_type_, 2)


def read_total_length(index_):
    total_length_ = line[index_:index_ + 15]
    index_ += 15
    total_length_ = int(total_length_, 2)
    print(f'Total length: {total_length_}')
    return index_, total_length_


def read_number_subpackets(index_):
    number_sub_packets_ = line[index_:index_ + 11]
    index_ += 11
    number_sub_packets_ = int(number_sub_packets_, 2)
    print(f'Number of sub-packets: {number_sub_packets_}')
    return index_, number_sub_packets_


def decode(encoded_string, i):
    i, packet_version = read_packet_version(i)
    global packet_version_sum
    packet_version_sum += packet_version

    i, packet_type = read_packet_type(i)
    if packet_type == 4:
        i, number = read_literal_value(i)
        return i, number

    i, length_type_id = read_length_type(i)
    if length_type_id == 0:
        i, total_length = read_total_length(i)
        saved_i = i
        while i - saved_i != total_length:
            i, number = decode(encoded_string, i)

    elif length_type_id == 1:
        i, number_sub_packets = read_number_subpackets(i)
        for _ in range(number_sub_packets):
            i, number = decode(encoded_string, i)
    else:
        print('Somethings wrong!')

    return i, None


if __name__ == '__main__':
    with open('test.txt') as f:
        line = f.readline().strip()

    line = format(int(line, 16), f'0{len(line) * 4}b')
    size = len(line)

    index = 0
    packet_version_sum = 0

    decode(line, 0)

    print(packet_version_sum)
