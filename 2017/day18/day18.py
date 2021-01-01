from typing import Dict, List

day_num = "18"
day_title = "Duet"


def read_input(filename: str) -> List[str]:
    with open(filename) as f:
        content = [x.strip('\n') for x in f.readlines()]
    return content


class Prog:
    def __init__(self, id: int, code: List[str], pipein: List[int], pipeout: List[int]):
        self.registers: Dict[str, int] = {}
        self.registers['p'] = id
        self.ID: int = id
        self.ptr: int = 0
        self.snd_count: int = 0
        self.rcv_count: int = 0
        self.instructions = code
        self.pipein = pipein
        self.pipeout = pipeout

    def lookup(self, value: str) -> int:
        try:
            return int(value)
        except:
            return self.registers.get(value, 0)

    def is_blocked(self) -> bool:
        return self.instructions[self.ptr].split()[0] == 'rcv' and len(self.pipein) == 0

    def execute(self):
        while self.ptr >= 0 and self.ptr < len(self.instructions):
            inst = self.instructions[self.ptr].split()
            cmd = inst[0]
            if cmd == 'snd':
                self.pipeout.append(self.lookup(inst[1]))
                self.snd_count += 1
            elif cmd == 'set':
                self.registers[inst[1]] = self.lookup(inst[2])
            elif cmd == 'add':
                self.registers[inst[1]] = self.lookup(
                    inst[1]) + self.lookup(inst[2])
            elif cmd == 'mul':
                self.registers[inst[1]] = self.lookup(
                    inst[1]) * self.lookup(inst[2])
            elif cmd == 'mod':
                self.registers[inst[1]] = self.lookup(
                    inst[1]) % self.lookup(inst[2])
            elif cmd == 'rcv':
                if len(self.pipein) > 0:
                    self.registers[inst[1]] = self.pipein.pop(0)
                    self.rcv_count += 1
                else:
                    return
            elif cmd == 'jgz':
                if self.lookup(inst[1]) > 0:
                    self.ptr += self.lookup(inst[2])
                    continue
            else:
                print('unknown command:', inst)

            self.ptr += 1

        return


def lookup(registers: Dict[str, int], value: str) -> int:
    try:
        return int(value)
    except:
        return registers.get(value, 0)


def execute(instructions: List[str]) -> int:
    registers: Dict[str, int] = {}
    ptr = 0
    sound = 0
    while ptr >= 0 and ptr < len(instructions):
        inst = instructions[ptr].split()
        cmd = inst[0]
        if cmd == 'snd':
            sound = lookup(registers, inst[1])
        elif cmd == 'set':
            registers[inst[1]] = lookup(registers, inst[2])
        elif cmd == 'add':
            registers[inst[1]] = lookup(
                registers, inst[1]) + lookup(registers, inst[2])
        elif cmd == 'mul':
            registers[inst[1]] = lookup(
                registers, inst[1]) * lookup(registers, inst[2])
        elif cmd == 'mod':
            registers[inst[1]] = lookup(
                registers, inst[1]) % lookup(registers, inst[2])
        elif cmd == 'rcv':
            if lookup(registers, inst[1]) != 0:
                return sound
        elif cmd == 'jgz':
            if lookup(registers, inst[1]) > 0:
                ptr += lookup(registers, inst[2])
                continue
        else:
            print('unknown command:', inst)
        ptr += 1

    return sound


def part1():
    lines = read_input('input.txt')
    print("Part 1:", execute(lines))


def part2():
    lines = read_input('input.txt')
    pipe_a: List[int] = []
    pipe_b: List[int] = []
    prog_a = Prog(0, lines, pipe_a, pipe_b)
    prog_b = Prog(1, lines, pipe_b, pipe_a)
    while not prog_a.is_blocked() or not prog_b.is_blocked():
        prog_a.execute()
        prog_b.execute()

    print('Part 2:', prog_b.snd_count)


def main():
    print(f"Day {day_num}: {day_title}")
    part1()
    part2()


if __name__ == '__main__':
    main()
