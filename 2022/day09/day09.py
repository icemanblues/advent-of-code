day_num = "09"
day_title = "Rope Bridge"


def parse(filename):
    with open(filename) as f:
        content = [x.strip('\n').split(' ') for x in f.readlines()]
    return content


movement = {
    'U': [0, 1],
    'L': [-1, 0],
    'D': [0, -1],
    'R': [1, 0],
}


def main():
    print("Day", day_num, ":", day_title)
    inst = parse("input.txt")

    head = (0, 0)
    tail = (0, 0)
    visited = set(tail)
    for i in inst:
        dir = i[0]
        times = int(i[1])

        for j in range(times):
            # move head
            m = movement[dir]
            head = (head[0]+m[0], head[1]+m[1])
            # move tail to match
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
            visited.add(tail)
    print("part1:", len(visited))


if __name__ == '__main__':
    main()
