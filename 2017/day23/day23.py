from typing import Dict, List

day_num = "23"
day_title = "Coprocessor Conflagration"


def read_input(filename: str) -> List[str]:
    with open(filename) as f:
        content = [x.strip('\n') for x in f.readlines()]
    return content

# This is copied and modified from day18
class Prog23:
    def __init__(self, code: List[str]):
        self.registers: Dict[str, int] = {}
        self.ptr: int = 0
        self.instructions = code
        self.mul_count = 0

    def lookup(self, value: str) -> int:
        try:
            return int(value)
        except:
            return self.registers.get(value, 0)

    def execute(self):
        while self.ptr >= 0 and self.ptr < len(self.instructions):
            inst = self.instructions[self.ptr].split()
            cmd = inst[0]
            if cmd == 'set':
                self.registers[inst[1]] = self.lookup(inst[2])
            elif cmd == 'sub':
                self.registers[inst[1]] = self.lookup(
                    inst[1]) - self.lookup(inst[2])
            elif cmd == 'mul':
                self.registers[inst[1]] = self.lookup(
                    inst[1]) * self.lookup(inst[2])
                self.mul_count += 1
            elif cmd == 'jnz':
                if self.lookup(inst[1]) != 0:
                    self.ptr += self.lookup(inst[2])
                    continue
            else:
                print('unknown command:', inst)

            self.ptr += 1

        return


def part1():
    lines = read_input('input.txt')
    prog = Prog23(lines)
    prog.execute()
    print("Part 1:", prog.mul_count)


def part2():
    # Need to analyze the coprocessor instructions (input) and figure out what it is trying to do
    # Then write it in a normal fashion, or short circuit it a lot
    print('Part 2:', 917)


def main():
    print(f"Day {day_num}: {day_title}")
    part1()
    part2()


if __name__ == '__main__':
    main()
