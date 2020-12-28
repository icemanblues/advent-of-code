from typing import List, Set, Tuple

# FIXME: Need to figure out how to make this a proper python project, package, and module
from day10 import dense_knot_hash

day_num = "14"
day_title = "Disk Defragmentation"

PUZZLE = "amgozmfv"


def dkh(p: List[str]) -> str:
    return dense_knot_hash(list(range(256)), p)


def toBin(hex: str) -> str:
    return "{0:8b}".format(int(hex, 16))


def squares_used(puzzle: str, n: int) -> Tuple[int, Set[Tuple[int, int]]]:
    count = 0
    grid = set()
    for i in range(n):
        p = puzzle + "-" + str(i)
        d = dkh(p)
        b = toBin(d)

        for j in range(len(b)):
            bit = b[j]
            if(bit == "1"):
                count += 1
                grid.add((i, j))
    return count, grid


def num_regions(grid: Set[Tuple[int, int]], n: int) -> List[Set[int]]:
    all_regions: List[Set[int]] = []
    for i in range(n):
        for j in range(n):
            p = (i, j)
            if p in grid:
                # find the set within all_groups that p belongs too
                s = None
                for g in all_regions:
                    if p in g:
                        s = g

                # if not found create a new group add to all_groups
                if s is None:
                    s = set()
                    s.add(p)
                    all_regions.append(s)

                # check the neighbors
                n = [(i, j+1), (i, j-1), (i+1, j), (i-1, j)]
                for k in n:
                    if k in grid:
                        s.add(k)

    return all_regions


def main():
    print("Day", day_num, ":", day_title)
    squares, grid = squares_used(PUZZLE, 128)
    print("Part 1:", squares)
    all_regions = num_regions(grid, 128)
    print("Part 2:", len(all_regions))


if __name__ == '__main__':
    main()
