from typing import Dict, List, Tuple

day_num = "03"
day_title = "Spiral Memory"

INPUT = 312051

up = 1
left = 2
down = 3
right = 4


def rotateLeft(dir: int) -> int:
    if dir == up:
        return left
    elif dir == left:
        return down
    elif dir == down:
        return right
    elif dir == right:
        return up
    else:
        print("unknown direction:", dir)
        return -1


def move(p: Tuple[int, int], dir: int) -> Tuple[int, int]:
    if dir == up:
        return p[0], p[1]+1
    elif dir == left:
        return p[0]-1, p[1]
    elif dir == down:
        return p[0], p[1]-1
    elif dir == right:
        return p[0]+1, p[1]
    else:
        print("unknown direction:", dir)
        return p


def part1(n: int) -> int:
    spiral: Dict[Tuple[int, int], int] = {}
    dir = right
    w = 1
    p = (0, 0)
    spiral[p] = w
    w = 2
    p = (1, 0)
    spiral[p] = w
    while w < n:
        turn = rotateLeft(dir)
        q = move(p, turn)
        if q in spiral:
            q = move(p, dir)
        else:
            dir = turn
        p = q
        w += 1
        spiral[p] = w

    return abs(p[0])+abs(p[1])


def adj(p: Tuple[int, int]) -> List[Tuple[int, int]]:
    l: List[Tuple[int, int]] = []
    for x in range(-1, 2):
        for y in range(-1, 2):
            if x != 0 or y != 0:
                l.append((p[0]+x, p[1]+y))
    return l


def part2(n: int) -> int:
    spiral: Dict[Tuple[int, int], int] = {}
    dir = right
    w = 1
    p = (0, 0)
    spiral[p] = w
    w = 1
    p = (1, 0)
    spiral[p] = w
    while w < n:
        turn = rotateLeft(dir)
        q = move(p, turn)
        if q in spiral:
            q = move(p, dir)
        else:
            dir = turn
        p = q

        sum = 0
        for a in adj(p):
            if a in spiral:
                sum += spiral[a]
        w = sum
        spiral[p] = w

    return w


def main():
    print(f"Day {day_num}: {day_title}")
    print("Part 1", part1(INPUT))
    print("Part 2", part2(INPUT))


if __name__ == '__main__':
    main()
