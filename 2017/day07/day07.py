from typing import List, Set, Dict, Tuple

day_num = "07"
day_title = "Recursive Circus"


def read_input(filename: str) -> List[str]:
    with open(filename) as f:
        content = [x.strip('\n') for x in f.readlines()]
    return content


class Prog:
    def __init__(self, name: str, weight: int, children: Set[str]):
        self.name = name
        self.weight = weight
        self.children = children


def parseProg(lines: List[str]) -> List[Prog]:
    progs: List[Prog] = []
    for line in lines:
        parts = line.split(' -> ')
        name_weight = parts[0].split(" (")
        name = name_weight[0]
        weight = int(name_weight[1][:-1])
        children: Set[str] = set()
        if len(parts) > 1:
            cc = parts[1].split(', ')
            for c in cc:
                children.add(c)
        progs.append(Prog(name, weight, children))
    return progs


def part1(progs: List[Prog]) -> Prog:
    start = progs[0]
    search = True
    while search:
        search = False
        for p in progs:
            if start.name in p.children:
                start = p
                search = True
    return start


def treeWeight(progsByName: Dict[str, Prog], root: Prog, memo: Dict[str, int]) -> int:
    if root.name in memo:
        return memo[root.name]

    w = root.weight
    for c in root.children:
        child = progsByName[c]
        w += treeWeight(progsByName, child, memo)
    memo[root.name] = w
    return w


def findBadProg(root: Prog, diff: int, progsByName: Dict[str, Prog], memo: Dict[str, int]) -> Tuple[str, int]:
    # get the sub-tree weights of all the children
    cw: Dict[str, int] = {}
    for c in root.children:
        cw[c] = memo[c]

    # find the odd man out
    c1 = list(root.children)[0]
    w1 = memo[c1]
    oddChildName = ""
    oddChildDiff = 0
    for c, w in cw.items():
        if w - w1 != 0:
            oddChildName = c
            oddChildDiff = w - w1

    # if the name is blank, its you
    if oddChildName == "":
        return root.name, diff
    else:
        oddChild = progsByName[oddChildName]
        return findBadProg(oddChild, oddChildDiff, progsByName, memo)


def part2(progs: List[Prog], root: Prog) -> int:
    progsByName: Dict[str, Prog] = {}
    for prog in progs:
        progsByName[prog.name] = prog

    memo: Dict[str, int] = {}
    for prog in progs:
        treeWeight(progsByName, prog, memo)

    bad, diff = findBadProg(root, 0, progsByName, memo)
    badProg = progsByName[bad]
    return badProg.weight - diff


def main():
    print(f"Day {day_num}: {day_title}")
    lines = read_input('input.txt')
    progs = parseProg(lines)
    root = part1(progs)
    print(f"Part 1: {root.name}")
    weight = part2(progs, root)
    print(f'Part 2: {weight}')


if __name__ == '__main__':
    main()
