from typing import List, Dict, Set

day_num = "12"
day_title = "Digital Plumber"


def read_input(filename: str) -> Dict[int, List[int]]:
    pipemap: Dict[int, List[int]] = {}
    with open(filename) as f:
        content = [x.strip('\n') for x in f.readlines()]
        for line in content:
            parts = line.split(' <-> ')
            program = int(parts[0])
            nodes: List[int] = [int(x) for x in parts[1].split(', ')]
            pipemap[program] = nodes
    return pipemap


def bfs(start: int, end: int, pipemap: Dict[int, List[int]]) -> bool:
    queue: List[int] = []
    queue.extend(pipemap[start])
    visited: Set[int] = {start}
    while len(queue) != 0:
        q = queue.pop(0)
        if q == end:
            return True
        if q in visited:
            continue
        visited.add(q)
        queue.extend(pipemap[q])

    return False


def part1():
    pipemap = read_input('input.txt')
    count = 0
    for p in pipemap.keys():
        if bfs(p, 0, pipemap):
            count += 1

    print("Part 1:", count)


def part2():
    print("Part 2")


def main():
    print(f"Day {day_num}: {day_title}")
    part1()
    part2()


if __name__ == '__main__':
    main()
