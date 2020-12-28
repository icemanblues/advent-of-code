from typing import Tuple

day_num = "09"
day_title = "Stream Processing"


def read_input(filename: str) -> str:
    with open(filename) as f:
        content = [x.strip('\n') for x in f.readlines()]
    return content[0]


def score(stream: str) -> Tuple[int, int]:
    score = 0
    garbageCount = 0

    depth = 0
    inGarbage = False
    negateNext = False

    for s in stream:
        if negateNext:
            negateNext = False
            continue

        if s == '!':
            negateNext = True
            continue

        if inGarbage and s == '>':
                inGarbage = False
        elif inGarbage:
            garbageCount += 1
        elif s == '{':
            depth+=1
        elif s == '}':
            score += depth
            depth -= 1
        elif s == '<':
            inGarbage = True            
    
    return score, garbageCount


def main():
    print(f"Day {day_num}: {day_title}")
    input = read_input('input.txt')
    p1, p2 = score(input)
    print("Part 1:", p1)
    print("Part 2:", p2)


if __name__ == '__main__':
    main()
