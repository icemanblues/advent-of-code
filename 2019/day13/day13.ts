import fs from 'fs';
import * as readline from 'readline';

const dayNum: string = "13";
const dayTitle: string = "Care Package";

function getOrDefault<K, V>(m: Map<K, V>, k: K, d: V): V {
    return m.has(k) ? m.get(k) : d;
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
    inputCallback: () => Promise<number>;

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
async function progAmp(amp: Amp) {
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
            let input: number = 0;
            if (amp.inputCallback) {
                input = await amp.inputCallback();
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

async function prog(amp: Amp): Promise<number[]> {
    while (!amp.isHalted) {
        await progAmp(amp);
    }
    return amp.outputs;
}

function str(x: number, y: number): string {
    return `${x},${y}`;
}

function count(game: number[], tile: number): number {
    const board: Map<string, number> = new Map<string, number>();

    for (let i: number = 0; i < game.length - 2; i += 3) {
        const key: string = str(game[i], game[i + 1]);
        const value: number = game[i + 2];
        board.set(key, value);
    }

    let counter: number = 0;
    board.forEach((v, k) => {
        if (v === tile) {
            counter++;
        }
    });
    return counter;
}

function paint(game: number[]) {
    let maxX: number = -1;
    let maxY: number = -1;
    const board: Map<string, number> = new Map<string, number>();
    for (let i: number = 0; i < game.length - 2; i += 3) {
        const key: string = str(game[i], game[i + 1]);
        const value: number = game[i + 2];
        board.set(key, value);

        if (maxX < game[i]) {
            maxX = game[i];
        }
        if (maxY < game[i + 1]) {
            maxY = game[i + 1];
        }
    }

    let line: string[] = [];
    for (let y: number = 0; y <= maxY; y++) {
        for (let x: number = 0; x <= maxX; x++) {
            const key: string = str(x, y);
            const value: number = board.get(key);
            if (value === 0) {
                line.push(' ');
            } else {
                line.push(String(value));
            }
            line.push();
        }
        console.log(line.join(''));
        line = [];
    }

    let score: number = board.get(str(-1, 0));
    console.log('score:', score);

    // find the ball and paddle
    let ball: number = -1;
    let paddle: number = -1;
    board.forEach((v, k) => {
        if (v === 4) {
            ball = Number(k.split(/,/)[0]);
        }
        if (v === 3) {
            paddle = Number(k.split(/,/)[0]);
        }
    });

    if (ball < paddle) {
        return -1;
    } else if (ball > paddle) {
        return 1;
    }
    return 0;
}

async function part1() {
    console.log('Part 1');
    const lines: number[] = readInputSync('input.txt').map(Number);
    const output: number[] = [];
    const amp: Amp = new Amp('part1', lines, [], output);
    await prog(amp);
    console.log('part1', count(output, 2));
}

async function part2() {
    console.log('Part 2');
    const lines: number[] = readInputSync('input.txt').map(Number);
    lines[0] = 2;
    const output: number[] = [];
    const amp: Amp = new Amp('Part 2', lines, [], output);

    const joystick = () => {
        paint(output);

        const rl = readline.createInterface({
            input: process.stdin,
            output: process.stdout
        });

        return new Promise<number>(resolve => {
            rl.question('joystick input?', (answer: string) => {
                console.log(`input ${answer}`);
                rl.close();
                resolve(Number(answer));
            });
        });
    }

    const autopilot = () => {
        const play: number = paint(output);
        return new Promise<number>(resolve => resolve(play));
    };

    amp.inputCallback = autopilot;
    await prog(amp);

    paint(output);
}

async function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);

    await part1();
    await part2(); // 12333 too low
}

main();
