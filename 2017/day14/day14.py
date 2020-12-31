from typing import List, Set, Tuple

from day10.day10 import dense_knot_hash

day_num = "14"
day_title = "Disk Defragmentation"

INPUT = "amgozmfv"


def dkh(p: str) -> str:
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


def num_regions(grid: Set[Tuple[int, int]], n: int) -> List[Set[Tuple[int, int]]]:
    all_regions: List[Set[Tuple[int, int]]] = []
    for i in range(n):
        for j in range(n):
            p = (i, j)
            if p in grid:
                # find the set within all_groups that p belongs too
                s: Set[Tuple[int, int]] = set()
                for g in all_regions:
                    if p in g:
                        s = g

                # if not found create a new group add to all_groups
                if len(s) == 0:
                    s.add(p)
                    all_regions.append(s)

                # check the neighbors
                neighbors: List[Tuple[int, int]] = [(i, j+1), (i, j-1), (i+1, j), (i-1, j)]
                for k in neighbors:
                    if k in grid:
                        s.add(k)

    return all_regions


def reduceSets(sets: List[Set[Tuple[int, int]]]) -> int:
    removed: Set[int] = set()
    for i in range(len(sets)-1):
        for j in range(1, len(sets)):
            if j in removed or i in removed:
                continue
            s1 = sets[i]
            s2 = sets[j]
            if len(s2.intersection(s1)) != 0:
                s1 = s1.union(s2)
                sets[i] = s1
                removed.add(j)

    return len(sets) - len(removed)

def main():
    print("Day", day_num, ":", day_title)
    squares, grid = squares_used(INPUT, 128)
    print("Part 1:", squares)
    all_regions = num_regions(grid, 128)
    reduced_regions = reduceSets(all_regions)
    print("Part 2:", reduced_regions)


if __name__ == '__main__':
    main()
