from typing import Dict, List

day_num = "17"
day_title = "Spinlock"

INPUT = 367


def spinlock(step_size: int, num_inserts: int) -> int:
    spin: List[int] = [0, 1]
    idx = 1
    for x in range(2, num_inserts+2):
        idx = (idx+step_size) % len(spin) + 1
        spin.insert(idx, x)

    f = spin.index(num_inserts)
    return spin[(f+1)%len(spin)]


def part1():
    print("Part 1:", spinlock(INPUT, 2017))


class Node:
    def __init__(self, d: int):
        self.data = d
        self.next = None


def nodelock(step_size: int, num_inserts: int) -> int:
    return 2


def part2():
    print("Part 2:", nodelock(INPUT, 50000000))


def main():
    print(f"Day {day_num}: {day_title}")
    part1()
    part2()


if __name__ == '__main__':
    main()
