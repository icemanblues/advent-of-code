import { Amp, progAmp, parseIntcode } from '../intcode';
import { strt, tuple } from '../util';

const dayNum: string = "15";
const dayTitle: string = "Oxygen System";

enum Direction {
    North = 1,
    South,
    West,
    East,
}

const ALL_DIR: Direction[] = [
    Direction.North,
    Direction.South,
    Direction.West,
    Direction.East
];

const offset: Map<Direction, [number, number]> = new Map([
    [Direction.North, [0, -1]],
    [Direction.South, [0, 1]],
    [Direction.West, [1, 0]],
    [Direction.East, [-1, 0]],
]);

function move(d: Direction, p: [number, number]): [number, number] {
    const adder: [number, number] = offset.get(d);
    const x: number = p[0] + adder[0];
    const y: number = p[1] + adder[1];
    return [x, y];
}

function adj(p: [number, number]): [number, number][] {
    return ALL_DIR.map(d => move(d, p));
}

function reverse(d: Direction): Direction {
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

function findDirection(start: [number, number],
    end: [number, number]): Direction {
    let d: Direction = Direction.North;

    const x: number = end[0] - start[0];
    const y: number = end[1] - start[1];
    offset.forEach((v, k) => {
        if (v[0] === x && v[1] === y) {
            d = k;
        }
    });

    return d;
}

enum Tile {
    Wall = 0,
    Valid,
    Oxygen,
}

function explore(robot: Robot, intcode: number[], grid: Map<string, number>): number {
    const starts: string = strt(robot.loc);
    grid.set(starts, Tile.Valid);

    const input: number[] = [];
    const output: number[] = [];
    const amp: Amp = new Amp('robot', intcode, input, output);

    let answer: number = -1;
    while (true) {
        // stop exploring when there is no more forward and no more backwards
        if (robot.backtrack && robot.pathStack.length === 0) {
            break;
        }


        // use the robot to pick our next move
        let dir: Direction = robot.pickNext(grid);
        input.push(dir);
        progAmp(amp);
        let resp = output.shift();

        // update the grid with the results
        const writeLoc: [number, number] = move(dir, robot.loc);
        grid.set(strt(writeLoc), resp);

        // update the robot here
        if (resp !== Tile.Wall) {
            robot.loc = writeLoc;
            if (!robot.backtrack) {
                robot.pathStack.push(dir);
            }
        }

        if (resp === Tile.Oxygen) {
            answer = robot.pathStack.length;
        }
    }

    return answer;
}

class Robot {
    loc: [number, number];
    backtrack: boolean;
    pathStack: Direction[];

    constructor() {
        this.loc = [0, 0];
        this.pathStack = [];
        this.backtrack = false;
    }

    pickNext(grid: Map<string, number>): Direction {
        if (grid.size === 1) {
            return Direction.North;
        }

        const possible: [number, number][] = adj(this.loc)
            .filter(p => !grid.has(strt(p)));

        if (possible.length === 0) {
            this.backtrack = true;
            const d: Direction = this.pathStack.pop();
            return reverse(d);
        } else {
            this.backtrack = false;
        }

        // check the first one, and convert it to a direction
        return findDirection(this.loc, possible[0]);
    }
}

function expand(grid: Map<string, number>): number {
    let oxygenStart: string;
    grid.forEach((v, k) => {
        if (v === Tile.Oxygen) {
            oxygenStart = k;
        }
    });

    let minutes = 0;
    let queue: string[] = [oxygenStart];
    while (queue.length !== 0) {
        let next: string[] = [];
        queue.forEach(s => {
            const possible = adj(tuple(s))
                .filter(p => grid.get(strt(p)) === Tile.Valid)
                .map(t => strt(t));

            if (possible.length !== 0) {
                next = next.concat(possible);
            }
        });

        if (next.length === 0) {
            break;
        }

        minutes++;
        next.forEach(s => grid.set(s, Tile.Oxygen));
        queue = next;
    }

    return minutes;
}


function part1() {
    const robot = new Robot();
    const intcode: number[] = parseIntcode('input.txt');
    const grid: Map<string, number> = new Map();
    console.log('Part 1', explore(robot, intcode, grid));
}

function part2() {
    const robot = new Robot();
    const intcode: number[] = parseIntcode('input.txt');
    const grid: Map<string, number> = new Map();
    explore(robot, intcode, grid);
    console.log('Part 2', expand(grid));
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);
    part1();
    part2();
}

main();
