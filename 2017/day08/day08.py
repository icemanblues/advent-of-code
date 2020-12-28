from typing import List, Dict, Tuple

day_num = "08"
day_title = "I Heard You Like Registers"


def read_input(filename: str) -> List[str]:
    with open(filename) as f:
        content = [x.strip('\n') for x in f.readlines()]
    return content


def process(lines: List[str]) -> Tuple[int, int]:
    registers: Dict[str, int] = {}
    max2 = 0
    for line in lines:
        words = line.split()
        cmd_reg = words[0]
        cmd = words[1]
        cmd_value = int(words[2])
        # 3 is 'if', worthless
        reg = words[4]
        op = words[5]
        value = int(words[6])

        v = registers.get(reg, 0)
        b = False
        if op == '>':
            b = v > value
        elif op == '<':
            b = v < value
        elif op == '>=':
            b = v >= value
        elif op == '<=':
            b = v <= value
        elif op == '==':
            b = v == value
        elif op == '!=':
            b = v != value
        else:
            print(f'Unknown operator in if statement: {op}')

        if b:
            v = registers.get(cmd_reg, 0)
            if cmd == 'inc':
                registers[cmd_reg] = v + cmd_value
            elif cmd == 'dec':
                registers[cmd_reg] = v - cmd_value
            else:
                print(f'Unknown cmd operator: {cmd}')

        if max2 < registers.get(cmd_reg, 0):
            max2 = registers[cmd_reg]

    max1 = 0
    for v in registers.values():
        if max1 < v:
            max1 = v
            
    return max1, max2


def main():
    print(f"Day {day_num}: {day_title}")
    lines = read_input("input.txt")
    p1, p2 = process(lines)
    print("Part 1:", p1)
    print("Part 2:", p2)


if __name__ == '__main__':
    main()
