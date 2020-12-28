from typing import List, Tuple

day_num = "06"
day_title = "Memory Reallocation"


def read_input(filename: str) -> List[int]:
    with open(filename) as f:
        content = [x.strip('\n') for x in f.readlines()]
        input_bank = content[0].split()
        input_bank = [int(x) for x in input_bank]
    return input_bank


def redistribute(l: List[int]) -> Tuple[int, int]:
    visited = {}
    count = 0

    while tuple(l) not in visited:
        visited[tuple(l)] = count

        # find the index of the max
        m = l[0]
        index_max = 0
        for i in range(1, len(l)):
            if l[i] > m:
                index_max = i
                m = l[i]

        # redistribute the wealth
        l[index_max] = 0
        for i in range(m):
            index_max += 1
            l[index_max % len(l)] += 1

        # increment the step counter
        count += 1

    first_seen = visited[tuple(l)]
    return count, count - first_seen


def main():
    print("Day", day_num, ":", day_title)
    input_bank = read_input("input.txt")
    p1, p2 = redistribute(input_bank)
    print(f"Part 1: {p1}")
    print(f"Part 2: {p2}")


if __name__ == '__main__':
    main()
