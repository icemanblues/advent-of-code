import fs from 'fs';

const dayNum: string = "15";
const dayTitle: string = "Oxygen System";

function readInputSync(filename: string): number[] {
    return fs.readFileSync(filename, "utf-8")
        .trimRight()
        .split(/,/)
        .map(Number);
}

function getOrDefault<K, V>(m: Map<K, V>, k: K, d: V) {
    if (!k) {
        console.log('undefined?', k, m);
    }
    return m.has(k) ? m.get(k) : d;
}

function str(x: number, y: number): string {
    return `${x},${y}`;
}

function strt(tup: [number, number]): string {
    return str(tup[0], tup[1]);
}

function tuple(s: string): [number, number] {
    const nums: number[] = s.split(/,/).map(Number);
    return [nums[0], nums[1]];
}

function opMode(inst: number): [number, number[]] {
    const s: string = String(inst);
    const op: number = Number(s.slice(s.length - 2));
    const mode: number[] = [];

    for (let i = s.length - 3; i >= 0; i--) {
        mode.push(Number(s.charAt(i)));
    }
    while (mode.length !== 3) {
        mode.push(0);
    }

    return [op, mode];
}

class Amp {
    name: string;
    isHalted: boolean;
    intcodes: number[];
    i: number;
    inputs: number[];
    outputs: number[];
    mem: Map<number, number>;
    relBase: number;
    inputCallback: () => number;

    constructor(name: string, ic: number[], ins: number[], outs: number[]) {
        this.name = name;
        this.intcodes = ic;
        this.inputs = ins;
        this.outputs = outs;

        this.isHalted = false;
        this.mem = new Map<number, number>();
        this.i = 0;
        this.relBase = 0;
    }

    _idx(k: number, mode: number): number {
        let idx: number = 0;

        if (mode === 0) { // position
            idx = this.intcodes[k];
        } else if (mode === 1) { // immediate
            idx = k;
        } else if (mode === 2) { // relative
            idx = this.relBase + this.intcodes[k];
        } else {
            console.log('this should never have happened');
        }

        return idx;
    }

    getValue(k: number, mode: number): number {
        const idx: number = this._idx(k, mode);
        if (idx < this.intcodes.length) {
            return this.intcodes[idx];
        }

        return getOrDefault(this.mem, idx, 0);
    }

    setValue(k: number, v: number, mode: number) {
        const idx: number = this._idx(k, mode);

        if (idx < this.intcodes.length) {
            this.intcodes[idx] = v;
        } else {
            this.mem.set(idx, v);
        }
    }
}

