from typing import List

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
    return spin[(f+1) % len(spin)]


def part1():
    print("Part 1:", spinlock(INPUT, 2017))


def zerolock(step_size: int, num_inserts: int) -> int:
    l = 2
    idx = 1
    next_to_zero = 1
    for x in range(2, num_inserts+1):
        idx = (idx+step_size) % l + 1
        l += 1
        if idx == 1:
            next_to_zero = x
    return next_to_zero


def part2():
    print("Part 2:", zerolock(INPUT, 50_000_000))


def main():
    print(f"Day {day_num}: {day_title}")
    part1()
    part2()


if __name__ == '__main__':
    main()
