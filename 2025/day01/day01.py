def parse_input(filename: str) -> list[tuple[str]]:
    with open(filename, 'r') as f:
        return [(line[0], int(line[1:])) for line in f]


def part1() -> int:
    x = 50
    zero = 0
    input = parse_input('input.txt')
    for i in input:
        if i[0] == 'R':
            x = (x + i[1]) % 100
        else:
            x = (x - i[1]) % 100
        if x == 0:
            zero += 1
    return zero


def part2() -> int:
    x = 50
    zero = 0
    input = parse_input('test.txt')
    for i in input:
        dir = -1
        if i[0] == 'R':
            dir = 1
        
        diff = (x + dir * i[1])
        rotate = abs(diff) // 100
        xi = diff % 100
        if x > 0 and 100-diff > x:
            zero += 1 + rotate
        if x < 0 and diff > 0:
            zero += 1 + rotate
        if xi == 0:
            zero += 1

        print(f"{i}: {x} -> {xi} (+{zero})")
        x = xi
    return zero


print("part 1: ", part1())
print("part 2: ", part2())