// run until it halts, or outputs
function progAmp(amp: Amp) {
    if (amp.isHalted) {
        console.log('Amp is halted.', amp.name);
        return;
    }

    while (true) {
        const inst: number = amp.intcodes[amp.i];
        let op: number;
        let modes: number[];
        [op, modes] = opMode(inst);

        if (99 === op) {
            amp.isHalted = true;
            break;
        }
        else if (1 === op) { // addition
            const a: number = amp.getValue(amp.i + 1, modes[0]);
            const b: number = amp.getValue(amp.i + 2, modes[1]);
            const c: number = a + b;
            amp.setValue(amp.i + 3, c, modes[2]);
            amp.i += 4;
        }
        else if (2 === op) { // multiplication
            const a: number = amp.getValue(amp.i + 1, modes[0]);
            const b: number = amp.getValue(amp.i + 2, modes[1]);
            const c: number = a * b;
            amp.setValue(amp.i + 3, c, modes[2]);
            amp.i += 4;
        }
        else if (3 === op) { // input
            let input: number;
            if (amp.inputCallback) {
                input = amp.inputCallback();
            } else {
                input = amp.inputs.shift();
            }

            amp.setValue(amp.i + 1, input, modes[0]);
            amp.i += 2;
        }
        else if (4 === op) { //output
            const output: number = amp.getValue(amp.i + 1, modes[0]);
            amp.outputs.push(output);
            amp.i += 2;
            return;
        }
        else if (5 === op) { // jump if not zero
            const a: number = amp.getValue(amp.i + 1, modes[0]);
            const b: number = amp.getValue(amp.i + 2, modes[1]);
            if (a !== 0) {
                amp.i = b;
            } else {
                amp.i += 3;
            }
        }
        else if (6 === op) { // jump if zero
            const a: number = amp.getValue(amp.i + 1, modes[0]);
            const b: number = amp.getValue(amp.i + 2, modes[1]);
            if (a === 0) {
                amp.i = b;
            } else {
                amp.i += 3;
            }
        }
        else if (7 === op) {
            const a: number = amp.getValue(amp.i + 1, modes[0]);
            const b: number = amp.getValue(amp.i + 2, modes[1]);
            const c: number = (a < b) ? 1 : 0;
            amp.setValue(amp.i + 3, c, modes[2]);
            amp.i += 4;
        }
        else if (8 === op) {
            const a: number = amp.getValue(amp.i + 1, modes[0]);
            const b: number = amp.getValue(amp.i + 2, modes[1]);
            const c: number = (a === b) ? 1 : 0;
            amp.setValue(amp.i + 3, c, modes[2]);
            amp.i += 4;
        }
        else if (9 === op) {
            const a: number = amp.getValue(amp.i + 1, modes[0]);
            amp.relBase = amp.relBase + a;
            amp.i += 2;
        }
        else {
            console.log('illegal op code', op);
            break;
        }
    }
}

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

function printTile(t: Tile): string {
    if (t === Tile.Wall) {
        return '#';
    } else if (t === Tile.Valid) {
        return '.';
    } else if (t === Tile.Oxygen) {
        return 'O';
    } else {
        return ' ';
    }
}

function displayBoard(grid: Map<string, number>, robot: Robot) {
    let minX: number = Number.MAX_VALUE;
    let minY: number = Number.MAX_VALUE;
    let maxX: number = Number.MIN_VALUE;
    let maxY: number = Number.MIN_VALUE;
    grid.forEach((v, k) => {
        const [x, y] = tuple(k);
        if (x < minX) {
            minX = x;
        }
        if (y < minY) {
            minY = y;
        }
        if (x > maxX) {
            maxX = x;
        }
        if (y > maxY) {
            maxY = y;
        }
    });

    for (let y: number = minY; y <= maxY; y++) {
        const line: string[] = [];
        for (let x: number = minX; x <= maxX; x++) {
            if (strt(robot.loc) === str(x, y)) {
                line.push('D');
            } else {
                const spot: number = grid.get(str(x, y));
                line.push(printTile(spot));
            }
        }
        console.log(line.join(''));
    }
}

function explore(robot: Robot, intcode: number[]): Map<string, number> {
    const grid: Map<string, number> = new Map();
    const starts: string = strt(robot.loc);
    grid.set(starts, Tile.Valid);

    const input: number[] = [];
    const output: number[] = [];
    const amp: Amp = new Amp('robot', intcode, input, output);

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

        // TODO: this should be a call to robot
        // TODO: move should be a robot function
        // update the robot here
        if (resp !== Tile.Wall) {
            robot.loc = writeLoc;
            if (!robot.backtrack) {
                robot.pathStack.push(dir);
            }
        }

        if (resp === Tile.Oxygen) {
            console.log('Part 1', robot.pathStack.length);
        }
    }

    return grid;
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
    console.log('Part 1');
    const robot = new Robot();
    const intcode: number[] = readInputSync('input.txt');
    explore(robot, intcode);
}

function part2() {
    console.log('Part 2');
    const robot = new Robot();
    const intcode: number[] = readInputSync('input.txt');
    const grid = explore(robot, intcode);
    console.log(expand(grid));
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);

    part1();
    part2();
}

main();
