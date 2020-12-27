day_num = "05"
day_title = "A Maze of Twisty Trampolines, All Alike"


def read_input(filename):
    with open(filename) as f:
        content = [int(x.strip('\n')) for x in f.readlines()]
    return content


def jump_out(filename):
    jumps = read_input('input.txt')
    curr = 0
    step_count = 0
    while curr >= 0 and curr < len(jumps):
        old_curr = curr
        curr = curr + jumps[curr]
        jumps[old_curr] += 1
        step_count += 1
    return step_count


def jump_strange(filename):
    jumps = read_input('input.txt')
    curr = 0
    step_count = 0
    while curr >= 0 and curr < len(jumps):
        old_curr = curr
        curr = curr + jumps[curr]
        if jumps[old_curr] >= 3:
            jumps[old_curr] -= 1
        else:
            jumps[old_curr] += 1
        step_count += 1

    return step_count


def part1():
    jump = jump_out('input.txt')
    print(f"Part 1: {jump}")


def part2():
    jump = jump_strange('input.txt')
    print(f"Part 2: {jump}")


def main():
    print("Day", day_num, ":", day_title)
    part1()
    part2()


if __name__ == '__main__':
    main()
