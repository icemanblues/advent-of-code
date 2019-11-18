###

day_num = "XX"
day_title = "Title"
sub_title_1 = ""
sub_title_2 = ""

def read_input(filename):
    with open(filename) as f:
        content = [x.strip('\n') for x in f.readlines()]
    return content


def part1():
    print("Part 1", sub_title_1)


def part2():
    print("Part 2", sub_title_2)


def main():
    print("Day", day_num, ":", day_title)
    part1()
    part2()


if __name__ == '__main__':
    main()
