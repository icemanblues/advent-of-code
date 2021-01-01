from typing import List, Tuple, Dict

day_num = "19"
day_title = "A Series of Tubes"

UP = (-1, 0)
DOWN = (1, 0)
LEFT = (0, -1)
RIGHT = (0, 1)

turn: Dict[Tuple[int, int], List[Tuple[int, int]]] = {
    UP: [LEFT, RIGHT],
    DOWN: [LEFT, RIGHT],
    LEFT: [UP, DOWN],
    RIGHT: [UP, DOWN],
}


def read_input(filename: str) -> List[str]:
    with open(filename) as f:
        content = [x.strip('\n') for x in f.readlines()]
    return content


def valid(d: Tuple[int, int], grid: List[str]) -> bool:
    return (d[0] >= 0 and d[0] < len(grid) and
            d[1] >= 0 and d[1] < len(grid[d[0]]) and
            grid[d[0]][d[1]] != ' ')


def tubing():
    grid = read_input('input.txt')
    x = grid[0].find('|')
    curr = (0, x)
    direct = DOWN
    steps = 0
    word = ''
    while valid(curr, grid):
        tile = grid[curr[0]][curr[1]]
        if tile.isalpha():
            word += tile

        if tile == '|' or tile == '-' or tile.isalpha():
            curr = curr[0] + direct[0], curr[1] + direct[1]
        else:
            # at the plus, need to make a turn
            for dirs in turn[direct]:
                d = curr[0] + dirs[0], curr[1] + dirs[1]
                if valid(d, grid):
                    direct = dirs
                    curr = d
                    break
        steps += 1

    print("Part 1:", word)
    print("Part 2:", steps)


def main():
    print(f"Day {day_num}: {day_title}")
    tubing()


if __name__ == '__main__':
    main()
