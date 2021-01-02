from typing import Set

day_num = "25"
day_title = "The Halting Problem"


def part1():
    tape: Set[int] = set()
    curr: int = 0
    state: str = 'a'

    for _ in range(12523873):
        is_one = curr in tape
        if state == 'a' and not is_one:
            tape.add(curr)
            curr += 1
            state = 'b'
        elif state == 'a' and is_one:
            tape.add(curr)
            curr -= 1
            state = 'e'
        elif state == 'b' and not is_one:
            tape.add(curr)
            curr += 1
            state = 'c'
        elif state == 'b' and is_one:
            tape.add(curr)
            curr += 1
            state = 'f'
        elif state == 'c' and not is_one:
            tape.add(curr)
            curr -= 1
            state = 'd'
        elif state == 'c' and is_one:
            tape.remove(curr)
            curr += 1
            state = 'b'
        elif state == 'd' and not is_one:
            tape.add(curr)
            curr += 1
            state = 'e'
        elif state == 'd' and is_one:
            tape.remove(curr)
            curr -= 1
            state = 'c'
        elif state == 'e' and not is_one:
            tape.add(curr)
            curr -= 1
            state = 'a'
        elif state == 'e' and is_one:
            tape.remove(curr)
            curr += 1
            state = 'd'
        elif state == 'f' and not is_one:
            tape.add(curr)
            curr += 1
            state = 'a'
        elif state == 'f' and is_one:
            tape.add(curr)
            curr += 1
            state = 'c'

    print("Part 1:", len(tape))


def main():
    print(f"Day {day_num}: {day_title}")
    part1()


if __name__ == '__main__':
    main()
