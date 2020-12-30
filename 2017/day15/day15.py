from typing import Iterator, List

day_num = "15"
day_title = "Dueling Generators"


DIVISOR = 2147483647


def read_input(filename: str) -> List[str]:
    with open(filename) as f:
        content = [x.strip('\n') for x in f.readlines()]
    return content


def make_generator(value: int, factor: int, mult: int) -> Iterator[int]:
    prev = value
    while True:
        value = (prev * factor) % DIVISOR
        if value % mult == 0:
            yield value
        prev = value


def judge(a: Iterator[int], b: Iterator[int], num_pairs: int) -> int:
    bits = 0
    for i in range(16):
        bits += 1 << i

    count = 0
    for _ in range(num_pairs):
        a_value = next(a)
        b_value = next(b)
        if (a_value & bits) == (b_value & bits):
            count += 1
    return count


def part1():
    a = make_generator(289, 16807, 1)
    b = make_generator(629, 48271, 1)
    count = judge(a, b, 40_000_000)
    print(f"Part 1: {count}")


def part2():
    a = make_generator(289, 16807, 4)
    b = make_generator(629, 48271, 8)
    count = judge(a, b, 5_000_000)
    print(f"Part 2: {count}")


def main():
    print(f"Day {day_num}: {day_title}")
    part1()
    part2()


if __name__ == '__main__':
    main()
