def readInput(fn):
    nums = []
    with open(fn) as f:
        for line in f:
            nums.append(int(line.rstrip()))
    return nums


def part1(fn):
    print("Part 1")
    inputs = readInput(fn)
    sum = 0
    for n in inputs:
        sum += n
    print(sum)

def part2(fn):
    print("Part 2")
    inputs = readInput(fn)
    freq = 0
    freq_set = {freq}
    idx = 0
    while True:
        freq += inputs[idx]
        if freq in freq_set:
            break
        
        freq_set.add(freq)
        idx = (idx + 1) % len(inputs)
    print("first duplicate frequency:", freq)


### main
print("Day 1: Chronal Calibration")

part1("input01.txt")
part2("input01.txt")
