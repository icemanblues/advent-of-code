from typing import List, Set, Tuple

from day10.day10 import dense_knot_hash

day_num = "14"
day_title = "Disk Defragmentation"

INPUT = "amgozmfv"


def dkh(p: str) -> str:
    return dense_knot_hash(list(range(256)), p)


def toBin(hex: str) -> str:
    return "{0:8b}".format(int(hex, 16))


def squares_used(puzzle: str, n: int) -> Set[Tuple[int, int]]:
    grid = set()
    for i in range(n):
        p = puzzle + "-" + str(i)
        d = dkh(p)

        # toBin drops the leading zeroes
        b = toBin(d)
        while len(b) < n:
            b = "0" + b

        for j in range(n):
            bit = b[j]
            if(bit == "1"):
                grid.add((i, j))
    return grid


def num_regions(grid: Set[Tuple[int, int]], n: int) -> List[Set[Tuple[int, int]]]:
    all_regions: List[Set[Tuple[int, int]]] = []
    for i in range(n):
        for j in range(n):
            p = (i, j)
            if p in grid:
                # check if it already exists in a set
                exists = False
                for s in all_regions:
                    if p in s:
                        exists = True
                        break
                if exists:
                    continue

                # not in an existing set, so start a new one and grow it
                # do bfs to find all of the neighbors and add them to the set
                queue: List[Tuple[int, int]] = [p]
                visited: Set[Tuple[int, int]] = set()
                while len(queue) != 0:
                    q = queue.pop(0)
                    if q in visited:
                        continue
                    visited.add(q)

                    neighbors: List[Tuple[int, int]] = [(q[0], q[1]+1), (q[0], q[1]-1),
                                                        (q[0]+1, q[1]), (q[0]-1, q[1])]
                    for sq in neighbors:
                        if sq in grid:
                            queue.append(sq)
                all_regions.append(visited)
    return all_regions


def main():
    print("Day", day_num, ":", day_title)
    grid = squares_used(INPUT, 128)
    print("Part 1:", len(grid))
    all_regions = num_regions(grid, 128)
    print("Part 2:", len(all_regions))


if __name__ == '__main__':
    main()
