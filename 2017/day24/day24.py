from typing import List, Tuple, Dict

day_num = "24"
day_title = "Electromagnetic Moat"


def read_input(filename: str) -> List[Tuple[int, int]]:
    ports: List[Tuple[int, int]] = []
    with open(filename) as f:
        for line in f.readlines():
            line = line.strip('\n')
            parts = line.split('/')
            ports.append((int(parts[0]), int(parts[1])))
    return ports


def strength(ports: List[Tuple[int, int]]) -> int:
    sum = 0
    for p in ports:
        sum += p[0] + p[1]
    return sum


def backtrack(must_match: int,
              ports: List[Tuple[int, int]],
              sol: List[Tuple[int, int]],
              all_sol: List[List[Tuple[int, int]]]) -> List[Tuple[int, int]]:
    if len(ports) == 0:
        return sol

    possible: List[Tuple[int, int]] = []
    for p in ports:
        if p[0] == must_match or p[1] == must_match:
            possible.append(p)
    for p in possible:
        sol.append(p)
        ports.remove(p)
        match = p[0] if p[1] == must_match else p[1]
        ans = backtrack(match, ports, sol, all_sol)
        all_sol.append(list(ans))
        sol.remove(p)
        ports.append(p)

    return sol


def find_max_str(all_sols: List[List[Tuple[int, int]]]) -> int:
    max_str: int = 0
    for solution in all_sols:
        s = strength(solution)
        if s > max_str:
            max_str = s
    return max_str


def build_bridges():
    ports = read_input('input.txt')

    # backtrack
    must_match: int = 0
    sol: List[Tuple[int, int]] = []
    all_sols: List[List[Tuple[int, int]]] = []
    backtrack(must_match, ports, sol, all_sols)

    # find the max strength of all solutions
    print("Part 1:", find_max_str(all_sols))

    # find the longest of all solutions with max strength
    lenToBridge: Dict[int, List[List[Tuple[int, int]]]] = {}
    max_len = 0
    for solution in all_sols:
        l = len(solution)
        bridge: List[List[Tuple[int, int]]] = lenToBridge.get(l, [])
        bridge.append(solution)
        lenToBridge[l] = bridge
        if l > max_len:
            max_len = l
    print('Part 2:', find_max_str(lenToBridge[max_len]))


def main():
    print(f"Day {day_num}: {day_title}")
    build_bridges()


if __name__ == '__main__':
    main()
