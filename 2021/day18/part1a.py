import math
from typing import Union

from llist import dllist, dllistnode


class Pair:
    def __init__(self, depth):
        self.left: Union[Pair, int, None] = None
        self.right: Union[Pair, int, None] = None
        self.parent: Union[Pair, None] = None
        self.depth: int = depth
        self.is_right = None

    def __str__(self, depth=False):
        if depth:
            if isinstance(self.left, Pair) and isinstance(self.right, Pair):
                return f'[{self.left.__str__(True)},{self.right.__str__(True)} ({self.depth})]'
            elif isinstance(self.left, Pair):
                return f'[{self.left.__str__(True)},{self.right.__str__()} ({self.depth})]'
            elif isinstance(self.right, Pair):
                return f'[{self.left.__str__()},{self.right.__str__(True)} ({self.depth})]'
            else:
                return f'[{self.left.__str__()},{self.right.__str__()} ({self.depth})]'

        return f'[{self.left.__str__()},{self.right.__str__()}]'

    def magnitude(self):
        if isinstance(self.left, int):
            left = self.left
        else:
            left = self.left.magnitude()

        if isinstance(self.right, int):
            right = self.right
        else:
            right = self.right.magnitude()

        return 3 * left + 2 * right


def create_tree(a: str) -> Pair:
    stack = []
    is_right = False
    current_pair = None
    for char in a:
        if char == '[':
            if current_pair is None:
                current_pair = Pair(0)
            else:
                new_pair = Pair(len(stack) + 1)
                new_pair.parent = current_pair
                new_pair.is_right = is_right

                if is_right:
                    current_pair.right = new_pair
                    is_right = False
                else:
                    current_pair.left = new_pair

                stack.append(current_pair)
                current_pair = new_pair
        elif char == ',':
            is_right = True
        elif char == ']':
            is_right = False
            if len(stack) > 0:
                current_pair = stack.pop()
        else:
            integer = int(char)
            if is_right:
                current_pair.right = integer
            else:
                current_pair.left = integer

    return current_pair


def find_left(n: dllistnode):
    node: dllistnode = n.prev
    while node is not None:
        if isinstance(node.value, int):
            return node
        node = node.prev

    return None


def find_right(n: dllistnode):
    node: dllistnode = n.next
    while node is not None:
        if isinstance(node.value, int):
            return node
        node = node.next

    return None


def explode(a: dllist):
    depth = 0
    node: dllistnode = a.first
    while node is not None:
        val = node.value

        if val == '[':
            depth += 1
        elif val == ']':
            depth -= 1

        if depth == 5:
            node = node.next
            a.remove(node.prev)

            left = find_left(node)
            if left is not None:
                left.value += node.value

            node.value = 0
            node = node.next

            right = find_right(node)
            if right is not None:
                right.value += node.value

            node = node.next
            a.remove(node.prev)

            node = node.next
            a.remove(node.prev)

            return True

        node = node.next

    return False


def split(a: dllist):
    node: dllistnode = a.first
    while node is not None:
        val = node.value

        if isinstance(val, int) and val >= 10:
            left = math.floor(val / 2)
            right = math.ceil(val / 2)

            node = node.prev
            a.remove(node.next)

            node = a.insertafter('[', node)
            node = a.insertafter(left, node)
            node = a.insertafter(right, node)
            a.insertafter(']', node)

            return True

        node = node.next

    return False


def add(a: dllist, b: dllist):
    a.extend(b)
    a.appendleft('[')
    a.appendright(']')

    # print(f'after addition: ')
    # print(to_str(a))

    while True:
        has_exploded = explode(a)
        # print(f'after explode: ')
        # print(to_str(a))
        while has_exploded:
            has_exploded = explode(a)
            # if has_exploded:
            #     print(f'after explode: ')
            #     print(to_str(a))
        if not split(a):
            break
        # print(f'after split: ')
        # print(to_str(a))

    return a


def to_str(a: dllist):
    res = ''
    node: dllistnode = a.first
    while node is not None:
        if (isinstance(node.value, int) or node.value == ']') and \
                node.next is not None and \
                (isinstance(node.next.value, int) or node.next.value == '['):
            res += f'{node.value},'
        else:
            res += f'{node.value}'

        node = node.next
    return res


if __name__ == '__main__':

    numbers = []
    with open('input.txt') as f:

        for line in f:
            line = line.strip()
            linked_list = dllist()

            previous_int = False
            for c in line:
                if c == '[' or c == ']':
                    linked_list.append(c)
                    previous_int = False
                elif c == ',':
                    previous_int = False
                else:
                    if previous_int:
                        linked_list[-1] *= 10
                        linked_list[-1] += int(c)
                    else:
                        linked_list.append(int(c))
                        previous_int = True

            numbers.append(linked_list)

    result = numbers[0]
    for number in numbers[1:]:
        result = add(result, number)

    tree = create_tree(to_str(result))
    print(tree.magnitude())
