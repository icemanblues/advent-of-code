from typing import List, Set, Tuple, Dict

day_num = "22"
day_title = "Sporifica Virus"

UP = (0, -1)
DOWN = (0, 1)
LEFT = (-1, 0)
RIGHT = (1, 0)


def turnRight(d: Tuple[int, int]) -> Tuple[int, int]:
    if d == UP:
        return RIGHT
    elif d == RIGHT:
        return DOWN
    elif d == DOWN:
        return LEFT
    elif d == LEFT:
        return UP
    else:
        print('unknown direction:', d)
        return d


def turnLeft(d: Tuple[int, int]) -> Tuple[int, int]:
    if d == UP:
        return LEFT
    elif d == RIGHT:
        return UP
    elif d == DOWN:
        return RIGHT
    elif d == LEFT:
        return DOWN
    else:
        print('unknown direction:', d)
        return d


def read_input(filename: str) -> Tuple[Set[Tuple[int, int]], Tuple[int, int]]:
    grid: Set[Tuple[int, int]] = set()
    with open(filename) as f:
        y = 0
        xmax = -1
        for line in f.readlines():
            line = line.strip()
            x = 0
            for s in line:
                if s == '#':
                    grid.add((x, y))
                x += 1
            y += 1
            xmax = x
        start = (xmax//2, y//2)
    return grid, start


def part1():
    grid, curr = read_input('input.txt')
    direct: Tuple[int, int] = UP

    infect_count = 0
    for _ in range(10000):
        infected = curr in grid
        if infected:
            direct = turnRight(direct)
            grid.remove(curr)
        else:
            direct = turnLeft(direct)
            grid.add(curr)
            infect_count += 1
        curr = curr[0] + direct[0], curr[1] + direct[1]

    print("Part 1:", infect_count)


WEAKENED = 'W'
INFECTED = 'I'
FLAGGED = 'F'


def reverse(d: Tuple[int, int]) -> Tuple[int, int]:
    if d == UP:
        return DOWN
    elif d == DOWN:
        return UP
    elif d == LEFT:
        return RIGHT
    elif d == RIGHT:
        return LEFT
    else:
        print("unknown direction:", d)
        return d


def part2():
    grid, curr = read_input('input.txt')
    direct: Tuple[int, int] = UP

    state: Dict[Tuple[int, int], str] = {}
    for p in grid:
        state[p] = INFECTED

    infect_count = 0
    for _ in range(10000000):
        clean = curr not in state
        if clean:
            direct = turnLeft(direct)
            state[curr] = WEAKENED
        elif state[curr] == WEAKENED:
            state[curr] = INFECTED
            infect_count += 1
        elif state[curr] == INFECTED:
            direct = turnRight(direct)
            state[curr] = FLAGGED
        elif state[curr] == FLAGGED:
            direct = reverse(direct)
            del state[curr]

        curr = curr[0] + direct[0], curr[1] + direct[1]

    print("Part 2:", infect_count)


def main():
    print(f"Day {day_num}: {day_title}")
    part1()
    part2()


if __name__ == '__main__':
    main()
