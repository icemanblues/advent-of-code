from typing import List

day_num = "02"
day_title = "Corruption Checksum"


def read_input(filename: str) -> List[str]:
    with open(filename) as f:
        content = [x.strip('\n') for x in f.readlines()]
    return content


def part1():
    lines = read_input('input.txt')
    sum = 0
    div = 0
    for line in lines:
        nums = [int(x) for x in line.split()]
        min = nums[0]
        max = nums[0]
        for n in nums:
            if min > n:
                min = n
            if max < n:
                max = n
            for m in nums:
                if n != m and n % m == 0:
                    div += int(n/m)
        sum += max-min

    print(f"Part 1: {sum}")
    print(f'Part 2: {div}')


def main():
    print("Day", day_num, ":", day_title)
    part1()


if __name__ == '__main__':
    main()
