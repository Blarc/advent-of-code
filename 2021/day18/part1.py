import math
from typing import Union


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

    def increase_depth(self):
        self.depth += 1
        if isinstance(self.left, Pair):
            self.left.increase_depth()
        if isinstance(self.right, Pair):
            self.right.increase_depth()

    def add(self, b):
        self.increase_depth()
        b.increase_depth()

        result = Pair(0)

        result.left = self
        self.parent = result

        result.right = b
        b.parent = result

        print_depth = False
        print(f'after addition: {result.__str__(print_depth)}')

        while True:
            _, _, has_exploded = result.explode()
            print(f'after explode: {result.__str__(print_depth)}')
            while has_exploded:
                _, _, has_exploded = result.explode()
                if has_exploded:
                    print(f'after explode: {result.__str__(print_depth)}')

            if not result.split():
                break
            print(f'after split: {result.__str__(print_depth)}')

        return result

    # Returns pair and take_right
    def find_left(self, is_right_explosion):

        if is_right_explosion:
            if isinstance(self.left, Pair):
                return self.left, True
            elif isinstance(self.left, int):
                return self, False
            else:
                raise ValueError('find_left: self.left is neither Pair or int')
        elif isinstance(self.left, int):
            return self, False
        elif self.parent:
            return self.parent.find_left(is_right_explosion)
        # else:
        #     tmp = self.parent
        #     while tmp is not None:
        #         if isinstance(tmp.left, int):
        #             return tmp, False
        #         elif tmp.is_right is not None and tmp.is_right and isinstance(tmp.left, Pair):
        #             return tmp.left, True
        #         tmp = tmp.parent
        # 
        return None, False

    # Returns pair and take_left
    def find_right(self, is_right_explosion):

        if not is_right_explosion:
            if isinstance(self.right, Pair):
                return self.right, True
            elif isinstance(self.right, int):
                return self, False
            else:
                raise ValueError('find_left: self.left is neither Pair or int')
        elif isinstance(self.right, int):
            return self, False
        elif self.parent:
            return self.parent.find_right(is_right_explosion)

        # else:
        #     tmp = self.parent
        #     while tmp is not None:
        #         if isinstance(tmp.right, int):
        #             return tmp, False
        #         elif tmp.is_right is not None and not tmp.is_right and isinstance(tmp.right, Pair):
        #             return tmp.right, True
        #         tmp = tmp.parent
        # 
        return None, False

    def explode(self) -> Union[tuple, None]:
        if self.depth >= 4:
            return self.left, self.right, True

        has_exploded_left = False
        if isinstance(self.left, Pair):
            left, right, has_exploded_left = self.left.explode()
            self.after_explode(left, right, False)

        has_exploded_right = False
        if isinstance(self.right, Pair):
            left, right, has_exploded_left = self.right.explode()
            self.after_explode(left, right, True)

        return None, None, has_exploded_left or has_exploded_right

    def after_explode(self, left, right, is_right_explosion):
        if left is not None:
            first_left, take_right = self.find_left(is_right_explosion)
            if first_left != self and first_left != self.left:
                self.left = 0
            if first_left is not None:
                if take_right:
                    first_left.right = left + first_left.right
                else:
                    first_left.left = left + first_left.left

        if right is not None:
            first_right, take_left = self.find_right(is_right_explosion)
            if first_right != self and first_right != self.right:
                self.right = 0
            if first_right is not None:
                if take_left:
                    first_right.left = right + first_right.left
                else:
                    first_right.right = right + first_right.right

    def split(self):
        left_split = False
        if isinstance(self.left, int):
            if self.left >= 10:
                p = Pair(self.depth + 1)
                p.left = math.floor(self.left / 2)
                p.right = math.ceil(self.left / 2)
                p.parent = self
                self.left = p
                left_split = True
        else:
            left_split = left_split or self.left.split()

        right_split = False
        if isinstance(self.right, int):
            if self.right >= 10:
                p = Pair(self.depth + 1)
                p.left = math.floor(self.right / 2)
                p.right = math.ceil(self.right / 2)
                p.parent = self
                self.right = p
                right_split = True
        else:
            right_split = right_split or self.right.split()

        return left_split or right_split


if __name__ == '__main__':

    numbers = []
    with open('test.txt') as f:

        for line in f:
            stack = []
            line = line.strip()

            is_right = False
            current_pair = None
            for c in line:
                if c == '[':
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
                elif c == ',':
                    is_right = True
                elif c == ']':
                    is_right = False
                    if len(stack) > 0:
                        current_pair = stack.pop()
                else:
                    integer = int(c)
                    if is_right:
                        current_pair.right = integer
                    else:
                        current_pair.left = integer

            numbers.append(current_pair)

    start = numbers[0]
    for number in numbers[1:]:
        start = start.add(number)

    print(start)
