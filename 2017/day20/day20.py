from typing import List, Dict, Set

day_num = "20"
day_title = "particle swarm"


class Particle:
    def __init__(self, pos: List[int], vel: List[int], acc: List[int]):
        self.pos = pos
        self.vel = vel
        self.acc = acc

    def tick(self):
        for i in range(3):
            self.vel[i] += self.acc[i]
            self.pos[i] += self.vel[i]

    def dist(self) -> int:
        sum = 0
        for p in self.pos:
            sum += abs(p)
        return sum

    def collide(self, other) -> bool:
        return (self.pos[0] == other.pos[0] and
                self.pos[1] == other.pos[1] and
                self.pos[2] == other.pos[2])


def read_input(filename: str) -> List[Particle]:
    particles: List[Particle] = []
    with open(filename) as f:
        for line in f.readlines():
            line = line.strip()
            pva = line.split(', ')
            p = [int(x) for x in pva[0][3:-1].split(',')]
            v = [int(x) for x in pva[1][3:-1].split(',')]
            a = [int(x) for x in pva[2][3:-1].split(',')]
            particles.append(Particle(p, v, a))
    return particles


def part1():
    particles = read_input('input.txt')
    for _ in range(1000):
        for p in particles:
            p.tick()

    min_dist = particles[0].dist()
    min_idx = 0
    for i in range(1, len(particles)):
        if particles[i].dist() < min_dist:
            min_dist = particles[i].dist()
            min_idx = i

    print("Part 1:", min_idx)


def part2():
    particles = read_input('input.txt')
    part_dict: Dict[int, Particle] = {}
    for i in range(len(particles)):
        part_dict[i] = particles[i]

    for _ in range(1000):
        delete: Set[int] = set()
        for i in range(len(particles)-1):
            for j in range(i+1, len(particles)):
                if particles[i].collide(particles[j]):
                    delete.add(i)
                    delete.add(j)

        deleter = list(delete)
        deleter = sorted(deleter, reverse=True)
        for d in deleter:
            particles.pop(d)

        for p in particles:
            p.tick()

    print("Part 2:", len(particles))


def main():
    print(f"Day {day_num}: {day_title}")
    part1()
    part2()


if __name__ == '__main__':
    main()
