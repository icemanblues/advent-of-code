import fs from 'fs';

const dayNum: string = "11";
const dayTitle: string = "Space Police";

function getOrDefault<K, V>(m: Map<K, V>, k: K, d: V) {
    if (!k) {
        console.log('undefined?', k, m);
    }
    return m.has(k) ? m.get(k) : d;
}

function str(x: number, y: number): string {
    return `${x},${y}`;
}

function readInputSync(filename: string): string[] {
    const contents: string = fs.readFileSync(filename, "utf-8");
    const lines: string[] = contents.trimRight().split(/,/);
    return lines;
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

function progPaint(amp: Amp, pb: Paintbot): number {
    amp.inputCallback = () => pb.scan();

    let i: number = 0;
    while (!amp.isHalted) {
        progAmp(amp);
        const o: number = amp.outputs.shift();

        if (i % 2 === 0) { // paint
            pb.paint(o);
        } else { // turn and step
            pb.move(o);
        }
        i++;
    }

    return pb.paintMap.size;
}

enum Direction {
    Up,
    Right,
    Down,
    Left,
}

const directionOffset: Map<Direction, [number, number]> = new Map<Direction, [number, number]>([
    [Direction.Up, [0, -1]],
    [Direction.Right, [1, 0]],
    [Direction.Down, [0, 1]],
    [Direction.Left, [-1, 0]],
]);

function step(start: [number, number], dir: Direction): [number, number] {
    const offset: [number, number] = directionOffset.get(dir);
    return [start[0] + offset[0], start[1] + offset[1]];
}

enum Turn {
    Left,
    Right,
}

function turn90(dir: Direction, turn: Turn): Direction {
    if (turn === Turn.Left) {
        switch (dir) {
            case Direction.Up:
                return Direction.Left;
            case Direction.Left:
                return Direction.Down;
            case Direction.Down:
                return Direction.Right;
            case Direction.Right:
                return Direction.Up;
        }
    } else if (turn === Turn.Right) {
        switch (dir) {
            case Direction.Up:
                return Direction.Right;
            case Direction.Right:
                return Direction.Down;
            case Direction.Down:
                return Direction.Left;
            case Direction.Left:
                return Direction.Up;
        }
    }

    console.log('unknown Turn', turn);
    return Direction.Up;
}

class Paintbot {
    x: number;
    y: number;
    dir: Direction;
    paintMap: Map<string, number>

    constructor() {
        this.x = 0;
        this.y = 0;
        this.dir = Direction.Up;
        this.paintMap = new Map<string, number>();
    }

    scan(): number {
        return getOrDefault(this.paintMap, str(this.x, this.y), 0);
    }

    paint(color: number) {
        this.paintMap.set(str(this.x, this.y), color);
    }

    move(turn: number) {
        this.dir = turn90(this.dir, turn);
        [this.x, this.y] = step([this.x, this.y], this.dir);
    }

    print() {
        for (let y: number = 0; y < 6; y++) {
            const line: string[] = [];
            for (let x: number = 0; x < 40; x++) {
                if (this.paintMap.get(str(x, y)) === 1) {
                    line[x] = '#';
                } else {
                    line[x] = ' ';
                }
            }
            console.log(line.join(''));
        }
    }
}

function part1() {
    console.log('Part 1');
    const lines: number[] = readInputSync('input.txt').map(Number);
    const amp: Amp = new Amp('part1', lines, [], []);
    const paintbot: Paintbot = new Paintbot();
    console.log(progPaint(amp, paintbot));
}

function part2() {
    console.log('Part 2');
    const lines: number[] = readInputSync('input.txt').map(Number);
    const amp: Amp = new Amp('part2', lines, [], []);
    const paintbot: Paintbot = new Paintbot();

    paintbot.paint(1);
    progPaint(amp, paintbot)

    paintbot.print();
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);

    part1();
    part2();
}

main();
