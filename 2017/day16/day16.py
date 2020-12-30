from typing import List, Set

day_num = "16"
day_title = "Permutation Promenade"


DANCERS = "abcdefghijklmnop"


def read_input(filename: str) -> List[str]:
    with open(filename) as f:
        content = [x.strip('\n') for x in f.readlines()]
        steps = content[0].split(",")
    return steps


def spin(s: str, x: int) -> str:
    return s[-x:] + s[:-x]


def exchange(s: str, x: int, y: int) -> str:
    cx = s[x]
    cy = s[y]
    s = s.replace(cx, "1")
    s = s.replace(cy, cx)
    s = s.replace("1", cy)
    return s


def partner(s: str, a: str, b: str) -> str:
    i = s.find(a)
    j = s.find(b)
    return exchange(s, i, j)


def dance(dancers: str, steps: List[str]) -> str:
    for step in steps:
        cmd = step[0]
        arg = step[1:]
        if cmd == 's':
            dancers = spin(dancers, int(arg))
        elif cmd == 'x':
            idxs = arg.split('/')
            dancers = exchange(dancers, int(idxs[0]), int(idxs[1]))
        elif cmd == 'p':
            idxs = arg.split('/')
            dancers = partner(dancers, idxs[0], idxs[1])
        else:
            print("unknown dance step:", step)
    return dancers


def part1():
    steps = read_input('input.txt')
    dancers = dance(DANCERS, steps)
    print("Part 1:", dancers)


def part2():
    steps = read_input('input.txt')
    dancers = DANCERS

    dance_starts: List[str] = []
    dance_set: Set[str] = set()
    n = 1000000000
    for _ in range(n):
        if dancers in dance_set:
            break
        dance_starts.append(dancers)
        dance_set.add(dancers)
        dancers = dance(dancers, steps)

    answer = dance_starts[n % len(dance_starts)]
    print("Part 2:", answer)


def main():
    print(f"Day {day_num}: {day_title}")
    part1()
    part2()


if __name__ == '__main__':
    main()
