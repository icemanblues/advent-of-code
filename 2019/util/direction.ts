export enum Direction {
    North = 1,
    South,
    West,
    East,
}

export const ALL_DIR: Direction[] = [
    Direction.North,
    Direction.South,
    Direction.West,
    Direction.East
];

export const OFFSET: Map<Direction, [number, number]> = new Map([
    [Direction.North, [0, -1]],
    [Direction.South, [0, 1]],
    [Direction.West, [1, 0]],
    [Direction.East, [-1, 0]],
]);

export function move(d: Direction, p: [number, number]): [number, number] {
    const adder: [number, number] = OFFSET.get(d);
    const x: number = p[0] + adder[0];
    const y: number = p[1] + adder[1];
    return [x, y];
}

export function adj(p: [number, number]): [number, number][] {
    return ALL_DIR.map(d => move(d, p));
}

export function reverse(d: Direction): Direction {
    switch (d) {
        case Direction.North:
            return Direction.South;
        case Direction.South:
            return Direction.North;
        case Direction.East:
            return Direction.West;
        case Direction.West:
            return Direction.East;
    }
}

export function findDirection(start: [number, number],
    end: [number, number]): Direction {
    let d: Direction = Direction.North;

    const x: number = end[0] - start[0];
    const y: number = end[1] - start[1];
    OFFSET.forEach((v, k) => {
        if (v[0] === x && v[1] === y) {
            d = k;
        }
    });

    return d;
}
