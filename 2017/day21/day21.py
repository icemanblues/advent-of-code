import math

day_num = "21"
day_title = "Fractal Art"

start = ['.#.', '..#', '###']


def read_input(filename: str):
    rules = {}
    with open(filename) as f:
        for line in f.readlines():
            line = line.strip('\n')
            parts = line.split(' => ')
            left = parts[0].split('/')
            right = parts[1].split('/')
            rules[tuple(left)] = right
    return rules


def rot90(art):
    rot = []
    for x in range(len(art)):
        row = ''
        for y in range(len(art)):
            row += art[len(art)-1-y][x]
        rot.append(row)
    return rot


def flip(art):
    f = []
    for i in range(len(art)):
        f.append(art[len(art)-1-i])
    return f


def orientations(art):
    o = []
    c = art
    f = flip(c)
    for _ in range(4):
        o.append(c)
        o.append(f)
        c = rot90(c)
        f = rot90(f)
    return o


def tick(art, rules):
    new_pieces = []
    n = 2 if len(art) % 2 == 0 else 3
    for y in range(0, len(art), n):
        for x in range(0, len(art), n):
            sub = []
            for j in range(n):
                row = ''
                for i in range(n):
                    row += art[y+j][x+i]
                sub.append(row)
            new_pieces.append(rules[tuple(sub)])
    new_art = stitch(new_pieces)
    return new_art


def stitch(pieces):
    n = len(pieces)
    s = int(math.sqrt(n))
    new_art = []
    for idx in range(0, n, s):
        for y in range(len(pieces[0])):
            row = ''
            for i in range(s):
                row += pieces[idx+i][y]
            new_art.append(row)
    return new_art


def count(art):
    c = 0
    for y in range(len(art)):
        for x in range(len(art[y])):
            if art[y][x] == '#':
                c += 1
    return c


def rot_rules(rules):
    all_rules = {}
    for k, v in rules.items():
        for o in orientations(k):
            all_rules[tuple(o)] = v
    return all_rules


def iterate(n, all_rules):
    curr = start
    for _ in range(n):
        curr = tick(curr, all_rules)
    return count(curr)


def main():
    print(f"Day {day_num}: {day_title}")
    rules = read_input('input.txt')
    all_rules = rot_rules(rules)
    print('Part 1:', iterate(5, all_rules))
    print('Part 2:', iterate(18, all_rules))


if __name__ == '__main__':
    main()
