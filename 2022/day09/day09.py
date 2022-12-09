day_num = "09"
day_title = "Rope Bridge"

movement = {
    'U': [0, 1],
    'L': [-1, 0],
    'D': [0, -1],
    'R': [1, 0],
}


def parse(filename):
    with open(filename) as f:
        content = [x.strip('\n').split(' ') for x in f.readlines()]
    return content


def move(head, tail):
    tail_dir = (head[0]-tail[0], head[1]-tail[1])
    x_norm = 1 if tail_dir[0] == 0 else tail_dir[0]//abs(tail_dir[0])
    y_norm = 1 if tail_dir[1] == 0 else tail_dir[1]//abs(tail_dir[1])
    tail_norm = (x_norm, y_norm)
    if tail_dir[0] == 0 and tail_dir[1] == 0:
        pass  # do nothing
    elif abs(tail_dir[0]) == 1 and abs(tail_dir[1]) == 0:
        pass  # do nothing
    elif abs(tail_dir[0]) == 0 and abs(tail_dir[1]) == 1:
        pass  # do nothing
    elif abs(tail_dir[0]) == 1 and abs(tail_dir[1]) == 1:
        pass  # do nothing
    elif abs(tail_dir[0]) >= 1 and abs(tail_dir[1]) >= 1:
        # diagonal
        tail = (tail[0] + tail_norm[0], tail[1] + tail_norm[1])
    elif abs(tail_dir[0]) >= 1:
        # move x
        tail = (tail[0]+tail_norm[0], tail[1])
    elif abs(tail_dir[1]) >= 0:
        # move y
        tail = (tail[0], tail[1]+tail_norm[1])
    else:
        print("unsure how to proceed")
    return tail


def rope_bridge(inst, rope_length):
    rope = [(0, 0) for _ in range(rope_length)]
    visited = {rope[-1]}
    for i in inst:
        for _ in range(int(i[1])):
            # move head
            m = movement[i[0]]
            rope[0] = (rope[0][0]+m[0], rope[0][1] + m[1])
            # move tail(s) to match
            for j in range(1, len(rope)):
                rope[j] = move(rope[j-1], rope[j])
            visited.add(rope[-1])
    return len(visited)


def main():
    print("Day", day_num, ":", day_title)
    inst = parse("input.txt")
    print("part1:", rope_bridge(inst, 2))
    print("part1:", rope_bridge(inst, 10))


if __name__ == '__main__':
    main()
