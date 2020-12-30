from typing import Dict, List, Tuple

day_num = "13"
day_title = "Packet Scanners"


UP: int = 0
DOWN: int = 1


class FWLayer:
    def __init__(self, depth: int, size: int):
        self.depth = depth
        self.size = size
        self.scanner = 0
        self.dir = UP

    def tick(self) -> None:
        if self.dir == UP:
            self.scanner += 1
        elif self.dir == DOWN:
            self.scanner -= 1

        if self.dir == UP and self.scanner >= self.size:
            self.dir = DOWN
            self.scanner -= 2
        if self.dir == DOWN and self.scanner < 0:
            self.dir = UP
            self.scanner += 2

    def is_top(self) -> bool:
        return self.scanner == 0

    def severity(self) -> int:
        return self.depth * self.size


def read_input(filename: str) -> Tuple[int, Dict[int, FWLayer]]:
    layers: Dict[int, FWLayer] = {}
    max_depth = 0
    with open(filename) as f:
        content = [x.strip('\n') for x in f.readlines()]
        for line in content:
            parts = line.split(": ")
            depth = int(parts[0])
            size = int(parts[1])
            layer = FWLayer(depth, size)
            layers[depth] = layer
            max_depth = depth
    return max_depth, layers


def tick_tock(layers: Dict[int, FWLayer], max_depth: int) -> int:
    ps = 0
    severity = 0
    while ps <= max_depth:
        if ps in layers and layers[ps].is_top():
            severity += layers[ps].severity()

        ps += 1
        for layer in layers.values():
            layer.tick()
    return severity


def fast_tick_tock(delay: int, layers: Dict[int, FWLayer], max_depth: int) -> Tuple[int, bool]:
    severity = 0
    caught = False
    for ps in range(max_depth+1):
        if ps in layers:
            if (ps + delay) % (2*(layers[ps].size-1)) == 0:
                severity += layers[ps].severity()
                caught = True
    return severity, caught


def part1():
    max_depth, layers = read_input('input.txt')
    severity = tick_tock(layers, max_depth)
    print(f"Part 1: {severity}")
    fast, _ = fast_tick_tock(0, layers, max_depth)
    print(f'Fast 1: {fast}')


def part2():
    max_depth, layers = read_input('input.txt')
    delay = 0
    caught = True
    while caught:
        delay += 1
        _, caught = fast_tick_tock(delay, layers, max_depth)
    print(f"Part 2: {delay}")


def main():
    print(f"Day {day_num}: {day_title}")
    part1()
    part2()


if __name__ == '__main__':
    main()
