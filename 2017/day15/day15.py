from typing import Iterator, List

day_num = "15"
day_title = "Dueling Generators"


DIVISOR = 2147483647

def read_input(filename: str) -> List[str]:
    with open(filename) as f:
        content = [x.strip('\n') for x in f.readlines()]
    return content


def make_generator(value: int, factor: int) -> Iterator[int]:
    prev = value
    while True:
        value = (prev * factor) % DIVISOR
        yield value
        prev = value

def part1():
    a = make_generator(289, 16807)
    b = make_generator(629, 48271)

    bits = 0
    for i in range(16):
        bits += 1<<i

    count = 0
    for _ in range(40_000_000):
        a_value = next(a)
        b_value = next(b)
        if (a_value & bits) == (b_value & bits):
            count += 1
    
    print(f"Part 1:", count)


def part2():
    print("Part 2")


def main():
    print(f"Day {day_num}: {day_title}")
    part1()
    part2()


if __name__ == '__main__':
    main()
