from typing import List

day_num = "10"
day_title = "Knot Hash"

SUFFIX = [17, 31, 73, 47, 23]


def read_input(filename: str) -> str:
    with open(filename) as f:
        content = [x.strip('\n') for x in f.readlines()]
    return content[0]


def swap(numbers: List[int], i: int, j: int):
    temp = numbers[i]
    numbers[i] = numbers[j]
    numbers[j] = temp


def reverse(numbers: List[int], curr: int, inst: int):
    end = curr + inst - 1
    num_swaps = int((inst)/2)
    for x in range(num_swaps):
        i = (curr + x) % len(numbers)
        j = (end - x) % len(numbers)
        swap(numbers, i, j)


def knot_hash(numbers: List[int], inst: List[int]) -> int:
    curr = 0
    skip_size = 0
    for i in inst:
        reverse(numbers, curr, i)
        curr = curr + i + skip_size
        skip_size += 1
    return numbers[0] * numbers[1]


def asciiCode(s: str) -> List[int]:
    b = []
    for i in s:
        b.append(ord(i))
    return b


def toHex(s: int) -> str:
    return "%0.2x" % s


def dense_knot_hash(numbers: List[int], inst: str) -> str:
    a = asciiCode(inst)
    ass = a + SUFFIX
    as64 = ass * 64
    knot_hash(numbers, as64)

    # XOR the 16th elements
    dense_hash: List[int] = list()
    prev = 0
    for i in range(16, len(numbers)+1, 16):
        block = numbers[prev:i]
        acc = block[0]
        for j in range(1, len(block)):
            acc = acc ^ block[j]
        # things to do before iterating
        dense_hash.append(acc)
        prev = i

    # convert dense hash to hex
    result = ""
    for i in dense_hash:
        result = result + toHex(i)
    return result


def main():
    print("Day", day_num, ":", day_title)
    inputs = read_input('input.txt')
    ints = [int(x) for x in inputs.split(',')]
    numbers = list(range(256))
    print("Part 1: ", knot_hash(numbers, ints))
    numbers = list(range(256))
    print("Part 2: ", dense_knot_hash(numbers, inputs))


if __name__ == '__main__':
    main()
