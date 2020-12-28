from typing import List

day_num = "01"
day_title = "Inverse Captcha"


def read_input(filename: str) -> List[str]:
    with open(filename) as f:
        content = [x.strip('\n') for x in f.readlines()]
    return content


def part1():
    line = read_input("input.txt")[0]
    curr = line[len(line)-1]
    count = 0
    for c in line:
        if curr == c:
            count += int(c)
        curr = c

    print("Part 1:", count)


def part2():
    line = read_input("input.txt")[0]
    half = len(line)/2
    count = 0
    for i in range(len(line)):
        j = int((i+half) % len(line))
        if line[i] == line[j]:
            count += int(line[i])

    print("Part 2:", count)


def main():
    print("Day", day_num, ":", day_title)
    part1()
    part2()


if __name__ == '__main__':
    main()
